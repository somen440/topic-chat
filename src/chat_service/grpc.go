package main

import (
	"context"
	"fmt"

	pb "github.com/somen440/topic-chat/src/chat_service/pb"
	"github.com/somen440/topic-chat/src/common"
	"google.golang.org/grpc"
)

type chatServiceServer struct {
	topicCatalogSvcAddr string
	topicCatalogSvcConn *grpc.ClientConn

	authSvcAddr string
	authSvcConn *grpc.ClientConn

	rooms RoomMap
}

func (srv *chatServiceServer) RecvMessage(req *pb.RecvMessageRequest, stream pb.ChatService_RecvMessageServer) error {
	topicID := TopicID(req.GetTopicId())
	userID := UserID(req.GetUserId())
	if topicID == 0 || userID == 0 {
		return fmt.Errorf("invalid param")
	}

	if !srv.ExistsRoom(topicID) {
		return fmt.Errorf("not found room")
	}

	log.WithField("userID", userID).
		WithField("topicID", topicID).
		Debug("get stream")

	r, err := srv.GetRoom(topicID)
	if err != nil {
		return err
	}

	c, err := r.GetClient(userID)
	if err != nil {
		return err
	}

	ctx := stream.Context()
	defer r.Leave(c)

	func() {
		for {
			select {
			case msg := <-c.send:
				stream.Send(msg)
			case <-ctx.Done():
				return
			}
		}
	}()

	log.WithField("userID", userID).
		WithField("topicID", topicID).
		Debug("end stream")

	return nil
}

func (srv *chatServiceServer) SendMessage(_ context.Context, req *pb.SendMessageRequest) (*pb.Empty, error) {
	topicID := TopicID(req.GetTopicId())
	msg := req.GetMessage()
	userID := UserID(msg.GetUserId())

	if topicID == 0 || userID == 0 {
		return nil, common.NewRuntimeException(
			"invalid param",
			map[string]string{
				"expect": fmt.Sprintf("userID > 0 && topicID > 0"),
				"actual": fmt.Sprintf("userID: [%d] topicID: [%d]", userID, topicID),
			},
		)
	}

	log.WithField("topicID", topicID).
		WithField("msg text", msg.GetText()).
		WithField("user id", msg.GetUserId()).
		Debug("send")

	r, err := srv.GetRoom(topicID)
	if err != nil {
		return nil, err
	}
	r.Forward(msg)

	return &pb.Empty{}, nil
}

func (srv *chatServiceServer) CreateRoom(_ context.Context, req *pb.CreateRoomRequest) (*pb.Empty, error) {
	topicID := TopicID(req.GetTopicId())

	log.WithField("topicID", topicID).
		Debug("create")

	if srv.ExistsRoom(topicID) {
		return nil, fmt.Errorf("exists")
	}

	r := newRoom(topicID)
	srv.rooms[topicID] = r
	go r.Run()

	return &pb.Empty{}, nil
}

func (srv *chatServiceServer) JoinRoom(ctx context.Context, req *pb.JoinRoomRequest) (*pb.JoinRoomResponse, error) {
	userID := UserID(req.GetUserId())
	topicID := TopicID(req.GetTopicId())
	if userID == 0 || topicID == 0 {
		return nil, fmt.Errorf("invalid param")
	}
	log.WithField("topicID", topicID).
		WithField("ruserID", userID).
		Debug("join room")

	_, err := srv.GetTopic(ctx, topicID)
	if err != nil {
		return nil, err
	}

	user, err := srv.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	r, err := srv.GetRoom(topicID)
	if err != nil {
		return nil, err
	}

	if r.ExistsUser(userID) {
		return nil, fmt.Errorf("already exists")
	}

	users := r.GetUsers()
	users = append(users, user)

	c := NewClient(user).(*client)
	r.Join(c)

	return &pb.JoinRoomResponse{
		Member: users,
	}, nil
}

func (srv *chatServiceServer) RecvMember(req *pb.RecvMemberRequest, stream pb.ChatService_RecvMemberServer) error {
	topicID := TopicID(req.GetTopicId())
	userID := UserID(req.GetUserId())
	if topicID == 0 || userID == 0 {
		return fmt.Errorf("invalid param")
	}

	if !srv.ExistsRoom(topicID) {
		return fmt.Errorf("not found room")
	}

	log.WithField("topicID", topicID).
		Debug("get stream")

	r, err := srv.GetRoom(topicID)
	if err != nil {
		return err
	}
	c, err := r.GetClient(userID)
	if err != nil {
		return err
	}

	ctx := stream.Context()

	func() {
		for {
			select {
			case user := <-c.add:
				stream.Send(user)
			case <-ctx.Done():
				return
			}
		}
	}()

	log.WithField("userID", userID).
		WithField("topicID", topicID).
		Debug("end stream")

	return nil
}

func (srv *chatServiceServer) GetUser(ctx context.Context, userID UserID) (*pb.User, error) {
	return pb.NewAuthServiceClient(srv.authSvcConn).
		GetUser(ctx, &pb.GetUserRequest{
			UserId: int32(userID),
		})
}

func (srv *chatServiceServer) GetTopic(ctx context.Context, topicID TopicID) (*pb.Topic, error) {
	return pb.NewTopicCatalogServiceClient(srv.topicCatalogSvcConn).
		GetTopic(ctx, &pb.GetTopicRequest{
			TopicId: int32(topicID),
		})
}

func newChatServiceServer() *chatServiceServer {
	return &chatServiceServer{
		rooms: mockRoom(),
	}
}
