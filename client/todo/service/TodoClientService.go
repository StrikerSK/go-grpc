package service

import (
	"context"
	"github.com/StrikerSK/go-grpc/proto/todo"
	"github.com/StrikerSK/go-grpc/server/Entity"
	"github.com/StrikerSK/go-grpc/src"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type TodoClientService struct {
	client todo.TodoServiceClient
}

func NewTodoClientService() TodoClientService {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(src.ResolvePortNumber(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	return TodoClientService{
		client: todo.NewTodoServiceClient(conn),
	}
}

func (r *TodoClientService) CreateTodo(input Entity.TodoStructure) string {
	response, err := r.client.CreateTodo(context.Background(), input.ConvertToProto())
	if err != nil {
		log.Printf("Error calling method: %v\n", err)
	}

	log.Printf("Server response: %s\n", response.Output)
	return response.Output
}

func (r *TodoClientService) ReadTodo(id string) (Entity.TodoStructure, error) {
	response, err := r.client.ReadTodo(context.Background(), &todo.StringRequest{Input: id})
	if err != nil {
		log.Printf("Error calling method: %v\n", err)
		return Entity.TodoStructure{}, err
	}

	return Entity.ConvertFromProto(response), nil
}

func (r *TodoClientService) UpdateTodo(input Entity.TodoStructure) string {
	response, err := r.client.UpdateTodo(context.Background(), input.ConvertToProto())
	if err != nil {
		log.Printf("Error calling method: %v\n", err)
	}
	return response.Output
}

func (r *TodoClientService) DeleteTodo(id string) (string, error) {
	response, err := r.client.DeleteTodo(context.Background(), &todo.StringRequest{Input: id})
	if err != nil {
		log.Printf("Error calling method: %v\n", err)
		return "", err
	}

	return response.Output, nil
}

func (r *TodoClientService) GetTodos() (output []Entity.TodoStructure) {
	response, err := r.client.FindAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Error calling method: %v\n", err)
	}

	for _, item := range response.Todos {
		output = append(output, Entity.ConvertFromProto(item))
	}

	return
}
