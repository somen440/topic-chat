package domain

import (
	pb "github.com/somen440/topic-chat/src/common/pb"
)

// UserRepository is user infrastructure interface
type UserRepository interface {
	Exists(UserID) bool
	Create(*pb.User) (*pb.User, error)
	Find(UserID) (*pb.User, error)
	FindAll() []*pb.User
}
