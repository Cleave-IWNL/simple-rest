package service

import "github.com/Cleave-IWNL/simpleRest/pkg/repository"

type Authorisation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorisation
	TodoItem
	TodoList
}

func NewService(repos repository.Repository) *Service {
	return &Service{}
}
