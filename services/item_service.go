package services

import (
	"net/http"

	"github.com/nvdhunter/golang-mvc/domain"
	"github.com/nvdhunter/golang-mvc/utils"
)

type itemService struct {
}

var (
	ItemService itemService
)

func (i *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
