package services

import (
	"github.com/nvdhunter/golang-mvc/domain"
	"github.com/nvdhunter/golang-mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
