/*
Copyright © 2022 Teh|Striker
*/
package cmd

import (
	"fmt"
	todoHandler "github.com/StrikerSK/go-grpc/client/todo/handler"
	todoService "github.com/StrikerSK/go-grpc/client/todo/service"
	todoServer "github.com/StrikerSK/go-grpc/server"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"log"
)

// todoCmd represents the todo command
var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		mode, err := cmd.Flags().GetString("mode")

		if err != nil {
			fmt.Println(err)
		}

		switch mode {
		case "server":
			todoServer.CreateTodoServer()
		case "client":
			app := fiber.New()
			handler := todoHandler.NewTodoServiceHandler(todoService.NewTodoClientService())
			handler.EnrichRouter(app)
			log.Fatal(app.Listen(fmt.Sprintf(":8080")))
		default:
			fmt.Println("No mode defined")
		}
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)
	todoCmd.Flags().StringP("mode", "m", "server", "Run server mode")
}
