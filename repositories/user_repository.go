package repositories

import (
	"bones/entities"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type UserRepository interface {
	Insert(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	All() ([]entities.User, error)
}

var userRepository UserRepository = new(MongoUserRepository)

// Override userRepository, e.g. for tests
func SetUserRepository(repo UserRepository) {
	userRepository = repo
}

func Users() UserRepository {
	return userRepository
}

type MongoUserRepository struct{}

func (r MongoUserRepository) Insert(user *entities.User) error {
	return r.collection().Insert(user)
}

func (r MongoUserRepository) FindByEmail(email string) (*entities.User, error) {
	user := new(entities.User)
	err := r.collection().Find(bson.M{"email": email}).One(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r MongoUserRepository) All() ([]entities.User, error) {
	users := []entities.User{}
	err := r.collection().Find(nil).All(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *MongoUserRepository) collection() *mgo.Collection {
	return session.DB(database).C("users")
}
