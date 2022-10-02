package domain

import (
	taskProto "github.com/StrikerSK/go-grpc/commons/proto/task"
)

type TaskStructure struct {
	Id          string   `bson:"id" json:"id"`
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Done        bool     `bson:"done" json:"done"`
	Tags        []string `bson:"tags" json:"tags"`
}

func ConvertFromProto(input *taskProto.Task) TaskStructure {
	return TaskStructure{
		Id:          input.Id,
		Name:        input.Name,
		Description: input.Description,
		Done:        input.Done,
		Tags:        input.Tags,
	}
}

func (r *TaskStructure) ConvertToProto() *taskProto.Task {
	return &taskProto.Task{
		Id:          r.Id,
		Name:        r.Name,
		Description: r.Description,
		Done:        r.Done,
		Tags:        r.Tags,
	}
}
