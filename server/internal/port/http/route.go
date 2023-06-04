package http

import (
	"net/http"
	"todo-service/internal/port/http/handler"
)

func initRoutes() {
	http.HandleFunc("/hello", handler.ShowHelloWorld)
}
