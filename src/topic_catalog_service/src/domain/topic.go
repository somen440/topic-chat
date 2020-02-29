package domain

import (
	"sort"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
)

// TopicID is pb.topic.Id alias
type TopicID int32

// TopicMap is pb.topic map alias
type TopicMap map[TopicID]*pb.Topic

// NewTopic is return Topic
func NewTopic(
	topicID TopicID,
	name string,
	description string,
	image string,
) *pb.Topic {
	return &pb.Topic{
		Id:          int32(topicID),
		Name:        name,
		Description: description,
		Image:       image,
	}
}

// ToTopicID is return TopiID
func ToTopicID(id int32) TopicID {
	return TopicID(id)
}

// CreateTopicMapMock is return TopicMap mock
func CreateTopicMapMock() TopicMap {
	return TopicMap{
		1: NewTopic(1, "Webエンジニア", "Webエンジニアとしてこの先生きのこるには", "https://www.pakutaso.com/shared/img/thumb/PAK86_smonitatocode20140517_TP_V.jpg"),
		2: NewTopic(2, "お笑い", "好きなお笑いについて語ろう", "https://www.pakutaso.com/shared/img/thumb/ookawaIMGL2044_TP_V.jpg"),
		3: NewTopic(3, "美味しいTKGについて", "TKGについて", "https://www.pakutaso.com/shared/img/thumb/YAMA845021_TP_V4.jpg"),
		4: NewTopic(4, "赤いきつねvs緑のたぬき", "どっち", "https://www.pakutaso.com/shared/img/thumb/AME19716038_TP_V.jpg"),
		5: NewTopic(5, "新型コロナについて", "新型コロナについて", "https://www.pakutaso.com/shared/img/thumb/KMKC428D343_TP_V.jpg"),
	}
}

// TopicMapToList return topic list
func TopicMapToList(topics TopicMap) []*pb.Topic {
	index := 0
	list := make([]*pb.Topic, len(topics))
	topicIDs := make([]int, len(topics))

	for topicID := range topics {
		topicIDs[index] = int(topicID)
		index++
	}
	index = 0

	sort.Ints(topicIDs)

	for _, topicID := range topicIDs {
		list[index] = topics[TopicID(topicID)]
		index++
	}

	return list
}
