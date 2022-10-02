package ports

import (
	todoDomain "github.com/StrikerSK/go-grpc/commons/todo/domain"
)

type ITodoRepository interface {
	FindAll() []todoDomain.TodoStructure
	CreateTodo(todoDomain.TodoStructure) error
	ReadTodo(string) (todoDomain.TodoStructure, error)
	UpdateTodo(todoDomain.TodoStructure) error
	DeleteTodo(string) error
}
