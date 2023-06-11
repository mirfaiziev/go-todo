package repository

import (
	"database/sql"
	"todo-service/internal/application/object"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) FindAll() ([]*object.Todo, error) {
	rows, err := r.db.Query(`SELECT * FROM todo`)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var todos []*object.Todo
	var row object.Todo
	for rows.Next() {
		rows.Scan(&row.Id, &row.Title, &row.State)
		todos = append(todos, &row)
	}

	return todos, nil
}

func (r *TodoRepository) AddTodo(todo *object.Todo) (*object.Todo, error) {
	lastInsertId := 0
	err := r.db.QueryRow(`INSERT INTO todo(title) VALUES($1) RETURNING id`, todo.Title).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	var newTodo object.Todo
	err = r.db.QueryRow(`SELECT * FROM todo WHERE id=$1`, lastInsertId).Scan(&newTodo.Id, &newTodo.Title, &newTodo.State)
	if err != nil {
		return nil, err
	}

	return &newTodo, nil
}
