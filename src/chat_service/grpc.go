package main

import (
	"context"
	"fmt"

	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

type chatServiceServer struct{}

func (srv *chatServiceServer) RecvMessage(req *pb.RecvMessageRequest, stream pb.ChatService_RecvMessageServer) error {
	topicID := TopicID(req.GetTopicId())
	userID := UserID(req.GetUser().GetId())

	if !ExistsRoom(topicID) {
		return fmt.Errorf("not found room")
	}

	log.WithField("userID", userID).
		WithField("topicID", topicID).
		Debug("get stream")

	r, err := GetRoom(topicID)
	if err != nil {
		return err
	}

	c := &client{
		userID: userID,
		send:   make(chan *pb.ChatMessage),
	}

	r.Join(c)
	defer r.Leave(c)

	for {
		msg := <-c.send
		stream.Send(msg)
	}
}

func (srv *chatServiceServer) SendMessage(_ context.Context, req *pb.SendMessageRequest) (*pb.Empty, error) {
	topicID := TopicID(req.GetTopicId())
	msg := req.GetMessage()

	log.WithField("topicID", topicID).
		WithField("msg text", msg.GetText()).
		Debug("send")

	r, err := GetRoom(topicID)
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

	if ExistsRoom(topicID) {
		return nil, fmt.Errorf("exists")
	}

	r := newRoom(topicID)
	rooms[topicID] = r
	go r.Run()

	return &pb.Empty{}, nil
}

func (srv *chatServiceServer) ListRoom(_ context.Context, _ *pb.Empty) (*pb.ListRoomResponse, error) {
	var result []*pb.Room
	for id := range rooms {
		result = append(result, &pb.Room{
			TopicId: int32(id),
		})
	}
	return &pb.ListRoomResponse{Rooms: result}, nil
}

func newChatServiceServer() *chatServiceServer {
	return &chatServiceServer{}
}
