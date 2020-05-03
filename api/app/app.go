package app

import (
	"github.com/pmanickam81/golang-microservices/api/controllers"
	"net/http"
)

func StartApp(){
	http.HandleFunc("/user", controllers.GetUser)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		panic(err)
	}
}
