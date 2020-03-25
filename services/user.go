package services

import (
	"github.com/jadahbakar/golang-microservices/domain"
	"github.com/jadahbakar/golang-microservices/utils"
)

// GetUser Domain
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
