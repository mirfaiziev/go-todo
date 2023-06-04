package main

import (
	"log"
	"todo-service/internal/adapter/database"
	"todo-service/internal/port/http"
)

func main() {
	database.InitConnection()

	if err := http.InitServer(); err != nil {
		log.Fatal("cannot init http server")
	}
}
