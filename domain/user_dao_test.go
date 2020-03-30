package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// initialization

	// Execution
	user, err := GetUser(0)

	// Validation
	assert.Nil(t, user, "we were not expecting user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exists", err.Message)

	// if user != nil {
	// 	t.Error("we were not expecting user with id 0")
	// }
	// if err == nil {
	// 	t.Error("we were expecting an error with user id is 0")
	// }
	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("we're expecting 404 when user not found")
	// }
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "Dedy", user.FirstName)
	assert.EqualValues(t, "Styawan", user.LastName)
	assert.EqualValues(t, "dedy@gmail.com", user.Email)

}
