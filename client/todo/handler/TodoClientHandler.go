package handler

import (
	todoService "github.com/StrikerSK/go-grpc/client/todo/service"
	"github.com/StrikerSK/go-grpc/server/Entity"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type TodoServiceHandler struct {
	service todoService.TodoClientService
}

func NewTodoServiceHandler(service todoService.TodoClientService) *TodoServiceHandler {
	return &TodoServiceHandler{
		service: service,
	}
}

func (r *TodoServiceHandler) ReadTodo(c *fiber.Ctx) error {
	parameter := c.Params("id")

	todo, err := r.service.ReadTodo(parameter)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(todo)
}

func (r *TodoServiceHandler) CreateTodo(c *fiber.Ctx) error {
	var tmpTodo Entity.TodoStructure
	_ = c.BodyParser(&tmpTodo)

	tmpTodo.Id = uuid.New().String()
	_ = r.service.CreateTodo(tmpTodo)

	return c.JSON(map[string]string{"data": tmpTodo.Id})
}

func (r *TodoServiceHandler) UpdateTodo(c *fiber.Ctx) error {
	var tmpTodo Entity.TodoStructure
	if err := c.BodyParser(&tmpTodo); err != nil {
		log.Printf("%v\n", err)
	}

	tmpTodo.Id = c.Params("id")
	_ = r.service.UpdateTodo(tmpTodo)
	return c.SendStatus(http.StatusOK)
}

func (r *TodoServiceHandler) DeleteTodo(c *fiber.Ctx) error {
	_, _ = r.service.DeleteTodo(c.Params("id"))
	return c.SendStatus(http.StatusOK)
}

func (r *TodoServiceHandler) FindTasks(c *fiber.Ctx) error {
	return c.JSON(r.service.GetTodos())
}
