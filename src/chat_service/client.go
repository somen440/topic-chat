package main

import (
	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

// UserID is user id
type UserID int

type client struct {
	user *pb.User
	send chan *pb.ChatMessage
	add  chan *pb.User
}

// Client interface
type Client interface {
	UserID() UserID
}

func (c *client) UserID() UserID {
	return UserID(c.user.GetId())
}

// NewClient return Client interface
func NewClient(user *pb.User) Client {
	return &client{
		user: user,
		send: make(chan *pb.ChatMessage),
		add:  make(chan *pb.User),
	}
}
