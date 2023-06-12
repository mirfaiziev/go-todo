package repository

import (
	"database/sql"
	"todo-service/internal/application/dto"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) FindAll() ([]*dto.Todo, error) {
	rows, err := r.db.Query(`SELECT * FROM todo`)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var todos []*dto.Todo
	var row dto.Todo
	for rows.Next() {
		rows.Scan(&row.Id, &row.Title, &row.State)
		todos = append(todos, &row)
	}

	return todos, nil
}

func (r *TodoRepository) AddTodo(todo *dto.Todo) (*dto.Todo, error) {
	lastInsertId := 0
	err := r.db.QueryRow(`INSERT INTO todo(title) VALUES($1) RETURNING id`, todo.Title).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	var newTodo dto.Todo
	err = r.db.QueryRow(`SELECT * FROM todo WHERE id=$1`, lastInsertId).Scan(&newTodo.Id, &newTodo.Title, &newTodo.State)
	if err != nil {
		return nil, err
	}

	return &newTodo, nil
}
