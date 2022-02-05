package service

import (
	"affordmed/models"
	"affordmed/repository"
	"affordmed/utils"
	"errors"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	InsertUser(signup models.User) error
	Login(login models.Login) error
}

type userService struct {
	userRepository repository.UserRepository
}

func (u userService) InsertUser(signup models.User) error {
	password, err := utils.HashPassword(signup.Password)
	if err != nil {
		logrus.Errorf("unable to hash the password %v", err)
		return err
	}
	signup.Password = password
	return u.userRepository.InsertUser(signup)
}

func (u userService) Login(login models.Login) error {
	user, err := u.userRepository.Get(login.Email)
	if err != nil {
		logrus.Errorf("unable to get user details %v", err)
		return err
	}

	if utils.CheckPasswordHash(login.Password, user.Password) {
		return nil
	}

	return errors.New("email/password invalid")
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return userService{userRepository: userRepository}
}
