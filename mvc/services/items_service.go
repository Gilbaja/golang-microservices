package services

import (
	"github.com/Gilbaja/golang-microservices/mvc/domain"
	"github.com/Gilbaja/golang-microservices/mvc/utils"
	"net/http"
)

type ItemService struct {
}

//var (
//	ItemService itemService
//)

func (s *ItemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message: "implement me",
		Status:  http.StatusInternalServerError,
		Code:    "internal_server_error",
	}
}