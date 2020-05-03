package services

import (
	"github.com/pmanickam81/golang-microservices/api/domain"
	"github.com/pmanickam81/golang-microservices/api/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError){
	return domain.GetUser(userId)
}
