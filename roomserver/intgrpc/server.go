package intgrpc

import (
	"context"
	"errors"
	"net"

	"github.com/matrix-org/dendrite/internal/caching"
	"github.com/matrix-org/dendrite/internal/config"
	"github.com/matrix-org/dendrite/roomserver/acls"
	"github.com/matrix-org/dendrite/roomserver/intgrpc/helper"
	"github.com/matrix-org/dendrite/roomserver/proto"
	"github.com/matrix-org/dendrite/roomserver/storage"
	"github.com/matrix-org/dendrite/roomserver/version"
	"github.com/matrix-org/gomatrixserverlib"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type RoomServiceServer struct {
	DB         storage.Database
	Cache      caching.RoomServerCaches
	ServerACLs *acls.ServerACLs
}

func NewRoomServiceServer(cfg *config.RoomServer, cache caching.RoomServerCaches, acl *acls.ServerACLs) *RoomServiceServer {
	db, err := storage.Open(&cfg.Database, cache)
	if err != nil {
		logrus.WithError(err).Panicf("failed to connect to room server db")
	}
	return &RoomServiceServer{
		DB:         db,
		Cache:      cache,
		ServerACLs: acl,
	}
}

// Listen starts a standalone grpc server
func (rs *RoomServiceServer) Listen(addr string) {
	logrus.Debugf("starting gRPC RoomServiceServer on %v", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.WithError(err).Fatal("Error listening")
	}

	s := grpc.NewServer()
	proto.RegisterRoomServiceServer(s, rs)
	if err := s.Serve(l); err != nil {
		logrus.WithError(err).Fatal("error serving")
	}
}

// Attach attaches this service to an existing grpc server
func (rs *RoomServiceServer) Attach(server *grpc.Server) {
	proto.RegisterRoomServiceServer(server, rs)
}

// ErrNoACLs is returned if a room does not have any ACLs.
var ErrNoACLs = errors.New("no server ACL tracking")

// ErrMissingRoomInfo is returned if there are no information about a room.
var ErrMissingRoomInfo = errors.New("missing room info for room")

// QueryServerBannedFromRoom queries the ACLs to check if the server is banned from a room
func (rs *RoomServiceServer) QueryServerBannedFromRoom(
	ctx context.Context, req *proto.ServerBannedFromRoomRequest,
) (*proto.ServerBannedFromRoomResponse, error) {
	res := &proto.ServerBannedFromRoomResponse{}
	if rs.ServerACLs == nil {
		return res, ErrNoACLs
	}
	res.Banned = rs.ServerACLs.IsServerBannedFromRoom(gomatrixserverlib.ServerName(req.ServerName), req.RoomID)
	return res, nil
}

// QuerySharedUsers queries the database to get the count of shared users per room.
func (rs *RoomServiceServer) QuerySharedUsers(
	ctx context.Context, req *proto.SharedUsersRequest,
) (*proto.SharedUsersResponse, error) {
	res := &proto.SharedUsersResponse{}
	roomIDs, err := rs.DB.GetRoomsByMembership(ctx, req.UserID, "join")
	if err != nil {
		return res, err
	}

	roomIDs = append(roomIDs, req.IncludeRoomIDs...)
	excludeMap := make(map[string]bool)
	for _, roomID := range req.ExcludeRoomIDs {
		excludeMap[roomID] = true
	}
	// filter out excluded rooms
	j := 0
	for i := range roomIDs {
		// move elements to include to the beginning of the slice
		// then trim elements on the right
		if !excludeMap[roomIDs[i]] {
			roomIDs[j] = roomIDs[i]
			j++
		}
	}
	roomIDs = roomIDs[:j]

	users, err := rs.DB.JoinedUsersSetInRooms(ctx, roomIDs)
	if err != nil {
		return res, err
	}

	// TODO: Replace JoinedUsersSetInRooms map[string]int with map[string]int64
	userResp := make(map[string]int64, len(users))
	for k, v := range users {
		userResp[k] = int64(v)
	}
	res.UserIDsToCount = userResp
	return res, nil
}

// QueryRoomsForUser gets the users rooms from the database.
func (rs *RoomServiceServer) QueryRoomsForUser(
	ctx context.Context, req *proto.RoomsForUserRequest,
) (*proto.RoomsForUserResponse, error) {
	res := &proto.RoomsForUserResponse{}
	roomIDs, err := rs.DB.GetRoomsByMembership(ctx, req.UserID, req.WantMembership)
	if err != nil {
		return res, err
	}
	res.RoomIDs = roomIDs
	return res, nil
}

// QueryPublishedRooms queries the published rooms.
func (rs *RoomServiceServer) QueryPublishedRooms(
	ctx context.Context,
	req *proto.PublishedRoomsRequest,
) (*proto.PublishedRoomsResponse, error) {
	res := &proto.PublishedRoomsResponse{}
	rooms, err := rs.DB.GetPublishedRooms(ctx)
	if err != nil {
		return res, err
	}
	res.RoomIDs = rooms
	return res, nil
}

// QueryRoomVersionForRoom gets the room version for a given room.
func (rs *RoomServiceServer) QueryRoomVersionForRoom(
	ctx context.Context,
	req *proto.RoomVersionForRoomRequest,
) (*proto.RoomVersionForRoomResponse, error) {
	res := &proto.RoomVersionForRoomResponse{}

	if roomVersion, ok := rs.Cache.GetRoomVersion(req.RoomID); ok {
		res.RoomVersion = helper.ToProtoRoomVersion(roomVersion)
		return res, nil
	}

	info, err := rs.DB.RoomInfo(ctx, req.RoomID)
	if err != nil {
		return res, err
	}
	if info == nil {
		return res, ErrMissingRoomInfo
	}
	res.RoomVersion = helper.ToProtoRoomVersion(info.RoomVersion)
	rs.Cache.StoreRoomVersion(req.RoomID, info.RoomVersion)
	return res, nil
}

// QueryRoomVersionCapabilities gets all version capabilities
// Attention: AvailableRoomVersions is a map[string]string instead of map[RoomVersion]string!
func (rs *RoomServiceServer) QueryRoomVersionCapabilities(
	ctx context.Context,
	req *proto.RoomVersionCapabilitiesRequest,
) (*proto.RoomVersionCapabilitiesResponse, error) {
	res := &proto.RoomVersionCapabilitiesResponse{}
	res.DefaultRoomVersion = helper.ToProtoRoomVersion(version.DefaultRoomVersion())

	res.AvailableRoomVersions = make(map[string]string)
	for v, desc := range version.SupportedRoomVersions() {
		if desc.Stable {
			res.AvailableRoomVersions[string(v)] = "stable"
		} else {
			res.AvailableRoomVersions[string(v)] = "unstable"
		}
	}

	return res, nil
}
