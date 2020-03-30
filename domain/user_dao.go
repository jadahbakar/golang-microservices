package domain

import (
	"fmt"
	"net/http"

	"github.com/jadahbakar/golang-microservices/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 1, FirstName: "Dedy", LastName: "Styawan", Email: "dedy@gmail.com"},
	}
)

// GetUser Domain
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	// user := users[userID]
	// if user == nil {
	// 	return nil, errors.New(fmt.Sprintf("user %v was not found", userID))
	// }
	if user := users[userID]; user != nil {
		return user, nil
	}
	// return nil, errors.New(fmt.Sprintf("user %v was not found", userID))
	// return nil, fmt.Errorf("user %v was not found", userID)
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exists", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
