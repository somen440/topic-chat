package handler

import (
	"context"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"

	"github.com/somen440/topic-chat/src/topic_catalog_service/src/domain"
)

type topicCatalogServiceServer struct {
	r domain.TopicRepositoryInterface
}

func (srv *topicCatalogServiceServer) ListTopics(_ context.Context, _ *pb.Empty) (*pb.ListTopicsResponse, error) {
}

func (srv *topicCatalogServiceServer) GetTopic(_ context.Context, req *pb.GetTopicRequest) (*pb.Topic, error) {

}

func (srv *topicCatalogServiceServer) AddTopic(_ context.Context, req *pb.AddTopicRequest) (*pb.Topic, error) {

}

func (srv *topicCatalogServiceServer) UpdateTopic(_ context.Context, req *pb.UpdateTopicRequest) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (srv *topicCatalogServiceServer) DeleteTopic(_ context.Context, req *pb.DeleteTopicRequest) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

// NewTopicCatalogServiceServer return TopicCatalogServiceServer
func NewTopicCatalogServiceServer(repository domain.TopicRepositoryInterface) pb.TopicCatalogServiceServer {
	return &topicCatalogServiceServer{
		r: repository,
	}
}
