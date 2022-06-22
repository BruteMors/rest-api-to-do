package repository

import (
	todo "github.com/brutemors/rest-api-to-do"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type Todoitem interface {
}

type Repository struct {
	Authorization
	TodoList
	Todoitem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostrgres(db),
	}
}
