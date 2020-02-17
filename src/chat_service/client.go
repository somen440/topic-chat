package main

import (
	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

// UserID is user id
type UserID int

type client struct {
	user *pb.User
	send chan *pb.ChatMessage
}

func (c *client) UserID() UserID {
	return UserID(c.user.GetId())
}
