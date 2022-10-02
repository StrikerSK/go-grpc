package ports

import (
	"github.com/StrikerSK/go-grpc/server/Entity"
)

type ITaskService interface {
	CreateTodo(Entity.TodoStructure) string
	ReadTodo(string) (Entity.TodoStructure, error)
	UpdateTodo(Entity.TodoStructure) string
	DeleteTodo(string) (string, error)
	GetTodos() []Entity.TodoStructure
}
