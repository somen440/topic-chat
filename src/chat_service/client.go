package main

import (
	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

type client struct {
	id   int
	send chan *pb.ChatMessage
}
