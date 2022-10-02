package ports

import (
	"context"
	taskDomain "github.com/StrikerSK/go-grpc/commons/task/domain"
)

type ITaskService interface {
	CreateTask(context.Context, taskDomain.TaskStructure) (string, error)
	ReadTask(context.Context, string) (taskDomain.TaskStructure, error)
	ReadTasks(context.Context) ([]taskDomain.TaskStructure, error)
	UpdateTask(context.Context, taskDomain.TaskStructure) error
	DeleteTask(context.Context, string) error
}
