package handler

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/somen440/topic-chat/src/auth_service/pb"
	"github.com/somen440/topic-chat/src/auth_service/src/domain"
	"github.com/somen440/topic-chat/src/common"
)

type authServiceServer struct {
	r   domain.UserRepository
	log *logrus.Logger
}

func (srv *authServiceServer) Join(ctx context.Context, req *pb.JoinRequest) (*pb.User, error) {
	name := req.GetName()
	avator := req.GetAvator()
	user, err := srv.r.Create(&pb.User{
		Name:   name,
		Avator: avator,
	})
	return user, err
}

func (srv *authServiceServer) LoggedIn(ctx context.Context, req *pb.LoggedInRequest) (*pb.User, error) {
	return nil, &common.NotImplementedException{}
}

func (srv *authServiceServer) Signout(ctx context.Context, req *pb.SignoutRequest) (*pb.Empty, error) {
	return nil, &common.NotImplementedException{}
}

func (srv *authServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	userID := domain.UserID(req.GetUserId())
	user, err := srv.r.Find(userID)
	return user, err
}

func (srv *authServiceServer) GetUserAll(ctx context.Context, _ *pb.Empty) (*pb.GetUserAllResponse, error) {
	return &pb.GetUserAllResponse{
		Users: srv.r.FindAll(),
	}, nil
}

// NewAuthServiceServer return pb AuthServiceServer
func NewAuthServiceServer(r domain.UserRepository, log *logrus.Logger) pb.AuthServiceServer {
	return &authServiceServer{
		r:   r,
		log: log,
	}
}
