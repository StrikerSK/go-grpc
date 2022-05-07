package client

import (
	"context"
	"github.com/StrikerSK/go-grpc/proto/todo"
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

func CreateTodo() string {
	customTodo := todo.CustomTodo{
		Name:        "First Todo",
		Description: "Created First Todo",
		Done:        false,
	}

	response, err := CreateClient().CreateTodo(context.Background(), &customTodo)
	if err != nil {
		log.Fatalf("Error calling method: %v\n", err)
	}

	log.Printf("Server response: %s\n", response.Id)
	return response.Id
}

func GetTodo(id string) {
	response, err := CreateClient().GetTodo(context.Background(), &todo.IdRequest{Id: id})
	if err != nil {
		log.Fatalf("Error calling method: %v\n", err)
	}

	log.Println(response)
}

func GetTodos() {
	response, err := CreateClient().GetTodos(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error calling method: %v\n", err)
	}

	log.Println(response)
}
