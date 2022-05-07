package client

import (
	"context"
	"github.com/StrikerSK/go-grpc/proto/todo"
	"github.com/StrikerSK/go-grpc/server/Entity"
	"github.com/StrikerSK/go-grpc/src"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func CreateClient() todo.TodoServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(src.ResolvePortNumber(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	return todo.NewTodoServiceClient(conn)
}

func CreateTodo(input Entity.TodoStructure) string {
	response, err := CreateClient().CreateTodo(context.Background(), input.ConvertToProto())
	if err != nil {
		log.Fatalf("Error calling method: %v\n", err)
	}

	log.Printf("Server response: %s\n", response.Output)
	return response.Output
}

func ReadTodo(id string) (Entity.TodoStructure, error) {
	response, err := CreateClient().ReadTodo(context.Background(), &todo.StringRequest{Input: id})
	if err != nil {
		log.Fatalf("Error calling method: %v\n", err)
		return Entity.TodoStructure{}, err
	}

	return Entity.ConvertFromProto(response), nil
}

func GetTodos() (output []Entity.TodoStructure) {
	response, err := CreateClient().FindAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error calling method: %v\n", err)
	}

	for _, item := range response.Todos {
		output = append(output, Entity.ConvertFromProto(item))
	}

	return
}
