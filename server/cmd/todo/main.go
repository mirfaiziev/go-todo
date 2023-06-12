package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-service/internal/adapter/database"
	httpserver "todo-service/internal/port/http"
)

func main() {
	// logger := log.New(os.Stderr, "", 0)
	// db, err := database.NewDBConn(logger)

	// if err != nil {
	// 	panic(err)
	// }
	// // close database
	// defer func() {
	// 	logger.Print("Closing database.")
	// 	db.Close()
	// }()

	// if err := http.InitServer(logger, db); err != nil {
	// 	logger.Fatal("Cannot init http server")
	// }

	// init loger
	logger := log.New(os.Stderr, "", 0)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := runServer(logger, ctx); err != nil {
		logger.Fatal(err)
	}
}

func runServer(logger *log.Logger, ctx context.Context) error {
	// init DB
	db, err := database.NewDBConn()
	defer func() {
		if err := db.Close(); err != nil {
			logger.Fatal("failed to close connection to the database:", err)
		}
	}()

	if err != nil {
		panic(err)
	}

	logger.Println("Database initialized.")

	// init http server
	serverErrors := make(chan error, 1)
	httpSrv := httpserver.NewServer()
	go func() {
		serverErrors <- httpSrv.ListenAndServe()
	}()
	logger.Printf("Start listening %s\n", httpSrv.Addr)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case <-ctx.Done():
		fmt.Errorf("server shutdown: %w", ctx.Err())
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if err := httpSrv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shutdown: %w", err)
		}

		return fmt.Errorf("Successfuly shutdown")
	}
}
