package intgrpc

import (
	"context"
	"encoding/json"
	"net"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"

	userapi "github.com/matrix-org/dendrite/userapi/api"

	"github.com/Shopify/sarama"
	"github.com/matrix-org/dendrite/eduserver/api"
	"github.com/matrix-org/dendrite/eduserver/cache"
	"github.com/matrix-org/dendrite/eduserver/proto"
	"github.com/matrix-org/dendrite/internal/config"
	"github.com/matrix-org/dendrite/internal/setup/kafka"
	"github.com/matrix-org/gomatrixserverlib"
	"github.com/sirupsen/logrus"
)

type EduServiceServer struct {
	Cache                        *cache.EDUCache
	OutputReceiptEventTopic      string
	OutputTypingEventTopic       string
	OutputSendToDeviceEventTopic string
	Producer                     sarama.SyncProducer
	ServerName                   gomatrixserverlib.ServerName
	UserAPI                      userapi.UserInternalAPI
	cfg                          *config.EDUServer
}

func NewEDUServiceGRPC(
	cfg *config.EDUServer,
	eduCache *cache.EDUCache,
	userAPI userapi.UserInternalAPI,
) *EduServiceServer {

	_, producer := kafka.SetupConsumerProducer(&cfg.Matrix.Kafka)

	return &EduServiceServer{
		ServerName:                   cfg.Matrix.ServerName,
		OutputReceiptEventTopic:      cfg.Matrix.Kafka.TopicFor(config.TopicOutputReceiptEvent),
		OutputTypingEventTopic:       cfg.Matrix.Kafka.TopicFor(config.TopicOutputTypingEvent),
		OutputSendToDeviceEventTopic: cfg.Matrix.Kafka.TopicFor(config.TopicOutputSendToDeviceEvent),
		Cache:                        eduCache,
		Producer:                     producer,
		UserAPI:                      userAPI,
		cfg:                          cfg,
	}
}

// Listen starts a standalone grpc server
func (e *EduServiceServer) Listen(addr string) {
	logrus.Debugf("starting gRPC EduServiceServer on %v", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.WithError(err).Fatal("Error listening")
	}

	s := grpc.NewServer()
	proto.RegisterEduServiceServer(s, e)
	if err := s.Serve(l); err != nil {
		logrus.WithError(err).Fatal("error serving")
	}
}

// Attach attaches this service to an existing grpc server
func (e *EduServiceServer) Attach(server *grpc.Server) {
	proto.RegisterEduServiceServer(server, e)
}

func (e *EduServiceServer) SendReceiptEvent(
	ctx context.Context,
	in *proto.ReceiptEvent,
) (*proto.EmptyResponse, error) {
	output := &api.OutputReceiptEvent{
		UserID:    in.UserID,
		RoomID:    in.RoomID,
		EventID:   in.EventID,
		Type:      in.Type,
		Timestamp: gomatrixserverlib.AsTimestamp(in.Timestamp.AsTime()),
	}
	logrus.WithFields(logrus.Fields{
		"userID":    output.UserID,
		"roomID":    output.RoomID,
		"eventID":   output.EventID,
		"type":      output.Type,
		"timestamp": output.Timestamp,
	}).Debug("received grpc event for receipts")
	js, err := json.Marshal(output)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to marshal json: %w", err)
	}
	m := &sarama.ProducerMessage{
		Topic: e.OutputReceiptEventTopic,
		Key:   sarama.StringEncoder(in.RoomID + ":" + in.UserID),
		Value: sarama.ByteEncoder(js),
	}
	_, _, err = e.Producer.SendMessage(m)
	if err != nil {
		return &proto.EmptyResponse{}, status.Errorf(codes.Internal, "%w", err)
	}
	return &proto.EmptyResponse{}, nil
}

func (e *EduServiceServer) SendTypingEvent(
	ctx context.Context,
	in *proto.TypingEvent,
) (*proto.EmptyResponse, error) {
	if in.Typing {
		expireTime := in.OriginServerTS.AsTime().Add(time.Duration(in.Timeout) * time.Millisecond)
		e.Cache.AddTypingUser(in.UserID, in.RoomID, &expireTime)
	} else {
		e.Cache.RemoveUser(in.UserID, in.RoomID)
	}
	err := e.sendTypingEvent(in)
	return &proto.EmptyResponse{}, err
}

func (e *EduServiceServer) sendTypingEvent(in *proto.TypingEvent) error {
	ev := &api.TypingEvent{
		Type:   gomatrixserverlib.MTyping,
		RoomID: in.RoomID,
		UserID: in.UserID,
		Typing: in.Typing,
	}
	ote := &api.OutputTypingEvent{
		Event: *ev,
	}

	if ev.Typing {
		expireTime := in.OriginServerTS.AsTime().Add(
			time.Duration(in.Timeout) * time.Millisecond,
		)
		ote.ExpireTime = &expireTime
	}

	eventJSON, err := json.Marshal(ote)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"room_id": in.RoomID,
		"user_id": in.UserID,
		"typing":  in.Typing,
	}).Infof("(grpc) Producing to topic '%s'", e.OutputTypingEventTopic)

	m := &sarama.ProducerMessage{
		Topic: string(e.OutputTypingEventTopic),
		Key:   sarama.StringEncoder(in.RoomID),
		Value: sarama.ByteEncoder(eventJSON),
	}

	_, _, err = e.Producer.SendMessage(m)
	return err
}

func (e *EduServiceServer) SendToDevice(ctx context.Context, ise *proto.SendToDeviceEvent) (*proto.EmptyResponse, error) {
	devices := []string{}
	_, domain, err := gomatrixserverlib.SplitID('@', ise.UserID)
	if err != nil {
		return &proto.EmptyResponse{}, err
	}

	// If the event is targeted locally then we want to expand the wildcard
	// out into individual device IDs so that we can send them to each respective
	// device. If the event isn't targeted locally then we can't expand the
	// wildcard as we don't know about the remote devices, so instead we leave it
	// as-is, so that the federation sender can send it on with the wildcard intact.
	if domain == e.ServerName && ise.DeviceID == "*" {
		var res userapi.QueryDevicesResponse
		err = e.UserAPI.QueryDevices(context.TODO(), &userapi.QueryDevicesRequest{
			UserID: ise.UserID,
		}, &res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "QueryDevices failed: %w", err)
		}
		for _, dev := range res.Devices {
			devices = append(devices, dev.ID)
		}
	} else {
		devices = append(devices, ise.DeviceID)
	}

	logrus.WithFields(logrus.Fields{
		"user_id":     ise.UserID,
		"num_devices": len(devices),
		"type":        ise.Type,
	}).Infof("Producing to topic '%s'", e.OutputSendToDeviceEventTopic)
	for _, device := range devices {
		ote := &api.OutputSendToDeviceEvent{
			UserID:   ise.UserID,
			DeviceID: device,
			SendToDeviceEvent: gomatrixserverlib.SendToDeviceEvent{
				Sender:  ise.Sender,
				Type:    ise.Type,
				Content: ise.Content,
			},
		}

		eventJSON, err := json.Marshal(ote)
		if err != nil {
			logrus.WithError(err).Error("sendToDevice failed json.Marshal")
			return &proto.EmptyResponse{}, status.Errorf(codes.Internal, "failed to marshal json: %w", err)
		}

		m := &sarama.ProducerMessage{
			Topic: string(e.OutputSendToDeviceEventTopic),
			Key:   sarama.StringEncoder(ote.UserID),
			Value: sarama.ByteEncoder(eventJSON),
		}

		_, _, err = e.Producer.SendMessage(m)
		if err != nil {
			logrus.WithError(err).Error("sendToDevice failed t.Producer.SendMessage")
			return &proto.EmptyResponse{}, status.Errorf(codes.Internal, "sendToDevice failed t.Producer.SendMessage: %w", err)
		}
	}
	return &proto.EmptyResponse{}, nil
}
