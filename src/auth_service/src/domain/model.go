package domain

import (
	"sort"

	pb "github.com/somen440/topic-chat/src/common/pb"
)

// UserID is pb user id
type UserID int32

// UserMap is user map
type UserMap map[UserID]*pb.User

// NewUser return user
func NewUser(id UserID, name string) *pb.User {
	return &pb.User{
		Id:   int32(id),
		Name: name,
	}
}

// ConvertUserMapToList is user map -> user list
func ConvertUserMapToList(users UserMap) []*pb.User {
	index := 0
	len := len(users)
	result := make([]*pb.User, len)
	userIDs := make([]int, len)

	for v := range users {
		userIDs[index] = int(v)
		index++
	}
	index = 0
	sort.Ints(userIDs)

	for _, v := range userIDs {
		result[index] = users[UserID(v)]
		index++
	}
	return result
}

// CreateMockUserMap return user map
func CreateMockUserMap() UserMap {
	results := UserMap{}
	for i, v := range []string{"user1", "user2", "user3", "user4", "user5"} {
		userID := UserID(i + 1)
		results[userID] = NewUser(userID, v)
	}
	return results
}
