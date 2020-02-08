package main

import (
	"fmt"

	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

type TopicID int

type room struct {
	id      TopicID
	join    chan *client
	leave   chan *client
	forward chan *pb.ChatMessage
	member  map[*client]bool
}

var rooms = map[TopicID]*room{}

func (r *room) Run() {
	for {
		select {
		case c := <-r.join:
			r.member[c] = true
		case c := <-r.leave:
			delete(r.member, c)
			close(c.send)
		case msg := <-r.forward:
			for c := range r.member {
				c.send <- msg
			}
		}
	}
}

func (r *room) NewClientId() int {
	return len(r.member) + 1
}

func (r *room) Join(c *client) {
	r.join <- c
}

func (r *room) Leave(c *client) {
	r.leave <- c
}

func (r *room) Forward(msg *pb.ChatMessage) {
	r.forward <- msg
}

func newRoom(id TopicID) *room {
	return &room{
		id:      id,
		join:    make(chan *client),
		leave:   make(chan *client),
		forward: make(chan *pb.ChatMessage),
		member:  map[*client]bool{},
	}
}

func GetRoom(topicID TopicID) (*room, error) {
	for id, r := range rooms {
		if topicID == id {
			return r, nil
		}
	}
	return nil, fmt.Errorf("room %v not found", topicID)
}
