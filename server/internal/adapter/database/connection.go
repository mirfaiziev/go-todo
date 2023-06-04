package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDBConn(logger *log.Logger) (*sql.DB, error) {
	// open database
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		return nil, err
	}

	// check db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	logger.Println("Database initialized.")

	return db, nil
}
