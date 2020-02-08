package main

import (
	"context"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
)

type topicCatalogServiceServer struct {
	topics TopicMap
}

func (srv *topicCatalogServiceServer) ListTopics(_ context.Context, _ *pb.Empty) (*pb.ListTopicsResponse, error) {
	log.Debug("list topics")

	return &pb.ListTopicsResponse{
		Topics: srv.Topics(),
	}, nil
}

func (srv *topicCatalogServiceServer) GetTopic(_ context.Context, req *pb.GetTopicRequest) (*pb.Topic, error) {
	topicID := TopicID(req.GetTopicId())

	log.WithField("topicID", topicID).
		Debug("get topic")

	return srv.Topic(topicID)
}

func newTopicCatalogServiceServer() *topicCatalogServiceServer {
	return &topicCatalogServiceServer{
		topics: mockTopic(),
	}
}
