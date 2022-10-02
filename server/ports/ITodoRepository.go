package ports

import (
	todoDomain "github.com/StrikerSK/go-grpc/commons/todo/domain"
)

type ITodoRepository interface {
	CreateTodo(*todoDomain.TodoStructure) error
	ReadTodo(string) (todoDomain.TodoStructure, error)
	ReadTodos() []todoDomain.TodoStructure
	UpdateTodo(todoDomain.TodoStructure) error
	DeleteTodo(string) error
}
