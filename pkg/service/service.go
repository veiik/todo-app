package service

import "todo/pkg/repository"

type TodoList interface {
}

type TodoItem interface {
}

type Authorisation interface {
}

type Service struct {
	TodoItem
	TodoList
	Authorisation
	bd *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{bd: repos}
}
