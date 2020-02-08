package main

import (
	"fmt"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
)

type TopicID int
type TopicMap map[TopicID]*pb.Topic
type TopicList []*pb.Topic

func newTopic(id TopicID, name string) *pb.Topic {
	return &pb.Topic{
		Id:   int32(id),
		Name: name,
	}
}

func mockTopic() TopicMap {
	topics := TopicMap{}
	for i, v := range []string{"topic1", "topic2", "topic3", "topic4", "topic5"} {
		topicID := TopicID(i + 1)
		topics[topicID] = newTopic(topicID, v)
	}
	return topics
}

func (srv *topicCatalogServiceServer) Topics() TopicList {
	var topics TopicList
	for _, v := range srv.topics {
		topics = append(topics, v)
	}
	return topics
}

func (srv *topicCatalogServiceServer) Topic(topicID TopicID) (*pb.Topic, error) {
	if val, ok := srv.topics[topicID]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("not exists topic %v ", topicID)
}
