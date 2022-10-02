package ports

import (
	todoDomain "github.com/StrikerSK/go-grpc/commons/todo/domain"
)

type ITaskService interface {
	CreateTodo(todoDomain.TodoStructure) (string, error)
	ReadTodo(string) (todoDomain.TodoStructure, error)
	UpdateTodo(todoDomain.TodoStructure) error
	DeleteTodo(string) error
	GetTodos() ([]todoDomain.TodoStructure, error)
}
