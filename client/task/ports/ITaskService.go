package ports

import (
	taskDomain "github.com/StrikerSK/go-grpc/commons/task/domain"
)

type ITaskService interface {
	CreateTask(taskDomain.TaskStructure) (string, error)
	ReadTask(string) (taskDomain.TaskStructure, error)
	ReadTasks() ([]taskDomain.TaskStructure, error)
	UpdateTask(taskDomain.TaskStructure) error
	DeleteTask(string) error
}
