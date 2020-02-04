package main

import (
	"context"
	"io"

	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

type chatServiceServer struct{}

func (srv *chatServiceServer) GetStream(stream pb.ChatService_GetStreamServer) error {
	newID := len(clients) + 1
	log.WithField("id", newID).Debug("get stream")

	client := &client{
		id:     newID,
		stream: stream,
	}
	clients = append(clients, client)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			continue
		}
		if err != nil {
			return err
		}
		stream.Send(req)
	}
}

func (srv *chatServiceServer) Send(_ context.Context, msg *pb.ChatMessage) (*pb.Empty, error) {
	log.WithField("msg", msg.GetText()).Debug("send")
	for _, c := range clients {
		c.stream.Send(msg)
	}
	return &pb.Empty{}, nil
}
