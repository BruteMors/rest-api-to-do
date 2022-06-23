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
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
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
		TodoList:      NewTodoListPostgres(db),
	}
}
