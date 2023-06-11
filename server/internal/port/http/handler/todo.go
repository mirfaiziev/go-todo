package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"todo-service/internal/application/object"
	"todo-service/internal/application/service"
)

type TodoHandler struct {
	HttpHandler
	service *service.TodoService
}

func NewTodoHandler(db *sql.DB) *TodoHandler {
	return &TodoHandler{service: service.NewTodoService(db)}
}

func (td *TodoHandler) ShowList(rw http.ResponseWriter, req *http.Request) {
	todos, _ := td.service.GetAll()
	
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	jsonResp, err := json.Marshal(todos)
	if err != nil {
		log.Fatalf("error parsing json:", err)
	}
	rw.Write(jsonResp)
}

func (td *TodoHandler) AddTodo(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var todo object.Todo
	err := decoder.Decode(&todo)
	if err != nil {
		log.Fatalf("error parsing json:", err)
	}

	newTodo, err := td.service.AddTodo(&todo)
	if err != nil {
		log.Fatalf("error parsing json:", err)
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
