package services

import (
	"github.com/pmanickam81/golang-microservices/api/domain"
	"github.com/pmanickam81/golang-microservices/api/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	_               userDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDAO = &userDaoMock{}
}

type userDaoMock struct{}

func (m *userDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestUserService_GetUser(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 does not exists",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}