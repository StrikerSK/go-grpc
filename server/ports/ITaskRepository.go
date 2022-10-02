package ports

import (
	taskDomain "github.com/StrikerSK/go-grpc/commons/task/domain"
)

type ITaskRepository interface {
	CreateTask(*taskDomain.TaskStructure) error
	ReadTask(string) (taskDomain.TaskStructure, error)
	ReadTasks() []taskDomain.TaskStructure
	UpdateTask(taskDomain.TaskStructure) error
	DeleteTask(string) error
}
