package domain

import (
	"reflect"
	"testing"

	pb "github.com/somen440/topic-chat/src/auth_service/pb"
	"github.com/somen440/topic-chat/src/auth_service/src/domain"
)

func TestNewUser(t *testing.T) {
	id := domain.UserID(3)
	name := "user3"
	avator := "avator3"

	expect := &pb.User{
		Id:     int32(id),
		Name:   name,
		Avator: avator,
	}
	actual := domain.NewUser(id, name, avator)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expect [%v] != actual [%v]", expect, actual)
	}
}

func TestConvertUserMapToList(t *testing.T) {
	a := &pb.User{Id: 1, Name: "user1", Avator: "avator1"}
	b := &pb.User{Id: 2, Name: "user2", Avator: "avator2"}
	c := &pb.User{Id: 3, Name: "user3", Avator: "avator3"}

	users := domain.UserMap{
		1: a,
		2: b,
		3: c,
	}

	expect := []*pb.User{a, b, c}
	actual := domain.ConvertUserMapToList(users)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expect [%v] != actual [%v]", expect, actual)
	}
}

func TestCreateMockUserMap(t *testing.T) {
	expect := domain.UserMap{
		1: &pb.User{Id: 1, Name: "user1", Avator: ""},
		2: &pb.User{Id: 2, Name: "user2", Avator: ""},
		3: &pb.User{Id: 3, Name: "user3", Avator: ""},
		4: &pb.User{Id: 4, Name: "user4", Avator: ""},
		5: &pb.User{Id: 5, Name: "user5", Avator: ""},
	}
	actual := domain.CreateMockUserMap()
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expect [%v] != actual [%v]", expect, actual)
	}
}
