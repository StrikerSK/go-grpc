package handler

import (
	taskService "github.com/StrikerSK/go-grpc/client/task/service"
	taskDomain "github.com/StrikerSK/go-grpc/commons/task/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type TaskServiceHandler struct {
	service taskService.TaskClientService
}

func NewTaskServiceHandler(service taskService.TaskClientService) *TaskServiceHandler {
	return &TaskServiceHandler{
		service: service,
	}
}

func (r *TaskServiceHandler) EnrichRouter(router fiber.Router) {
	router.Post("/task", r.CreateTask)
	router.Get("/task/:id", r.ReadTask)
	router.Get("/task", r.ReadTasks)
	router.Put("/task/:id", r.UpdateTask)
	router.Delete("/task/:id", r.DeleteTask)
}

func (r *TaskServiceHandler) CreateTask(c *fiber.Ctx) error {
	var tmpTask taskDomain.TaskStructure
	_ = c.BodyParser(&tmpTask)

	newID := uuid.New().String()
	tmpTask.Id = newID
	_, _ = r.service.CreateTask(tmpTask)

	return c.JSON(map[string]string{"data": newID})
}

func (r *TaskServiceHandler) ReadTask(c *fiber.Ctx) error {
	parameter := c.Params("id")

	task, err := r.service.ReadTask(parameter)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(task)
}

func (r *TaskServiceHandler) ReadTasks(c *fiber.Ctx) error {
	tasks, err := r.service.ReadTasks()
	if err != nil {
		log.Printf("%v\n", err)
		c.Status(http.StatusInternalServerError)
	}

	return c.JSON(tasks)
}

func (r *TaskServiceHandler) UpdateTask(c *fiber.Ctx) error {
	var tmpTask taskDomain.TaskStructure
	if err := c.BodyParser(&tmpTask); err != nil {
		log.Printf("%v\n", err)
	}

	tmpTask.Id = c.Params("id")
	if err := r.service.UpdateTask(tmpTask); err != nil {
		log.Printf("%v\n", err)
		c.Status(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func (r *TaskServiceHandler) DeleteTask(c *fiber.Ctx) error {
	if err := r.service.DeleteTask(c.Params("id")); err != nil {
		log.Printf("%v\n", err)
		c.Status(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}
