package Repository

import (
	taskDomain "github.com/StrikerSK/go-grpc/commons/task/domain"
	"github.com/google/uuid"
	"log"
)

type LocalTaskRepository struct{}

func NewLocalTaskRepository() LocalTaskRepository {
	return LocalTaskRepository{}
}

func (r LocalTaskRepository) ReadTask(ID string) (outputResult taskDomain.TaskStructure, err error) {
	log.Printf("User provided ID to read: %s\n", ID)
	return taskDomain.TaskStructure{
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

func (r LocalTaskRepository) ReadTasks() []taskDomain.TaskStructure {
	return []taskDomain.TaskStructure{
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

func (r LocalTaskRepository) CreateTask(inputTask *taskDomain.TaskStructure) (err error) {
	log.Println("User provide new Task input: ", inputTask)
	inputTask.Id = uuid.NewString()
	return
}

func (r LocalTaskRepository) UpdateTask(inputTask taskDomain.TaskStructure) (err error) {
	log.Println("User provide updated Task input for ID [", inputTask.Id, "]: ", inputTask)
	return
}

func (r LocalTaskRepository) DeleteTask(ID string) (err error) {
	log.Printf("User provided ID to delete: %s\n", ID)
	return nil
}
