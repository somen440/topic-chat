package main

import (
	"context"

	pb "github.com/somen440/topic-chat/src/frontend/pb"
)

// -----------------------------------------------------------------------------------------
// Auth
// -----------------------------------------------------------------------------------------

func (fe *frontendServer) Join(ctx context.Context, name string) (*pb.User, error) {
	return pb.NewAuthServiceClient(fe.authSvcConn).
		Join(ctx, &pb.JoinRequest{
			Name: name,
		})
}

func (fe *frontendServer) LoggedIn(ctx context.Context, userID string) (*pb.User, error) {
	return pb.NewAuthServiceClient(fe.authSvcConn).
		LoggedIn(ctx, &pb.LoggedInRequest{
			UserId: userID,
		})
}

func (fe *frontendServer) Signout(ctx context.Context, userID string) error {
	_, err := pb.NewAuthServiceClient(fe.authSvcConn).
		Signout(ctx, &pb.SignoutRequest{
			UserId: userID,
		})
	return err
}

// -----------------------------------------------------------------------------------------
// Topic Catalog
// -----------------------------------------------------------------------------------------

func (fe *frontendServer) GetTopic(ctx context.Context, topicID string) (*pb.Topic, error) {
	return pb.NewTopicCatalogServiceClient(fe.topicCatalogSvcConn).
		GetTopic(ctx, &pb.GetTopicRequest{
			Id: topicID,
		})
}

func (fe *frontendServer) ListTopics(ctx context.Context) ([]*pb.Topic, error) {
	res, err := pb.NewTopicCatalogServiceClient(fe.topicCatalogSvcConn).
		ListTopics(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	return res.GetTopics(), err
}
