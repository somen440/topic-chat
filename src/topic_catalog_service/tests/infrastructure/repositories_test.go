package infrastructure

import (
	"reflect"
	"testing"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/domain"
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/infrastructure"
)

func TestInMemoryTopicRepository(t *testing.T) {
	var r domain.TopicRepositoryInterface
	var err error
	var topic *pb.Topic

	mockTopic := &pb.Topic{Name: "hoge"}

	t.Run("NewInMemoryTopicRepository", func(t *testing.T) {
		expect := "*infrastructure.inMemoryTopicRepository"

		r = infrastructure.NewInMemoryTopicRepository()
		actual := reflect.TypeOf(r).String()

		if expect != actual {
			t.Errorf("expect [%v] != actual [%v]", expect, actual)
		}
	})

	t.Run("Create Success", func(t *testing.T) {
		topic, err = r.Create(mockTopic)
		if err != nil {
			t.Errorf("create error")
		}
	})

	t.Run("Created topic", func(t *testing.T) {
		expect := &pb.Topic{
			Id:   6,
			Name: mockTopic.GetName(),
		}
		actual := topic
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("%v != %v", expect, actual)
		}
	})
}
