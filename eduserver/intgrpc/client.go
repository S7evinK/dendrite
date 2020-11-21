package intgrpc

import (
	"context"

	"github.com/matrix-org/dendrite/eduserver/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type EduServiceClient struct {
	client proto.EduServiceClient
}

func NewEDUServiceGRPCClient(addr string) EduServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logrus.WithError(err).Fatalf("unable to dial: %+v", err)
	}

	c := proto.NewEduServiceClient(conn)
	return EduServiceClient{client: c}
}

func (e EduServiceClient) SendReceiptEvent(ctx context.Context, in *proto.ReceiptEvent) (*proto.EmptyResponse, error) {
	return e.client.SendReceiptEvent(ctx, in)
}

func (e EduServiceClient) SendTypingEvent(ctx context.Context, request *proto.TypingEvent) (*proto.EmptyResponse, error) {
	return e.client.SendTypingEvent(ctx, request)
}

func (e EduServiceClient) SendToDevice(ctx context.Context, request *proto.SendToDeviceEvent) (*proto.EmptyResponse, error) {
	return e.client.SendToDevice(ctx, request)
}
