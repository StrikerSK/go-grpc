package ports

import (
	"context"
	todoDomain "github.com/StrikerSK/go-grpc/commons/todo/domain"
)

type ITodoService interface {
	CreateTodo(context.Context, todoDomain.TodoStructure) (string, error)
	GetTodo(context.Context, string) (todoDomain.TodoStructure, error)
	GetTodos(context.Context) ([]todoDomain.TodoStructure, error)
	UpdateTodo(context.Context, todoDomain.TodoStructure) error
	DeleteTodo(context.Context, string) error
}
