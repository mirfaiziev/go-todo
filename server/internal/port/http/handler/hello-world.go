package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

type HttpHandler struct {
	db *sql.DB
}

func NewHttpHandler(db *sql.DB) *HttpHandler {
	return &HttpHandler{
		db: db,
	}
}

func (h *HttpHandler) ShowHelloWorld(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello World\n")
}
