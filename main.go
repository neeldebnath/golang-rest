package main

import (
	"net/http"

	handlersusers "github.com/neeldebnath/golang-rest/handlers/users"
)

func main() {
	http.HandleFunc("/user/login", handlersusers.Login)
}