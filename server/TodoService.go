package server

import (
	"context"
	"errors"
	"github.com/StrikerSK/go-grpc/proto/todo"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type TodoServer struct {
	todo.UnimplementedTodoServiceServer
	todos []todo.PersistedTodo
}

func (s *TodoServer) CreateTodo(ctx context.Context, newTodo *todo.NewTodo) (*todo.IdRequest, error) {
	customID := uuid.NewString()

	s.todos = append(s.todos, todo.PersistedTodo{
		Id:          customID,
		Name:        newTodo.Name,
		Description: newTodo.Description,
		Done:        newTodo.Done,
	})

	log.Printf("New Todo created: %s\n", customID)
	return &todo.IdRequest{Id: customID}, nil
}

func (s *TodoServer) GetTodo(ctx context.Context, id *todo.IdRequest) (*todo.PersistedTodo, error) {
	for index := range s.todos {
		if s.todos[index].Id == id.GetId() {
			return &s.todos[index], nil
		}
	}

	return nil, errors.New("item not found")
}

func (s *TodoServer) GetTodos(empty *emptypb.Empty, todoStream todo.TodoService_GetTodosServer) error {
	for index := range s.todos {
		if err := todoStream.Send(&s.todos[index]); err != nil {
			return err
		}
	}

	return nil
}
