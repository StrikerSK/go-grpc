package service

import (
	"context"
	"errors"
	taskProto "github.com/StrikerSK/go-grpc/commons/proto/task"
	"github.com/StrikerSK/go-grpc/commons/src"
	taskDomain "github.com/StrikerSK/go-grpc/commons/task/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type TaskClientService struct {
	client taskProto.TaskServiceClient
}

func NewTaskClientService() TaskClientService {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(src.ResolvePortNumber(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	return TaskClientService{
		client: taskProto.NewTaskServiceClient(conn),
	}
}

func (r *TaskClientService) CreateTask(input taskDomain.TaskStructure) (string, error) {
	response, err := r.client.CreateTask(context.Background(), input.ConvertToProto())
	if err != nil {
		err = ProcessGrpcError(err)
		log.Printf("Error calling method: %v\n", err)
		return "", err
	}

	id := response.GetId()
	log.Printf("Server response: %s\n", id)
	return id, nil
}

func (r *TaskClientService) ReadTask(id string) (taskDomain.TaskStructure, error) {
	taskRequest := &taskProto.TaskRequest{
		Id: id,
	}

	response, err := r.client.ReadTask(context.Background(), taskRequest)
	if err != nil {
		err = ProcessGrpcError(err)
		log.Printf("Error calling method: %v\n", err)
		return taskDomain.TaskStructure{}, err
	}

	return taskDomain.ConvertFromProto(response), nil
}

func (r *TaskClientService) ReadTasks() (output []taskDomain.TaskStructure, err error) {
	response, err := r.client.ReadTasks(context.Background(), &emptypb.Empty{})
	if err != nil {
		err = ProcessGrpcError(err)
		log.Printf("Error calling method: %v\n", err)
		return nil, err
	}

	for _, item := range response.Tasks {
		output = append(output, taskDomain.ConvertFromProto(item))
	}

	return
}

func (r *TaskClientService) UpdateTask(input taskDomain.TaskStructure) error {
	if _, err := r.client.UpdateTask(context.Background(), input.ConvertToProto()); err != nil {
		err = ProcessGrpcError(err)
		log.Printf("Error calling method: %v\n", err)
		return ProcessGrpcError(err)
	}

	return nil
}

func (r *TaskClientService) DeleteTask(id string) error {
	_, err := r.client.DeleteTask(context.Background(), &taskProto.TaskRequest{Id: id})
	if err != nil {
		err = ProcessGrpcError(err)
		log.Printf("Error calling method: %v\n", err)
		return err
	}

	return nil
}

func ProcessGrpcError(err error) error {
	statusError, _ := status.FromError(err)
	message := statusError.Message()
	return errors.New(message)
}
