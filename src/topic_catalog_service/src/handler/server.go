package handler

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/domain"
)

type topicCatalogServiceServer struct {
	r   domain.TopicRepositoryInterface
	log *logrus.Logger
}

func (srv *topicCatalogServiceServer) ListTopics(_ context.Context, _ *pb.Empty) (*pb.ListTopicsResponse, error) {
	list := srv.r.FindAll()

	srv.log.
		WithField("listCount", len(list)).
		Info("list topic")

	return &pb.ListTopicsResponse{
		Topics: list,
	}, nil
}

func (srv *topicCatalogServiceServer) GetTopic(_ context.Context, req *pb.GetTopicRequest) (*pb.Topic, error) {
	topicID := domain.ToTopicID(req.GetTopicId())
	t, err := srv.r.Find(topicID)
	if err != nil {
		return nil, err
	}

	srv.log.
		WithField("topicID", topicID).
		Info("get topic")

	return t, nil
}

func (srv *topicCatalogServiceServer) AddTopic(_ context.Context, req *pb.AddTopicRequest) (*pb.Topic, error) {
	createdTopic, err := srv.r.Create(req.GetTopic())
	if err != nil {
		return nil, err
	}

	srv.log.
		WithField("topicID", createdTopic.GetId()).
		Info("create topic")

	return createdTopic, nil
}

func (srv *topicCatalogServiceServer) UpdateTopic(_ context.Context, req *pb.UpdateTopicRequest) (*pb.Empty, error) {
	topic := req.GetTopic()
	err := srv.r.Update(topic)
	if err != nil {
		return nil, err
	}

	srv.log.
		WithField("topicID", topic.GetId()).
		Info("update topic")

	return &pb.Empty{}, nil
}

func (srv *topicCatalogServiceServer) DeleteTopic(_ context.Context, req *pb.DeleteTopicRequest) (*pb.Empty, error) {
	topicID := domain.ToTopicID(req.GetTopicId())
	err := srv.r.Delete(topicID)
	if err != nil {
		return nil, err
	}

	srv.log.
		WithField("topicID", topicID).
		Info("delete topic")

	return &pb.Empty{}, nil
}

// NewTopicCatalogServiceServer return TopicCatalogServiceServer
func NewTopicCatalogServiceServer(
	repository domain.TopicRepositoryInterface,
	log *logrus.Logger,
) pb.TopicCatalogServiceServer {
	return &topicCatalogServiceServer{
		r:   repository,
		log: log,
	}
}
