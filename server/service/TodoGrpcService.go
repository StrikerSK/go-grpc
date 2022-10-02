package service

import (
	"context"
	todoProto "github.com/StrikerSK/go-grpc/commons/proto/todo"
	todoDomain "github.com/StrikerSK/go-grpc/commons/todo/domain"
	todoPort "github.com/StrikerSK/go-grpc/server/ports"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoGrpcService struct {
	service todoPort.ITodoService
	todoProto.UnimplementedTodoServiceServer
}

func NewTodoGrpcService(service todoPort.ITodoService) *TodoGrpcService {
	return &TodoGrpcService{
		service: service,
	}
}

func (s *TodoGrpcService) CreateTodo(ctx context.Context, newTodo *todoProto.Todo) (*todoProto.TodoRequest, error) {
	id, err := s.service.CreateTodo(ctx, todoDomain.ConvertFromProto(newTodo))
	if err != nil {
		return &todoProto.TodoRequest{}, status.New(500, err.Error()).Err()
	}

	return &todoProto.TodoRequest{Id: id}, nil
}

func (s *TodoGrpcService) GetTodo(ctx context.Context, id *todoProto.TodoRequest) (*todoProto.Todo, error) {
	tmpTodo, err := s.service.GetTodo(ctx, id.GetId())
	return tmpTodo.ConvertToProto(), err
}

func (s *TodoGrpcService) GetTodos(ctx context.Context, empty *emptypb.Empty) (*todoProto.TodoList, error) {
	var outputSlice []*todoProto.Todo
	persistedTodos, err := s.service.GetTodos(ctx)
	if err != nil {
		return &todoProto.TodoList{}, status.New(500, err.Error()).Err()
	}

	for _, item := range persistedTodos {
		outputSlice = append(outputSlice, item.ConvertToProto())
	}

	return &todoProto.TodoList{
		Todos: outputSlice,
	}, nil
}

func (s *TodoGrpcService) UpdateTodo(ctx context.Context, input *todoProto.Todo) (*emptypb.Empty, error) {
	if err := s.service.UpdateTodo(ctx, todoDomain.ConvertFromProto(input)); err != nil {
		return &emptypb.Empty{}, status.New(500, err.Error()).Err()
	}

	return &emptypb.Empty{}, nil
}

func (s *TodoGrpcService) DeleteTodo(ctx context.Context, input *todoProto.TodoRequest) (*emptypb.Empty, error) {
	if err := s.service.DeleteTodo(ctx, input.GetId()); err != nil {
		return &emptypb.Empty{}, status.New(500, err.Error()).Err()
	}

	return &emptypb.Empty{}, nil
}
