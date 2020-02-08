package main

import (
	"context"

	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

type chatServiceServer struct {
	room *room
}

func (srv *chatServiceServer) RecvMessage(_ *pb.Empty, stream pb.ChatService_RecvMessageServer) error {
	newID := srv.room.NewClientId()
	log.WithField("id", newID).Debug("get stream")

	client := &client{
		id:   newID,
		send: make(chan *pb.ChatMessage),
	}
	srv.room.Join(client)
	defer srv.room.Leave(client)

	for {
		msg := <-client.send
		stream.Send(msg)
	}
}

func (srv *chatServiceServer) SendMessage(_ context.Context, msg *pb.ChatMessage) (*pb.Empty, error) {
	log.Debug("send")
	srv.room.Forword(msg)
	return &pb.Empty{}, nil
}

func newChatServiceServer(r *room) *chatServiceServer {
	return &chatServiceServer{
		room: r,
	}
}
