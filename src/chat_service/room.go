package main

import (
	pb "github.com/somen440/topic-chat/src/chat_service/pb"
)

type room struct {
	id      int
	join    chan *client
	leave   chan *client
	forword chan *pb.ChatMessage
	member  map[*client]bool
}

func (r *room) Run() {
	for {
		select {
		case c := <-r.join:
			r.member[c] = true
		case c := <-r.leave:
			delete(r.member, c)
			close(c.send)
		case msg := <-r.forword:
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

func (r *room) Forword(msg *pb.ChatMessage) {
	r.forword <- msg
}

func newRoom() *room {
	return &room{
		id:      1, // todo: topic id
		join:    make(chan *client),
		leave:   make(chan *client),
		forword: make(chan *pb.ChatMessage),
		member:  map[*client]bool{},
	}
}
