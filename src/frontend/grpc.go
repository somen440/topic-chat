package main

import (
	"context"

	pb "github.com/somen440/topic-chat/src/common/pb"
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

func (fe *frontendServer) LoggedIn(ctx context.Context, userID int) (*pb.User, error) {
	return pb.NewAuthServiceClient(fe.authSvcConn).
		LoggedIn(ctx, &pb.LoggedInRequest{
			UserId: int32(userID),
		})
}

func (fe *frontendServer) Signout(ctx context.Context, userID int) error {
	_, err := pb.NewAuthServiceClient(fe.authSvcConn).
		Signout(ctx, &pb.SignoutRequest{
			UserId: int32(userID),
		})
	return err
}

// -----------------------------------------------------------------------------------------
// Topic Catalog
// -----------------------------------------------------------------------------------------

func (fe *frontendServer) GetTopic(ctx context.Context, topicID int) (*pb.Topic, error) {
	return pb.NewTopicCatalogServiceClient(fe.topicCatalogSvcConn).
		GetTopic(ctx, &pb.GetTopicRequest{
			TopicId: int32(topicID),
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
