package intgrpc

import (
	"context"
	"time"

	"github.com/matrix-org/dendrite/roomserver/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type RoomServiceClient struct {
	client proto.RoomServiceClient
}

func NewRoomServerServiceGRPCClient(addr string) RoomServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(clientInterceptor))
	if err != nil {
		logrus.WithError(err).Fatalf("unable to dial: %+v", err)
	}
	c := proto.NewRoomServiceClient(conn)
	return RoomServiceClient{client: c}
}

func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	logrus.Debugf("clientInterceptor: method=%s; Duration=%s; Error=%v", method, time.Since(start), err)
	return err
}

func (r RoomServiceClient) QueryServerBannedFromRoom(ctx context.Context, in *proto.ServerBannedFromRoomRequest) (*proto.ServerBannedFromRoomResponse, error) {
	return r.client.QueryServerBannedFromRoom(ctx, in)
}

func (r RoomServiceClient) QuerySharedUsers(ctx context.Context, in *proto.SharedUsersRequest) (*proto.SharedUsersResponse, error) {
	return r.client.QuerySharedUsers(ctx, in)
}

func (r RoomServiceClient) QueryRoomsForUser(ctx context.Context, in *proto.RoomsForUserRequest) (*proto.RoomsForUserResponse, error) {
	return r.client.QueryRoomsForUser(ctx, in)
}

func (r RoomServiceClient) QueryPublishedRooms(ctx context.Context, in *proto.PublishedRoomsRequest) (*proto.PublishedRoomsResponse, error) {
	return r.client.QueryPublishedRooms(ctx, in)
}

func (r RoomServiceClient) QueryRoomVersionForRoom(ctx context.Context, in *proto.RoomVersionForRoomRequest) (*proto.RoomVersionForRoomResponse, error) {
	return r.client.QueryRoomVersionForRoom(ctx, in)
}

func (r RoomServiceClient) QueryRoomVersionCapabilities(ctx context.Context, in *proto.RoomVersionCapabilitiesRequest) (*proto.RoomVersionCapabilitiesResponse, error) {
	return r.client.QueryRoomVersionCapabilities(ctx, in)
}

func (r RoomServiceClient) QueryServerAllowedToSeeEvent(ctx context.Context, in *proto.ServerAllowedToSeeEventRequest) (*proto.ServerAllowedToSeeEventResponse, error) {
	return r.client.QueryServerAllowedToSeeEvent(ctx, in)
}

func (r RoomServiceClient) QueryServerJoinedToRoomGRPC(ctx context.Context, in *proto.ServerJoinedToRoomRequest) (*proto.ServerJoinedToRoomResponse, error) {
	return r.client.QueryServerJoinedToRoom(ctx, in)
}
