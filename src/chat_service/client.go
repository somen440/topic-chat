package main

import (
	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

type UserID string

type client struct {
	userID UserID
	send   chan *pb.ChatMessage
}
