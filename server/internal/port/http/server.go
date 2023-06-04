package http

import (
	"database/sql"
	"log"
	"net/http"
	"todo-service/internal/port/http/handler"
)

func InitServer(logger *log.Logger, db *sql.DB) error {
	initRoutes(db)
	return http.ListenAndServe(":8080", nil)
}

func initRoutes(db *sql.DB) {
	httpHandler := handler.NewHttpHandler(db)
	http.HandleFunc("/hello", httpHandler.ShowHelloWorld)
}
