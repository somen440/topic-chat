package main

import (
	"fmt"

	pb "github.com/somen440/topic-chat/src/auth_service/pb"
)

func newUser(id int, name string) *pb.User {
	return &pb.User{
		Id:       int32(id),
		Name:     name,
		LoggedIn: true,
	}
}

func (srv *authServiceServer) findUser(id int) (*pb.User, error) {
	if v, ok := srv.users[id]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("not found user id [%v]", id)
}

func (srv *authServiceServer) getNewUserID() int {
	return len(srv.users) + 1
}

func mockUser() map[int]*pb.User {
	result := map[int]*pb.User{}
	for i, v := range []string{"hoge", "foo", "bar", "foobar", "hogebar"} {
		userID := i + 1
		result[userID] = newUser(userID, v)
	}
	return result
}
