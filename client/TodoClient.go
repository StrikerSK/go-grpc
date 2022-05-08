package client

import (
	"context"
	"github.com/StrikerSK/go-grpc/proto/todo"
	"github.com/StrikerSK/go-grpc/server/Entity"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func CreateTodo(input Entity.TodoStructure) string {
	response, err := GetClient().CreateTodo(context.Background(), input.ConvertToProto())
	if err != nil {
		log.Printf("Error calling method: %v\n", err)
	}

	log.Printf("Server response: %s\n", response.Output)
	return response.Output
}

func ReadTodo(id string) (Entity.TodoStructure, error) {
	response, err := GetClient().ReadTodo(context.Background(), &todo.StringRequest{Input: id})
	if err != nil {
		log.Printf("Error calling method: %v\n", err)
		return Entity.TodoStructure{}, err
	}

	return Entity.ConvertFromProto(response), nil
}

func UpdateTodo(input Entity.TodoStructure) string {
	response, err := GetClient().UpdateTodo(context.Background(), input.ConvertToProto())
	if err != nil {
		log.Printf("Error calling method: %v\n", err)
	}
	return response.Output
}

func DeleteTodo(id string) (string, error) {
	response, err := GetClient().DeleteTodo(context.Background(), &todo.StringRequest{Input: id})
	if err != nil {
		log.Printf("Error calling method: %v\n", err)
		return "", err
	}

	return response.Output, nil
}

func GetTodos() (output []Entity.TodoStructure) {
	response, err := GetClient().FindAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Error calling method: %v\n", err)
	}

	for _, item := range response.Todos {
		output = append(output, Entity.ConvertFromProto(item))
	}

	return
}
