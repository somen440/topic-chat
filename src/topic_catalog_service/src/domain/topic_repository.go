package domain

// TopicRepositoryInterface is the topic infrastructure interface
type TopicRepositoryInterface interface {
	Create(*Topic) (*Topic, error)
	Update(*Topic) error
	Find(TopicID) (*Topic, error)
	FindAll() TopicList
	Exists(TopicID) bool
	Delete(*Topic) error
}
