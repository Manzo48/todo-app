package service

import (
	"github.com/Manzo48/todo-app"
	"github.com/Manzo48/todo-app/pkg/repository"
)

// import "github.com/Manzo48/todo-app/pkg/repository"

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetAllUsers() ([]todo.User, error) 
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}