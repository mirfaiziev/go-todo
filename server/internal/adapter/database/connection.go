package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func NewDBConn() (*sql.DB, error) {
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

	return db, nil
}
