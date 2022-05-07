package server

import (
	"context"
	"github.com/StrikerSK/go-grpc/proto/todo"
	"github.com/StrikerSK/go-grpc/server/Entity"
	"github.com/StrikerSK/go-grpc/server/Repository"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoServer struct {
	todo.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, newTodo *todo.CustomTodo) (*todo.StringResponse, error) {
	customID := uuid.NewString()
	_ = Repository.GetRepository().CreateTodo(Entity.ConvertFromProto(newTodo))
	return &todo.StringResponse{Output: customID}, nil
}

func (s *TodoServer) ReadTodo(ctx context.Context, id *todo.StringRequest) (*todo.CustomTodo, error) {
	tmpTodo, err := Repository.GetRepository().ReadTodo(id.GetInput())
	return tmpTodo.ConvertToProto(), err
}

func (s *TodoServer) FindAll(context.Context, *emptypb.Empty) (*todo.TodoArray, error) {
	var outputSlice []*todo.CustomTodo
	for _, item := range Repository.GetRepository().FindAll() {
		outputSlice = append(outputSlice, item.ConvertToProto())
	}
	return &todo.TodoArray{Todos: outputSlice}, nil
}
