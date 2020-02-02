package main

import (
	"context"
	"fmt"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
)

type topicCatalogServiceServer struct{}

func (t *topicCatalogServiceServer) ListTopics(_ context.Context, _ *pb.Empty) (*pb.ListTopicsResponse, error) {
	log.Info("list topics")

	return &pb.ListTopicsResponse{
		Topics: topics,
	}, nil
}

func (t *topicCatalogServiceServer) GetTopic(_ context.Context, req *pb.GetTopicRequest) (*pb.Topic, error) {
	log.WithField("id", req.GetId()).Info("get topic")

	for _, v := range topics {
		if v.GetId() == req.GetId() {
			return v, nil
		}
	}
	return nil, fmt.Errorf("not found")
}
