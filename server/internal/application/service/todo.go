package service

import (
	"database/sql"
	"todo-service/internal/adapter/database/repository"
	"todo-service/internal/application/object"
)

type TodoService struct {
	todoRepo *repository.TodoRepository
}

func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{
		todoRepo: repository.NewTodoRepository(db),
	}
}

func (t *TodoService) AddTodo(todo *object.Todo) (*object.Todo, error) {
	return t.todoRepo.AddTodo(todo)
}

func (t *TodoService) GetAll() ([]*object.Todo, error) {
	return t.todoRepo.FindAll()
}
