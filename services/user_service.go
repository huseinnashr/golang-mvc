package services

import "github.com/nvdhunter/golang-mvc/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
