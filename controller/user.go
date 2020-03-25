package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jadahbakar/golang-microservices/utils"

	"github.com/jadahbakar/golang-microservices/services"
)

// GetUser function
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIDParm := req.URL.Query().Get("userid")
	log.Printf("request user_id %v", userIDParm)
	userID, err := strconv.ParseInt(userIDParm, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "User ID must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		// resp.WriteHeader(http.StatusBadRequest)
		// resp.Write([]byte("User ID must be a number"))
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		// resp.WriteHeader(http.StatusNotFound)
		// resp.Write([]byte(err.Error()))
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
