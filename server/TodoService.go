package server

import (
	"context"
	"github.com/StrikerSK/go-grpc/proto/todo"
	"github.com/StrikerSK/go-grpc/server/Entity"
	"github.com/StrikerSK/go-grpc/server/Repository"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
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

func (s *TodoServer) FindAll(empty *emptypb.Empty, todoStream todo.TodoService_FindAllServer) error {
	tasks := Repository.GetRepository().FindAll()
	for index := range tasks {
		log.Println(tasks[index])
		if err := todoStream.Send(tasks[index].ConvertToProto()); err != nil {
			return err
		}
	}
	return nil
}
