package handler

import (
	todoService "github.com/StrikerSK/go-grpc/client/todo/service"
	todoDomain "github.com/StrikerSK/go-grpc/commons/todo/domain"
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

func (r *TodoServiceHandler) EnrichRouter(router fiber.Router) {
	router.Get("/todo/:id", r.ReadTodo)
	router.Post("/todo", r.CreateTodo)
	router.Put("/todo/:id", r.UpdateTodo)
	router.Delete("/todo/:id", r.DeleteTodo)
	router.Get("/todo", r.FindTasks)
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
	var tmpTodo todoDomain.TodoStructure
	_ = c.BodyParser(&tmpTodo)

	newID := uuid.New().String()
	tmpTodo.Id = newID
	_, _ = r.service.CreateTodo(tmpTodo)

	return c.JSON(map[string]string{"data": newID})
}

func (r *TodoServiceHandler) UpdateTodo(c *fiber.Ctx) error {
	var tmpTodo todoDomain.TodoStructure
	if err := c.BodyParser(&tmpTodo); err != nil {
		log.Printf("%v\n", err)
	}

	tmpTodo.Id = c.Params("id")
	if err := r.service.UpdateTodo(tmpTodo); err != nil {
		log.Printf("%v\n", err)
		c.Status(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func (r *TodoServiceHandler) DeleteTodo(c *fiber.Ctx) error {
	if err := r.service.DeleteTodo(c.Params("id")); err != nil {
		log.Printf("%v\n", err)
		c.Status(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func (r *TodoServiceHandler) FindTasks(c *fiber.Ctx) error {
	todos, err := r.service.GetTodos()
	if err != nil {
		log.Printf("%v\n", err)
		c.Status(http.StatusInternalServerError)
	}

	return c.JSON(todos)
}
