package repositories

import (
	"bones/entities"
)

type UserRepository interface {
	Insert(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindById(id int) (*entities.User, error)
	All() ([]entities.User, error)
}

var Users UserRepository = new(SqlUserRepository)

type SqlUserRepository struct{}

func (r SqlUserRepository) Insert(user *entities.User) error {
	return exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
}

func (r SqlUserRepository) FindByEmail(email string) (*entities.User, error) {
	rc := new(usersRowCollector)
	err := queryRow(rc, "SELECT * FROM users WHERE email = $1 LIMIT 1", email)
	return rc.firstOrErr(err)
}

func (r SqlUserRepository) FindById(id int) (*entities.User, error) {
	rc := new(usersRowCollector)
	err := queryRow(rc, "SELECT * FROM users WHERE id = $1 LIMIT 1", id)
	return rc.firstOrErr(err)
}

func (r SqlUserRepository) All() ([]entities.User, error) {
	rc := new(usersRowCollector)
	err := query(rc, "SELECT * FROM users")
	return rc.allOrErr(err)
}

type usersRowCollector struct {
	users []entities.User
}

func (rc *usersRowCollector) collectRow(rs rowScanner) error {
	user := entities.User{}

	err := rs.Scan(&user.Id, &user.Password, &user.Email)

	if err != nil {
		return err
	}

	rc.users = append(rc.users, user)

	return nil
}

func (rc *usersRowCollector) allOrErr(err error) ([]entities.User, error) {
	if err != nil {
		return nil, err
	}

	return rc.users, nil
}

func (rc *usersRowCollector) firstOrErr(err error) (*entities.User, error) {
	if err != nil {
		return nil, err
	}

	return rc.first(), nil
}

func (rc *usersRowCollector) first() *entities.User {
	if len(rc.users) > 0 {
		return &rc.users[0]
	}

	return nil
}
