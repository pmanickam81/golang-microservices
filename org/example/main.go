package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Welcome to Golang Microservices!!!"))
	})

	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		panic(err)
	}
}