package sqlrepositories

import (
	"bones/entities"
	"bones/repositories"
	"database/sql"
	"github.com/peterskeide/veil"
)

func NewUserRepository() *UserRepository {
	return &UserRepository{dbveil}
}

type UserRepository struct {
	veil.Veil
}

func (r UserRepository) Insert(user *entities.User) error {
	_, err := r.FindByEmail(user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return r.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
		}

		return err
	}

	return repositories.DuplicateEmailError
}

func (r UserRepository) FindByEmail(email string) (*entities.User, error) {
	rc := new(usersRowCollector)
	err := r.QueryRow(rc, "SELECT * FROM users WHERE email = $1 LIMIT 1", email)
	return rc.firstOrErr(err)
}

func (r UserRepository) FindById(id int) (*entities.User, error) {
	rc := new(usersRowCollector)
	err := r.QueryRow(rc, "SELECT * FROM users WHERE id = $1 LIMIT 1", id)
	return rc.firstOrErr(err)
}

func (r UserRepository) Find(id int) (interface{}, error) {
	return r.FindById(id)
}

func (r UserRepository) All() ([]entities.User, error) {
	rc := new(usersRowCollector)
	err := r.Query(rc, "SELECT * FROM users")
	return rc.allOrErr(err)
}

type usersRowCollector struct {
	users []entities.User
}

func (rc *usersRowCollector) CollectRow(rs veil.RowScanner) error {
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
		if err == sql.ErrNoRows {
			return nil, repositories.NotFoundError
		}

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
