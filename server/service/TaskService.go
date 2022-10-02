package service

import (
	"context"
	taskDomain "github.com/StrikerSK/go-grpc/commons/task/domain"
	taskPorts "github.com/StrikerSK/go-grpc/server/ports"
)

type TaskLocalService struct {
	repository taskPorts.ITaskRepository
}

func NewTaskLocalService(repository taskPorts.ITaskRepository) *TaskLocalService {
	return &TaskLocalService{
		repository: repository,
	}
}

func (r TaskLocalService) CreateTask(ctx context.Context, task taskDomain.TaskStructure) (string, error) {
	err := r.repository.CreateTask(&task)
	if err != nil {
		return "", err
	}

	return task.Id, nil
}

func (r TaskLocalService) ReadTask(ctx context.Context, id string) (taskDomain.TaskStructure, error) {
	return r.repository.ReadTask(id)
}

func (r TaskLocalService) ReadTasks(ctx context.Context) ([]taskDomain.TaskStructure, error) {
	return r.repository.ReadTasks(), nil
}

func (r TaskLocalService) UpdateTask(ctx context.Context, task taskDomain.TaskStructure) error {
	return r.repository.UpdateTask(task)
}

func (r TaskLocalService) DeleteTask(ctx context.Context, id string) error {
	return r.repository.DeleteTask(id)
}
