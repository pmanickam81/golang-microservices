package services

import "github.com/pmanickam81/golang-microservices/org/example/domain"

func GetUser(userId int64) (*domain.User, error){
	return domain.GetUser(userId)
}
