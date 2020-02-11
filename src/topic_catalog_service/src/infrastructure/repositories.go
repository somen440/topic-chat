package infrastructure

import (
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/domain"
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/exception"
)

// InMemoryTopicRepository is save to variable
type InMemoryTopicRepository struct {
	topics domain.TopicMap
}

// Create topic
func (r *InMemoryTopicRepository) Create(topic domain.Topic) (domain.Topic, error) {
}

// Update topic
func (r *InMemoryTopicRepository) Update(topic domain.Topic) error {

}

// FindAll topic
func (r *InMemoryTopicRepository) Find(topicID domain.TopicID) (domain.Topic, error) {

}

// FindAll topics
func (r *InMemoryTopicRepository) FindAll() domain.TopicList {
}

// Exists is exists topic from topicID
func (r *InMemoryTopicRepository) Exists(topicID domain.TopicID) bool {
	_, ok := r.topics[topicID]
	return ok
}

// Delete topic
func (r *InMemoryTopicRepository) Delete(topic domain.Topic) error {
	topicID := topic.TopicID()
	if !r.Exists(topicID) {
		return exception.NewRuntimeException("record not found", map[string]interface{}{
			"topicID": topicID,
		})
	}
	delete(r.topics, topicID)
	return nil
}

// NewInMemoryTopicRepository is return InMemoryTopicRepository
func NewInMemoryTopicRepository() *InMemoryTopicRepository {
	return &InMemoryTopicRepository{
		topics: domain.CreateTopicMapMock(),
	}
}
