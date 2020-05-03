package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we are not expecting user with  id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "Not Found", err.Description)
	assert.EqualValues(t, "User 0 does not exists", err.Message)
}

func TestGetUser(t *testing.T) {
	user, err := GetUser(123)

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Hello", user.FirstName)
	assert.EqualValues(t, "World", user.LastName)
	assert.EqualValues(t, "user@test.com", user.Email)
}