package Repository

import (
	todoDomain "github.com/StrikerSK/go-grpc/commons/todo/domain"
	"log"
)

type LocalTodoRepository struct{}

func NewLocalTodoRepository() LocalTodoRepository {
	return LocalTodoRepository{}
}

func (r *LocalTodoRepository) ReadTodo(ID string) (outputResult todoDomain.TodoStructure, err error) {
	log.Printf("User provided ID to read: %s\n", ID)
	return todoDomain.TodoStructure{
		Id:          "123",
		Name:        "MainTask",
		Description: "This represents main task",
		Done:        false,
		Tags: []string{
			"tag1",
			"tag2",
		},
	}, nil
}

func (r *LocalTodoRepository) FindAll() []todoDomain.TodoStructure {
	return []todoDomain.TodoStructure{
		{
			Id:          "123",
			Name:        "MainTask",
			Description: "This represents task 1",
			Done:        false,
			Tags: []string{
				"tag1",
				"tag2",
			},
		},
		{
			Id:          "123",
			Name:        "MainTask",
			Description: "This represents task 2",
			Done:        false,
			Tags: []string{
				"tag1",
				"tag2",
			},
		},
	}
}

func (r *LocalTodoRepository) CreateTodo(inputTask todoDomain.TodoStructure) (err error) {
	log.Println("User provide new Task input: ", inputTask)
	return
}

func (r *LocalTodoRepository) UpdateTodo(inputTask todoDomain.TodoStructure) (err error) {
	log.Println("User provide updated Task input for ID [", inputTask.Id, "]: ", inputTask)
	return
}

func (r *LocalTodoRepository) DeleteTodo(ID string) (err error) {
	log.Printf("User provided ID to delete: %s\n", ID)
	return nil
}
