package http

import (
	"database/sql"
	"log"
	"net/http"
)

func InitServer(logger *log.Logger, db *sql.DB) error {
	return http.ListenAndServe(":8080", http.HandlerFunc(router(db)))
}

