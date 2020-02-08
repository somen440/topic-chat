package main

import (
	"context"

	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

type chatServiceServer struct{}

func (srv *chatServiceServer) RecvMessage(req *pb.RecvMessageRequest, stream pb.ChatService_RecvMessageServer) error {
	topicID := TopicID(req.GetTopicId())
	r, err := GetRoom(topicID)
	if err != nil {
		r = newRoom(topicID)
		rooms[topicID] = r
		go r.Run()
	}

	newClientID := r.NewClientId()
	log.WithField("newClientID", newClientID).
		WithField("topicID", topicID).
		Debug("get stream")

	client := &client{
		id:   newClientID,
		send: make(chan *pb.ChatMessage),
	}
	r.Join(client)
	defer r.Leave(client)

	for {
		msg := <-client.send
		stream.Send(msg)
	}
}

func (srv *chatServiceServer) SendMessage(_ context.Context, req *pb.SendMessageRequest) (*pb.Empty, error) {
	topicID := TopicID(req.GetTopicId())
	msg := req.GetMessage()

	log.WithField("topicID", topicID).
		WithField("msg", msg).
		Debug("send")

	r, err := GetRoom(topicID)
	if err != nil {
		return nil, err
	}
	r.Forward(msg)

	return &pb.Empty{}, nil
}

func newChatServiceServer() *chatServiceServer {
	return &chatServiceServer{}
}
