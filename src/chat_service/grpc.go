package main

import (
	"context"

	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

type chatServiceServer struct{}

type client struct {
	id   int
	send chan []byte
}

var clients []*client

func (srv *chatServiceServer) GetStream(_ *pb.Empty, stream pb.ChatService_GetStreamServer) error {
	log.Debug("get stream")

	c := &client{
		id:   len(clients) + 1,
		send: make(chan []byte),
	}
	clients = append(clients, c)

	for {
		select {
		case txt := <-c.send:
			log.WithField("txt", string(txt)).Debug("get txt")
			stream.Send(&pb.Message{
				Text: string(txt),
			})
		}
	}
}

func (srv *chatServiceServer) Send(ctx context.Context, req *pb.Message) (*pb.Empty, error) {
	log.WithField("txt", req.GetText()).Debug("send")
	msg := []byte(req.GetText())
	for _, c := range clients {
		select {
		case c.send <- msg:
			log.WithField("client", c).Debug("send message")
		default:
			log.WithField("client", c).Error("failed send message")
		}
	}
	return &pb.Empty{}, nil
}
