package services

import (
	"github.com/nvdhunter/golang-mvc/domain"
	"github.com/nvdhunter/golang-mvc/utils"
)

type userService struct {
}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
