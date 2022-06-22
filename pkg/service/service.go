package service

import (
	todo "github.com/brutemors/rest-api-to-do"
	"github.com/brutemors/rest-api-to-do/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type Todoitem interface {
}

type Service struct {
	Authorization
	TodoList
	Todoitem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
