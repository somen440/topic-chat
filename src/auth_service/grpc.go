package main

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/somen440/topic-chat/src/auth_service/pb"
)

type authServiceServer struct{}

func (srv *authServiceServer) Join(_ context.Context, req *pb.JoinRequest) (*pb.User, error) {
	name := req.GetName()
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}

	id := strconv.Itoa(len(users) + 1)
	user := &pb.User{
		Id:       id,
		Name:     name,
		LoggedIn: true,
	}
	users = append(users, user)

	log.WithField("name", name).
		WithField("id", id).
		WithField("users", users).
		Debug("join user")

	return user, nil
}

func (srv *authServiceServer) LoggedIn(_ context.Context, req *pb.LoggedInRequest) (*pb.User, error) {
	id := req.GetUserId()
	log.WithField("id", id).
		Debug("logged in user")

	user, err := findUser(id)
	if err != nil {
		return nil, err
	}
	user.LoggedIn = true

	return user, nil
}

func (srv *authServiceServer) Signout(_ context.Context, req *pb.SignoutRequest) (*pb.Empty, error) {
	id := req.GetUserId()
	log.WithField("id", id).
		Debug("sighout user")

	user, err := findUser(id)
	if err != nil {
		return nil, err
	}
	user.LoggedIn = false

	return &pb.Empty{}, nil
}

func findUser(id string) (*pb.User, error) {
	for _, v := range users {
		if v.GetId() == id {
			return v, nil
		}
	}
	return nil, fmt.Errorf("not found user")
}
