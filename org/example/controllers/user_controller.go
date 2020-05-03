package controllers

import (
	"encoding/json"
	"github.com/pmanickam81/golang-microservices/org/example/services"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, r *http.Request){
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"),10, 64)
	if err !=nil {
		resp.WriteHeader(http.StatusNotFound)
		_, _ = resp.Write([]byte("user id must be a number"))
		return
	}
	user, err := services.GetUser(userId)
	if err!=nil {
		resp.WriteHeader(http.StatusNotFound)
		_,_ = resp.Write([]byte(err.Error()))
		return
	}
	userJson, _:= json.Marshal(user)
	_, _ = resp.Write(userJson)
}