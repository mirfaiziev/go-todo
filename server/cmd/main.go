package main

import (
	"log"
	"os"
	"todo-service/internal/adapter/database"
	"todo-service/internal/port/http"
)

func main() {
	logger := log.New(os.Stderr, "", 0)
	db, err := database.NewDBConn(logger)

	if err != nil {
		panic(err)
	}
	// close database
	defer func() {
		logger.Print("Closing database.")
		db.Close()
	}()

	if err := http.InitServer(logger, db); err != nil {
		logger.Fatal("Cannot init http server")
	}
}
