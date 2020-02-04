package main

import (
	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

type client struct {
	id     int
	stream pb.ChatService_GetStreamServer
}

var clients []*client
