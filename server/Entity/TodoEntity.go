package Entity

import "github.com/StrikerSK/go-grpc/proto/todo"

type TodoStructure struct {
	Id          string `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Done        bool   `bson:"done" json:"done"`
}

func ConvertFromProto(input *todo.CustomTodo) TodoStructure {
	return TodoStructure{
		Id:          input.Id,
		Name:        input.Name,
		Description: input.Description,
		Done:        input.Done,
	}
}

func (r *TodoStructure) ConvertToProto() *todo.CustomTodo {
	return &todo.CustomTodo{
		Id:          r.Id,
		Name:        r.Name,
		Description: r.Description,
		Done:        r.Done,
	}
}
