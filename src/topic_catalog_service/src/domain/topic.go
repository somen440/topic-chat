package domain

import (
	"sort"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
)

// TopicID is pb.topic.Id alias
type TopicID int32

// TopicMap is pb.topic map alias
type TopicMap map[TopicID]*pb.Topic

// NewTopic is return Topic
func NewTopic(topicID TopicID, name string) *pb.Topic {
	return &pb.Topic{
		Id:   int32(topicID),
		Name: name,
	}
}

// ToTopicID is return TopiID
func ToTopicID(id int32) TopicID {
	return TopicID(id)
}

// CreateTopicMapMock is return TopicMap mock
func CreateTopicMapMock() TopicMap {
	topics := TopicMap{}
	for i, v := range []string{"topic1", "topic2", "topic3", "topic4", "topic5"} {
		topicID := TopicID(i + 1)
		topics[topicID] = NewTopic(topicID, v)
	}
	return topics
}

// TopicMapToList return topic list
func TopicMapToList(topics TopicMap) []*pb.Topic {
	index := 0
	list := make([]*pb.Topic, len(topics))
	topicIDs := make([]int, len(topics))

	for topicID := range topics {
		topicIDs[index] = int(topicID)
		index++
	}
	index = 0

	sort.Ints(topicIDs)

	for _, topicID := range topicIDs {
		list[index] = topics[TopicID(topicID)]
		index++
	}

	return list
}
