package service

import (
	"database/sql"
	"log"
	"todo-service/internal/adapter/database/repository"
	"todo-service/internal/application/dto"
)

type TodoService struct {
	todoRepo *repository.TodoRepository
}

func NewTodoService(logger *log.Logger, db *sql.DB) *TodoService {
	return &TodoService{
		todoRepo: repository.NewTodoRepository(db),
	}
}

func (t *TodoService) AddTodo(todo *dto.Todo) (*dto.Todo, error) {
	return t.todoRepo.AddTodo(todo)
}

func (t *TodoService) GetAll() ([]*dto.Todo, error) {
	return t.todoRepo.FindAll()
}
