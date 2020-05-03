package domain

import (
	"fmt"
	"github.com/pmanickam81/golang-microservices/api/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "Hello",
			LastName:  "World",
			Email:     "user@test.com",
		},
	}
	UserDAO userDAOInterface
)

func init(){
	UserDAO = &userDAO{}
}

type userDAOInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDAO struct{}

func (u *userDAO) GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:     fmt.Sprintf("User %v does not exists", userId),
		StatusCode:  http.StatusNotFound,
		Description: "Not Found",
	}

}
