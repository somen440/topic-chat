package main

import (
	"reflect"
	"testing"

	"github.com/somen440/topic-chat/src/topic_catalog_service/src/domain"
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/infrastructure"
)

func TestCreate(t *testing.T) {
	r := infrastructure.NewInMemoryTopicRepository()
	topic, err := r.Create(&domain.Topic{
		Name: "test topic",
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	if topic.Id != 6 {
		t.Errorf("%v", topic)
	}
	if topic.Name != "test topic" {
		t.Errorf("%v", topic)
	}
	actualLen := len(r.FindAll())
	if actualLen != 6 {
		t.Errorf("%v", actualLen)
	}
	actual, err := r.Find(6)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(topic, actual) {
		t.Errorf("%v", actual)
	}
}
