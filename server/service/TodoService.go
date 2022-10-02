package service

import (
	"context"
	todoDomain "github.com/StrikerSK/go-grpc/commons/todo/domain"
	todoPorts "github.com/StrikerSK/go-grpc/server/ports"
)

type TodoLocalService struct {
	repository todoPorts.ITodoRepository
}

func NewTodoLocalService(repository todoPorts.ITodoRepository) *TodoLocalService {
	return &TodoLocalService{
		repository: repository,
	}
}

func (r TodoLocalService) CreateTodo(ctx context.Context, todo todoDomain.TodoStructure) (string, error) {
	err := r.repository.CreateTodo(&todo)
	if err != nil {
		return "", err
	}

	return todo.Id, nil
}

func (r TodoLocalService) GetTodo(ctx context.Context, id string) (todoDomain.TodoStructure, error) {
	return r.repository.ReadTodo(id)
}

func (r TodoLocalService) GetTodos(ctx context.Context) ([]todoDomain.TodoStructure, error) {
	return r.repository.ReadTodos(), nil
}

func (r TodoLocalService) UpdateTodo(ctx context.Context, todo todoDomain.TodoStructure) error {
	return r.repository.UpdateTodo(todo)
}

func (r TodoLocalService) DeleteTodo(ctx context.Context, id string) error {
	return r.repository.DeleteTodo(id)
}
