package main

import (
	"fmt"
	"github.com/StrikerSK/go-grpc/client/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	todoPath := app.Group("/api/todo")
	todoPath.Get("/:id", service.ReadTodo)
	todoPath.Get("", service.FindTasks)
	todoPath.Post("", service.CreateTodo)
	todoPath.Put("/:id", service.UpdateTodo)
	todoPath.Delete("/:id", service.DeleteTodo)

	log.Fatal(app.Listen(fmt.Sprintf(":8080")))
}
