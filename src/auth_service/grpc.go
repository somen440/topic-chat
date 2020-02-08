package main

import (
	"context"
	"fmt"

	pb "github.com/somen440/topic-chat/src/auth_service/pb"
)

type authServiceServer struct {
	users map[int]*pb.User
}

func (srv *authServiceServer) Join(_ context.Context, req *pb.JoinRequest) (*pb.User, error) {
	name := req.GetName()
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}

	userID := srv.getNewUserID()
	user := newUser(userID, name)
	srv.users[userID] = user

	log.WithField("name", name).
		WithField("userID", userID).
		Debug("join user")

	return user, nil
}

func (srv *authServiceServer) LoggedIn(_ context.Context, req *pb.LoggedInRequest) (*pb.User, error) {
	userID := int(req.GetUserId())

	log.WithField("userID", userID).
		Debug("logged in user")

	user, err := srv.findUser(userID)
	if err != nil {
		return nil, err
	}
	user.LoggedIn = true

	return user, nil
}

func (srv *authServiceServer) Signout(_ context.Context, req *pb.SignoutRequest) (*pb.Empty, error) {
	userID := int(req.GetUserId())

	log.WithField("userID", userID).
		Debug("sighout user")

	user, err := srv.findUser(userID)
	if err != nil {
		return nil, err
	}
	user.LoggedIn = false

	return &pb.Empty{}, nil
}

func (srv *authServiceServer) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	userID := int(req.GetUserId())

	log.WithField("userID", userID).
		Debug("get user")

	user, err := srv.findUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (srv *authServiceServer) GetUserAll(_ context.Context, _ *pb.Empty) (*pb.GetUserAllResponse, error) {
	var users []*pb.User
	for _, v := range srv.users {
		users = append(users, v)
	}
	return &pb.GetUserAllResponse{
		Users: users,
	}, nil
}

func newAuthServiceServer() *authServiceServer {
	return &authServiceServer{
		users: mockUser(),
	}
}
