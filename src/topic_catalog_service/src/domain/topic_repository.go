package domain

import (
	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
)

// TopicRepositoryInterface is the topic infrastructure interface
type TopicRepositoryInterface interface {
	Create(*pb.Topic) (*pb.Topic, error)
	Update(*pb.Topic) error
	Find(TopicID) (*pb.Topic, error)
	FindAll() []*pb.Topic
	Exists(TopicID) bool
	Delete(TopicID) error
}
