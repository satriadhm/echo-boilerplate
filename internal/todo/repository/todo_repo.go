package repository

import (
	"database/sql"
	"errors"

	"github.com/satriadhm-echo-boilerplate/internal/todo"
)

type TodoRepository interface {
	Create(todo *todo.Todo) error
	FindById(id int) (*todo.Todo, error)
	Update(todo *todo.Todo) error
	Delete(id int) error
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (repo *todoRepository) Create(todo *todo.Todo) error {
	stmt, err := repo.db.Prepare("INSERT INTO todos (name, is_done) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(todo.Name, todo.IsDone)
	return err
}

func (repo *todoRepository) FindById(id int) (*todo.Todo, error) {
	row := repo.db.QueryRow("SELECT id, name, is_done FROM todos WHERE id = ?", id)
	todo := &todo.Todo{}
	if err := row.Scan(&todo.ID, &todo.Name, &todo.IsDone); err != nil {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}

func (repo *todoRepository) Update(todo *todo.Todo) error {
	stmt, err := repo.db.Prepare("UPDATE todos SET name = ?, is_done = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(todo.Name, todo.IsDone, todo.ID)
	return err
}

func (repo *todoRepository) Delete(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
