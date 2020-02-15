package infrastructure

import (
	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/domain"
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/exception"
)

type inMemoryTopicRepository struct {
	topics domain.TopicMap
}

func (r *inMemoryTopicRepository) Create(topic *pb.Topic) (*pb.Topic, error) {
	topicID := domain.TopicID(len(r.topics) + 1)
	topic.Id = int32(topicID)
	r.topics[topicID] = topic
	return topic, nil
}

func (r *inMemoryTopicRepository) Update(topic *pb.Topic) error {
	topicID := domain.TopicID(topic.GetId())
	if !r.Exists(topicID) {
		return exception.NewRuntimeException("record not found", map[string]interface{}{
			"topicID": topicID,
		})
	}
	r.topics[topicID] = topic
	return nil
}

func (r *inMemoryTopicRepository) Find(topicID domain.TopicID) (*pb.Topic, error) {
	if !r.Exists(topicID) {
		return nil, exception.NewRuntimeException("record not found", map[string]interface{}{
			"topicID": topicID,
		})
	}
	return r.topics[topicID], nil
}

func (r *inMemoryTopicRepository) FindAll() []*pb.Topic {
	return domain.TopicMapToList(r.topics)
}

func (r *inMemoryTopicRepository) Exists(topicID domain.TopicID) bool {
	_, ok := r.topics[topicID]
	return ok
}

func (r *inMemoryTopicRepository) Delete(topicID domain.TopicID) error {
	if !r.Exists(topicID) {
		return exception.NewRuntimeException("record not found", map[string]interface{}{
			"topicID": topicID,
		})
	}
	delete(r.topics, topicID)
	return nil
}

// NewInMemoryTopicRepository is return InMemoryTopicRepository
func NewInMemoryTopicRepository() domain.TopicRepositoryInterface {
	return &inMemoryTopicRepository{
		topics: domain.CreateTopicMapMock(),
	}
}
