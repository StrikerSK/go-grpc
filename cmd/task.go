/*
Copyright © 2022 Teh|Striker
*/
package cmd

import (
	"fmt"
	taskHandler "github.com/StrikerSK/go-grpc/client/task/handler"
	taskService "github.com/StrikerSK/go-grpc/client/task/service"
	taskServer "github.com/StrikerSK/go-grpc/server"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"log"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		mode, err := cmd.Flags().GetString("mode")

		if err != nil {
			fmt.Println(err)
		}

		switch mode {
		case "server":
			taskServer.NewTaskGrpcServer().RunServer()
		case "client":
			app := fiber.New()
			handler := taskHandler.NewTaskServiceHandler(taskService.NewTaskClientService())
			handler.EnrichRouter(app)
			log.Fatal(app.Listen(fmt.Sprintf(":8080")))
		default:
			fmt.Println("No mode defined")
		}
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
	taskCmd.Flags().StringP("mode", "m", "server", "Run server mode")
}
