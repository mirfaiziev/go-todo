package http

import (
	"database/sql"
	"log"
	"net/http"
	"todo-service/internal/port/http/handler"
)

func InitServer(logger *log.Logger, db *sql.DB) error {
	return http.ListenAndServe(":8080", http.HandlerFunc(router(logger, db)))
}

func NewServer() *http.Server {
	return &http.Server{
		// todo move port address to env
		Addr:    "0.0.0.0:8080",
		Handler: http.HandlerFunc(handler.NewHelloWorldHandler().ShowHelloWorld),
	}
}
