package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/klevtcov/todo"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Autorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
	}
}
