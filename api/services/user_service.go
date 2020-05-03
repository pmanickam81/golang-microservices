package services

import (
	"github.com/pmanickam81/golang-microservices/api/domain"
	"github.com/pmanickam81/golang-microservices/api/utils"
)

var (
	UserService userService
)

type userService struct{}

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDAO.GetUser(userId)
}
