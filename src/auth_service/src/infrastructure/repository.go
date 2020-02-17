package infrastructure

import (
	"github.com/sirupsen/logrus"
	pb "github.com/somen440/topic-chat/src/auth_service/pb"
	"github.com/somen440/topic-chat/src/auth_service/src/domain"
	"github.com/somen440/topic-chat/src/common"
)

type inMemoryUserRepository struct {
	users domain.UserMap
	log   *logrus.Logger
}

func (r *inMemoryUserRepository) Exists(userID domain.UserID) bool {
	_, ok := r.users[userID]
	return ok
}

func (r *inMemoryUserRepository) Create(user *pb.User) (*pb.User, error) {
	userID := domain.UserID(len(r.users) + 1)
	createdUser := domain.NewUser(userID, user.GetName())
	r.users[userID] = createdUser
	return createdUser, nil
}

func (r *inMemoryUserRepository) Find(userID domain.UserID) (*pb.User, error) {
	v, ok := r.users[userID]
	if !ok {
		return nil, common.NewRuntimeException("not found", map[string]interface{}{
			"userID": userID,
		})
	}
	return v, nil
}

func (r *inMemoryUserRepository) FindAll() []*pb.User {
	r.log.Debug("find all")
	return domain.ConvertUserMapToList(r.users)
}

// NewInMemoryUserRepository return user repository
func NewInMemoryUserRepository(users domain.UserMap, log *logrus.Logger) domain.UserRepository {
	return &inMemoryUserRepository{
		users: users,
		log:   log,
	}
}
