package main

import (
	"io"

	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

type chatServiceServer struct{}

func (srv *chatServiceServer) RecvMessage(_ *pb.Empty, stream pb.ChatService_RecvMessageServer) error {
	newID := len(clients) + 1
	log.WithField("id", newID).Debug("get stream")

	client := &client{
		id:   newID,
		send: make(chan []byte),
	}
	clients[client] = true

	for {
		msg := <-client.send
		stream.Send(&pb.ChatMessage{Text: string(msg)})
	}
}

func (srv *chatServiceServer) SendMessage(stream pb.ChatService_SendMessageServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Empty{})
		}

		if err != nil {
			return err
		}

		for c := range clients {
			select {
			case c.send <- []byte(req.GetText()):
				log.WithField("msg", req.GetText()).
					WithField("id", c.id).
					Debug("send")
			default:
				delete(clients, c)
				close(c.send)
			}
		}

		return stream.SendAndClose(&pb.Empty{})
	}
}
