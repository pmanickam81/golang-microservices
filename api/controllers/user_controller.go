package controllers

import (
	"encoding/json"
	"github.com/pmanickam81/golang-microservices/api/services"
	"github.com/pmanickam81/golang-microservices/api/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, r *http.Request){
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"),10, 64)
	if err !=nil {
		apiErr := &utils.ApplicationError{
			Message:    "user Id must be a number",
			StatusCode: http.StatusBadRequest,
			Description:       "bad_request",
		}
		errorJson,_ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		_, _ = resp.Write(errorJson)
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr!=nil {
		errorJson,_ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		_,_ = resp.Write(errorJson)
		return
	}
	userJson, _:= json.Marshal(user)
	_, _ = resp.Write(userJson)
}