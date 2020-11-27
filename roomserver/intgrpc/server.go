package intgrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/matrix-org/dendrite/internal/caching"
	"github.com/matrix-org/dendrite/internal/config"
	"github.com/matrix-org/dendrite/roomserver/acls"
	"github.com/matrix-org/dendrite/roomserver/internal/helpers"
	"github.com/matrix-org/dendrite/roomserver/intgrpc/helper"
	"github.com/matrix-org/dendrite/roomserver/proto"
	"github.com/matrix-org/dendrite/roomserver/storage"
	"github.com/matrix-org/dendrite/roomserver/version"
	"github.com/matrix-org/gomatrixserverlib"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	s := grpc.NewServer(grpc.UnaryInterceptor(rs.Interceptor))
	proto.RegisterRoomServiceServer(s, rs)
	if err := s.Serve(l); err != nil {
		logrus.WithError(err).Fatal("error serving")
	}
}

// TODO: create useful interceptor; just for debugging
func (rs *RoomServiceServer) Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	defer func(t time.Time) {
		logrus.WithFields(logrus.Fields{
			"duration": time.Since(t),
			"method":   info.FullMethod,
			"response": resp,
			"request":  req,
		}).Debugf("RoomServiceServer Interceptor")
	}(start)
	resp, err = handler(ctx, req)
	return
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
		return res, status.Error(codes.Internal, ErrNoACLs.Error())
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
		return res, status.Error(codes.Internal, fmt.Errorf("GetRoomsByMembership failed: %w", err).Error())
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
		return res, status.Error(codes.Internal, fmt.Errorf("JoinedUsersSetInRooms failed: %w", err).Error())
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
		return res, status.Error(codes.Internal, fmt.Errorf("GetRoomsByMembership failed: %w", err).Error())
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
		return res, status.Error(codes.Internal, fmt.Errorf("GetPublishedRooms failed: %w", err).Error())
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
		return res, status.Error(codes.Internal, fmt.Errorf("RoomInfo failed: %w", err).Error())
	}
	if info == nil {
		return res, status.Error(codes.Internal, ErrMissingRoomInfo.Error())
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

// QueryServerAllowedToSeeEvent checks to see if a server is allowed to see a specific event.
func (rs *RoomServiceServer) QueryServerAllowedToSeeEvent(
	ctx context.Context,
	req *proto.ServerAllowedToSeeEventRequest,
) (*proto.ServerAllowedToSeeEventResponse, error) {
	res := &proto.ServerAllowedToSeeEventResponse{}
	events, err := rs.DB.EventsFromIDs(ctx, []string{req.EventID})
	if err != nil {
		return res, status.Error(codes.Internal, err.Error())
	}
	if len(events) == 0 {
		res.AllowedToSeeEvent = false // event doesn't exist so not allowed to see
		return res, nil
	}
	roomID := events[0].RoomID()
	isServerInRoom, err := helpers.IsServerCurrentlyInRoom(
		ctx, rs.DB, gomatrixserverlib.ServerName(req.ServerName), roomID,
	)
	if err != nil {
		return res, status.Error(codes.Internal, err.Error())
	}
	info, err := rs.DB.RoomInfo(ctx, roomID)
	if err != nil {
		return res, status.Error(codes.Internal, err.Error())
	}
	if info == nil {
		return res, status.Error(codes.Internal, fmt.Errorf("QueryServerAllowedToSeeEvent: no room info for room %s", roomID).Error())
	}

	if res.AllowedToSeeEvent, err = helpers.CheckServerAllowedToSeeEvent(
		ctx, rs.DB, *info, req.EventID, gomatrixserverlib.ServerName(req.ServerName), isServerInRoom,
	); err != nil {
		return res, status.Error(status.Code(err), err.Error())
	}

	return res, nil
}
