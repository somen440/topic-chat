package main

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/somen440/topic-chat/src/auth_service/pb"
)

type authServiceServer struct{}

func (srv *authServiceServer) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	name := req.GetName()
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}

	id := strconv.Itoa(len(users) + 1)
	user := &pb.User{
		Id:   id,
		Name: name,
	}
	users = append(users, user)

	return &pb.JoinResponse{
		User: user,
	}, nil
}

func (srv *authServiceServer) LoggedIn(_ context.Context, _ *pb.LoggedInRequest) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
