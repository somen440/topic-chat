package domain

import (
	"reflect"
	"testing"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/domain"
)

func TestNewTopic(t *testing.T) {
	id := int32(1)
	name := "topic1"
	description := "description1"
	image := "image"

	expect := &pb.Topic{
		Id:          id,
		Name:        name,
		Description: description,
		Image:       image,
	}
	actual := domain.NewTopic(domain.TopicID(id), name, description, image)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expect [%v] != actual [%v]", expect, actual)
	}
}

func TestCreateTopicMapMock(t *testing.T) {
	expect := map[int32]*pb.Topic{
		1: domain.NewTopic(1, "Webエンジニア", "Webエンジニアとしてこの先生きのこるには", ""),
		2: domain.NewTopic(2, "お笑い", "好きなお笑いについて語ろう", ""),
		3: domain.NewTopic(3, "美味しいTKGについて", "TKGについて", ""),
		4: domain.NewTopic(4, "赤いきつねvs緑のたぬき", "どっち", ""),
		5: domain.NewTopic(5, "新型コロナについて", "新型コロナについて", ""),
	}
	actual := domain.CreateTopicMapMock()

	if reflect.DeepEqual(expect, actual) {
		t.Errorf("expect [%v] != actual [%v]", expect, actual)
	}
}

func TestTopicMapToList(t *testing.T) {
	a := domain.NewTopic(domain.TopicID(1), "topic1", "description1", "image1")
	b := domain.NewTopic(domain.TopicID(2), "topic2", "description2", "image2")
	c := domain.NewTopic(domain.TopicID(3), "topic3", "description3", "image3")

	topicMap := domain.TopicMap{
		1: a,
		2: b,
		3: c,
	}
	expectList := []*pb.Topic{a, b, c}
	actualList := domain.TopicMapToList(topicMap)
	if !reflect.DeepEqual(expectList, actualList) {
		t.Errorf("expect [%v] != actual [%v]", expectList, actualList)
	}
}
