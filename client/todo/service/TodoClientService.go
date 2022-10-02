package service

import (
	"context"
	"errors"
	todoProto "github.com/StrikerSK/go-grpc/commons/proto/todo"
	"github.com/StrikerSK/go-grpc/commons/src"
	todoDomain "github.com/StrikerSK/go-grpc/commons/todo/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type TodoClientService struct {
	client todoProto.TodoServiceClient
}

func NewTodoClientService() TodoClientService {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(src.ResolvePortNumber(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	return TodoClientService{
		client: todoProto.NewTodoServiceClient(conn),
	}
}

func (r *TodoClientService) CreateTodo(input todoDomain.TodoStructure) (string, error) {
	response, err := r.client.CreateTodo(context.Background(), input.ConvertToProto())
	if err != nil {
		err = ProcessGrpcError(err)
		log.Printf("Error calling method: %v\n", err)
		return "", err
	}

	todoID := response.GetId()
	log.Printf("Server response: %s\n", todoID)
	return todoID, nil
}

func (r *TodoClientService) ReadTodo(id string) (todoDomain.TodoStructure, error) {
	todoRequest := &todoProto.TodoRequest{
		Id: id,
	}

	response, err := r.client.GetTodo(context.Background(), todoRequest)
	if err != nil {
		err = ProcessGrpcError(err)
		log.Printf("Error calling method: %v\n", err)
		return todoDomain.TodoStructure{}, err
	}

	return todoDomain.ConvertFromProto(response), nil
}

func (r *TodoClientService) UpdateTodo(input todoDomain.TodoStructure) error {
	if _, err := r.client.UpdateTodo(context.Background(), input.ConvertToProto()); err != nil {
		err = ProcessGrpcError(err)
		log.Printf("Error calling method: %v\n", err)
		return ProcessGrpcError(err)
	}

	return nil
}

func (r *TodoClientService) DeleteTodo(id string) error {
	_, err := r.client.DeleteTodo(context.Background(), &todoProto.TodoRequest{Id: id})
	if err != nil {
		err = ProcessGrpcError(err)
		log.Printf("Error calling method: %v\n", err)
		return err
	}

	return nil
}

func (r *TodoClientService) GetTodos() (output []todoDomain.TodoStructure, err error) {
	response, err := r.client.GetTodos(context.Background(), &emptypb.Empty{})
	if err != nil {
		err = ProcessGrpcError(err)
		log.Printf("Error calling method: %v\n", err)
		return nil, ProcessGrpcError(err)
	}

	for _, item := range response.Todos {
		output = append(output, todoDomain.ConvertFromProto(item))
	}

	return
}

func ProcessGrpcError(err error) error {
	statusError, _ := status.FromError(err)
	message := statusError.Message()
	return errors.New(message)
}
