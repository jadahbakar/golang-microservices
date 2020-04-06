package services

import (
	"net/http"

	"github.com/jadahbakar/golang-microservices/domain"
	"github.com/jadahbakar/golang-microservices/utils"
)

func GetItem(itemID string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}

}
