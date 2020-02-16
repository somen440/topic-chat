package main

import (
	"fmt"

	pb "github.com/somen440/topic-chat/src/common/pb"
)

// TopicID is topicID
type TopicID int

// RoomMap is room map
type RoomMap map[TopicID]*room

type room struct {
	id      TopicID
	join    chan *client
	leave   chan *client
	forward chan *pb.ChatMessage
	member  map[*client]bool
}

func (r *room) Run() {
	for {
		select {
		case c := <-r.join:
			log.WithField("userId", c.UserID()).
				Debug("join")
			r.member[c] = true
		case c := <-r.leave:
			log.WithField("userId", c.UserID()).
				Debug("leave")
			delete(r.member, c)
			close(c.send)
		case msg := <-r.forward:
			for c := range r.member {
				log.WithField("userId", c.UserID()).
					Debug("send message")
				select {
				case c.send <- msg:
				default:
					r.Leave(c)
				}
			}
		}
	}
}

func (r *room) NewClientID() int {
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

func (r *room) ExistsUser(userID UserID) bool {
	for c := range r.member {
		if userID == c.UserID() {
			return true
		}
	}
	return false
}

func (r *room) GetClient(userID UserID) (*client, error) {
	for c := range r.member {
		if userID == c.UserID() {
			return c, nil
		}
	}
	return nil, fmt.Errorf("client %v not found", userID)
}

func (r *room) GetID() TopicID {
	return r.id
}

func (r *room) GetUsers() []*pb.User {
	users := []*pb.User{}
	for v := range r.member {
		users = append(users, v.user)
	}
	return users
}

func newRoom(id TopicID) *room {
	r := &room{
		id:      id,
		join:    make(chan *client),
		leave:   make(chan *client),
		forward: make(chan *pb.ChatMessage),
		member:  map[*client]bool{},
	}
	go r.Run()
	return r
}

func (srv *chatServiceServer) ExistsRoom(topicID TopicID) bool {
	for id := range srv.rooms {
		if topicID == id {
			return true
		}
	}
	return false
}

func (srv *chatServiceServer) GetRoom(topicID TopicID) (*room, error) {
	for id, r := range srv.rooms {
		if topicID == id {
			return r, nil
		}
	}
	return nil, fmt.Errorf("room %v not found", topicID)
}

func mockRoom() RoomMap {
	return RoomMap{
		1: newRoom(1),
		2: newRoom(2),
		3: newRoom(3),
		4: newRoom(4),
		5: newRoom(5),
	}
}
