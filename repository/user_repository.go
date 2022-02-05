package repository

import (
	"affordmed/models"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	InsertUser(signup models.User) error
	Get(email string) (*models.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return userRepository{db: db}
}

const (
	createUser = "insert into users(first_name, last_name,email, password)values($1, $2, $3, $4)"
	getUser    = "select * from users where email = $1"
)

func (u userRepository) InsertUser(signup models.User) error {
	result, err := u.db.Exec(createUser, signup.FirstName, signup.LastName, signup.Email, signup.Password)
	if err != nil {
		logrus.Errorf("unable to create user %v", err)
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		logrus.Errorf("unable to retrive affected rows %v", err)
		return err
	}
	if affected == 1 {
		return nil
	}

	return errors.New("something went wrong")
}

func (u userRepository) Get(email string) (*models.User, error) {
	var user models.User
	err := u.db.Get(&user, getUser, email)
	if err != nil {
		logrus.Errorf("unable to fetch user details %v", err)
		return nil, err
	}

	return &user, nil
}
