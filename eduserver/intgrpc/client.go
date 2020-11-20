package intgrpc

import (
	"context"

	"github.com/matrix-org/dendrite/eduserver/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type eduserviceClient struct {
	client proto.EduServiceClient
}

func NewEDUServiceGRPCClient(addr string) *eduserviceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logrus.WithError(err).Fatal("unable to dial")
	}

	return &eduserviceClient{
		client: proto.NewEduServiceClient(conn),
	}
}

func (e eduserviceClient) SendReceiptEvent(ctx context.Context, request *proto.ReceiptEvent) (*proto.EmptyResponse, error) {
	return e.client.SendReceiptEvent(ctx, request)
}

func (e eduserviceClient) SendTypingEvent(ctx context.Context, request *proto.TypingEvent) (*proto.EmptyResponse, error) {
	return e.client.SendTypingEvent(ctx, request)
}

func (e eduserviceClient) SendToDevice(ctx context.Context, request *proto.SendToDeviceEvent) (*proto.EmptyResponse, error) {
	return e.client.SendToDevice(ctx, request)
}
