package domain

import (
	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
)

// Topic is pb.topic alias
type Topic pb.Topic

// TopicID is pb.topic.Id alias
type TopicID int32

// TopicMap is pb.topic map alias
type TopicMap map[TopicID]*Topic

// TopicList is pb.topic list alias
type TopicList []*Topic

// NewTopic is return Topic
func NewTopic(topicID TopicID, name string) *Topic {
	return &Topic{
		Id:   int32(topicID),
		Name: name,
	}
}

// TopicID is return TopiID
func (t *Topic) TopicID() TopicID {
	return TopicID(t.Id)
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
func TopicMapToList(topics TopicMap) TopicList {
	var list TopicList
	for _, v := range topics {
		list = append(list, v)
	}
	return list
}
