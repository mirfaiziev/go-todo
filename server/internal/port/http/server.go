package http

import (
	"net/http"
)

func InitServer() error {
	initRoutes()
	return http.ListenAndServe(":8080", nil)
}
