package service

import (
	"context"
	todoDomain "github.com/StrikerSK/go-grpc/commons/todo/domain"
	"github.com/StrikerSK/go-grpc/proto/todo"
	todoPort "github.com/StrikerSK/go-grpc/server/ports"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoService struct {
	repository todoPort.ITodoRepository
	todo.UnimplementedTodoServiceServer
}

func NewTodoService(repository todoPort.ITodoRepository) *TodoService {
	return &TodoService{
		repository: repository,
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, newTodo *todo.Todo) (*todo.TodoRequest, error) {
	customID := uuid.NewString()
	_ = s.repository.CreateTodo(todoDomain.ConvertFromProto(newTodo))

	return &todo.TodoRequest{
		Id: customID,
	}, nil
}

func (s *TodoService) GetTodo(ctx context.Context, id *todo.TodoRequest) (*todo.Todo, error) {
	tmpTodo, err := s.repository.ReadTodo(id.GetId())
	return tmpTodo.ConvertToProto(), err
}

func (s *TodoService) GetTodos(context.Context, *emptypb.Empty) (*todo.TodoList, error) {
	var outputSlice []*todo.Todo
	for _, item := range s.repository.FindAll() {
		outputSlice = append(outputSlice, item.ConvertToProto())
	}

	return &todo.TodoList{
		Todos: outputSlice,
	}, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, input *todo.Todo) (*emptypb.Empty, error) {
	if err := s.repository.UpdateTodo(todoDomain.ConvertFromProto(input)); err != nil {
		return &emptypb.Empty{}, status.New(500, err.Error()).Err()
	}

	return &emptypb.Empty{}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, input *todo.TodoRequest) (*emptypb.Empty, error) {
	if err := s.repository.DeleteTodo(input.GetId()); err != nil {
		return &emptypb.Empty{}, status.New(500, err.Error()).Err()
	}

	return &emptypb.Empty{}, nil
}
