package infrastructure

import (
	"reflect"

	"github.com/sirupsen/logrus"
	"github.com/somen440/topic-chat/src/common"
	pb "github.com/somen440/topic-chat/src/common/pb"
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/domain"
)

type inMemoryTopicRepository struct {
	topics domain.TopicMap
	log    *logrus.Logger
}

func (r *inMemoryTopicRepository) Create(topic *pb.Topic) (*pb.Topic, error) {
	topicID := domain.TopicID(len(r.topics) + 1)
	topic.Id = int32(topicID)
	r.topics[topicID] = topic

	r.log.
		WithField("newTopic", topic).
		Debug("create topic")

	return topic, nil
}

func (r *inMemoryTopicRepository) Update(topic *pb.Topic) error {
	topicID := domain.TopicID(topic.GetId())

	beforeTopic, err := r.Find(topicID)
	if err != nil {
		return err
	}

	if reflect.DeepEqual(beforeTopic, topic) {
		r.log.
			WithField("beforeTopic", topic).
			Debug("not update topic")
		return nil
	}

	r.topics[topicID] = topic

	r.log.
		WithField("beforeTopic", topic).
		WithField("afterTopic", topic).
		Debug("update topic")

	return nil
}

func (r *inMemoryTopicRepository) Find(topicID domain.TopicID) (*pb.Topic, error) {
	if !r.Exists(topicID) {
		return nil, common.NewRuntimeException("record not found", map[string]interface{}{
			"topicID": topicID,
		})
	}

	r.log.
		WithField("topicID", topicID).
		Debug("find topic")

	return r.topics[topicID], nil
}

func (r *inMemoryTopicRepository) FindAll() []*pb.Topic {
	r.log.
		Debug("find all topic")
	return domain.TopicMapToList(r.topics)
}

func (r *inMemoryTopicRepository) Exists(topicID domain.TopicID) bool {
	_, ok := r.topics[topicID]

	r.log.
		WithField("ok", ok).
		Debug("exists topic")

	return ok
}

func (r *inMemoryTopicRepository) Delete(topicID domain.TopicID) error {
	if !r.Exists(topicID) {
		return common.NewRuntimeException("record not found", map[string]interface{}{
			"topicID": topicID,
		})
	}
	delete(r.topics, topicID)

	r.log.
		WithField("topicID", topicID).
		Debug("delete topic")

	return nil
}

// NewInMemoryTopicRepository is return InMemoryTopicRepository
func NewInMemoryTopicRepository(log *logrus.Logger) domain.TopicRepositoryInterface {
	return &inMemoryTopicRepository{
		topics: domain.CreateTopicMapMock(),
		log:    log,
	}
}
