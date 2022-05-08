package Repository

import (
	"github.com/StrikerSK/go-grpc/server/Entity"
	"log"
)

type LocalTodoRepository struct{}

func SetLocalRepository() {
	SetMainRepository(&LocalTodoRepository{})
}

func (r *LocalTodoRepository) ReadTodo(ID string) (outputResult Entity.TodoStructure, err error) {
	log.Printf("User provided ID to read: %s\n", ID)
	return Entity.TodoStructure{
		Id:          "123",
		Name:        "MainTask",
		Description: "This represents main task",
		Done:        false,
	}, nil
}

func (r *LocalTodoRepository) FindAll() []Entity.TodoStructure {
	return []Entity.TodoStructure{
		{
			Id:          "123",
			Name:        "MainTask",
			Description: "This represents task 1",
			Done:        false,
		},
		{
			Id:          "123",
			Name:        "MainTask",
			Description: "This represents task 2",
			Done:        false,
		},
	}
}

func (r *LocalTodoRepository) CreateTodo(inputTask Entity.TodoStructure) (err error) {
	log.Println("User provide new Task input: ", inputTask)
	return
}

func (r *LocalTodoRepository) UpdateTodo(inputTask Entity.TodoStructure) (err error) {
	log.Println("User provide updated Task input for ID [", inputTask.Id, "]: ", inputTask)
	return
}

func (r *LocalTodoRepository) DeleteTodo(ID string) (err error) {
	log.Printf("User provided ID to delete: %s\n", ID)
	return nil
}
