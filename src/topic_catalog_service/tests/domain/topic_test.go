package domain

import (
	"reflect"
	"testing"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"

	"github.com/somen440/topic-chat/src/topic_catalog_service/src/domain"
)

func TestNewTopic(t *testing.T) {
	expect := &pb.Topic{
		Id:   1,
		Name: "topic1",
	}
	actual := domain.NewTopic(1, "topic1")

	if reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestCreateTopicMapMock(t *testing.T) {
	expect := map[int32]*pb.Topic{
		1: &pb.Topic{Id: 1, Name: "topic1"},
		2: &pb.Topic{Id: 2, Name: "topic2"},
		3: &pb.Topic{Id: 3, Name: "topic3"},
		4: &pb.Topic{Id: 4, Name: "topic4"},
		5: &pb.Topic{Id: 5, Name: "topic5"},
	}
	actual := domain.CreateTopicMapMock()

	if reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestTopicID(t *testing.T) {
	expect := domain.TopicID(3)
	actual := domain.NewTopic(3, "test topic").TopicID()
	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestTopicMapToList(t *testing.T) {
	a := domain.NewTopic(1, "topic1")
	b := domain.NewTopic(2, "topic2")
	c := domain.NewTopic(3, "topic3")

	topicMap := domain.TopicMap{
		1: a,
		2: b,
		3: c,
	}
	expectList := domain.TopicList{a, b, c}
	actualList := domain.TopicMapToList(topicMap)
	if !reflect.DeepEqual(expectList, actualList) {
		t.Errorf("%v != %v", expectList, actualList)
	}
}
