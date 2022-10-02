package service

import (
	"context"
	taskProto "github.com/StrikerSK/go-grpc/commons/proto/task"
	taskDomain "github.com/StrikerSK/go-grpc/commons/task/domain"
	taskPort "github.com/StrikerSK/go-grpc/server/ports"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TaskGrpcService struct {
	service taskPort.ITaskService
	taskProto.UnimplementedTaskServiceServer
}

func NewTaskGrpcService(service taskPort.ITaskService) *TaskGrpcService {
	return &TaskGrpcService{
		service: service,
	}
}

func (s *TaskGrpcService) CreateTask(ctx context.Context, newTask *taskProto.Task) (*taskProto.TaskRequest, error) {
	id, err := s.service.CreateTask(ctx, taskDomain.ConvertFromProto(newTask))
	if err != nil {
		return &taskProto.TaskRequest{}, status.New(500, err.Error()).Err()
	}

	return &taskProto.TaskRequest{Id: id}, nil
}

func (s *TaskGrpcService) ReadTask(ctx context.Context, id *taskProto.TaskRequest) (*taskProto.Task, error) {
	tmpTask, err := s.service.ReadTask(ctx, id.GetId())
	return tmpTask.ConvertToProto(), err
}

func (s *TaskGrpcService) ReadTasks(ctx context.Context, empty *emptypb.Empty) (*taskProto.TaskList, error) {
	var outputSlice []*taskProto.Task
	persistedTasks, err := s.service.ReadTasks(ctx)
	if err != nil {
		return &taskProto.TaskList{}, status.New(500, err.Error()).Err()
	}

	for _, item := range persistedTasks {
		outputSlice = append(outputSlice, item.ConvertToProto())
	}

	return &taskProto.TaskList{
		Tasks: outputSlice,
	}, nil
}

func (s *TaskGrpcService) UpdateTask(ctx context.Context, input *taskProto.Task) (*emptypb.Empty, error) {
	if err := s.service.UpdateTask(ctx, taskDomain.ConvertFromProto(input)); err != nil {
		return &emptypb.Empty{}, status.New(500, err.Error()).Err()
	}

	return &emptypb.Empty{}, nil
}

func (s *TaskGrpcService) DeleteTask(ctx context.Context, input *taskProto.TaskRequest) (*emptypb.Empty, error) {
	if err := s.service.DeleteTask(ctx, input.GetId()); err != nil {
		return &emptypb.Empty{}, status.New(500, err.Error()).Err()
	}

	return &emptypb.Empty{}, nil
}
