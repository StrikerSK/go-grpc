package domain

import (
	"github.com/StrikerSK/go-grpc/commons/proto/todo"
)

type TodoStructure struct {
	Id          string   `bson:"id" json:"id"`
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Done        bool     `bson:"done" json:"done"`
	Tags        []string `bson:"tags" json:"tags"`
}

func ConvertFromProto(input *todo.Todo) TodoStructure {
	return TodoStructure{
		Id:          input.Id,
		Name:        input.Name,
		Description: input.Description,
		Done:        input.Done,
		Tags:        input.Tags,
	}
}

func (r *TodoStructure) ConvertToProto() *todo.Todo {
	return &todo.Todo{
		Id:          r.Id,
		Name:        r.Name,
		Description: r.Description,
		Done:        r.Done,
		Tags:        r.Tags,
	}
}
