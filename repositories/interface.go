package repositories

import (
	"bones/entities"
	"errors"
)

var NotFoundError = errors.New("Entity not found")
var DuplicateEmailError = errors.New("This email address is already in use")

type EntityFinder interface {
	Find(id int) (interface{}, error)
}

type UserRepository interface {
	EntityFinder
	// Adds a new user record to the database.
	// Returns an error if the client attempts to
	// insert a user with an email address that is already
	// registered.
	Insert(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindById(id int) (*entities.User, error)
	All() ([]entities.User, error)
}
