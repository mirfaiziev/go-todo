package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"todo-service/internal/application/dto"
	"todo-service/internal/application/service"
)

type TodoHandler struct {
	db      *sql.DB
	service *service.TodoService
	logger  *log.Logger
}

func NewTodoHandler(logger *log.Logger, db *sql.DB) *TodoHandler {
	return &TodoHandler{service: service.NewTodoService(logger, db)}
}

func (td *TodoHandler) ShowList(rw http.ResponseWriter, req *http.Request) {
	todos, _ := td.service.GetAll()

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	jsonResp, err := json.Marshal(todos)
	if err != nil {
		td.logger.Fatalf("error parsing json:", err)
	}
	rw.Write(jsonResp)
}

func (td *TodoHandler) AddTodo(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var todo dto.Todo
	err := decoder.Decode(&todo)
	if err != nil {
		td.logger.Fatalf("error parsing json:", err)
	}

	newTodo, err := td.service.AddTodo(&todo)
	if err != nil {
		td.logger.Fatalf("error parsing json:", err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	jsonResp, err := json.Marshal(newTodo)
	if err != nil {
		log.Fatalf("error parsing json:", err)
	}
	rw.Write(jsonResp)
}

func (td *TodoHandler) GetTodo(rw http.ResponseWriter, req *http.Request) {

}

func (td *TodoHandler) UpdateTodo(id int) {

}
