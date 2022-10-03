/*
Copyright © 2022 Teh|Striker
*/
package cmd

import (
	"fmt"
	authorizationClient "github.com/StrikerSK/go-grpc/client/auth"
	authorizationServer "github.com/StrikerSK/go-grpc/server/auth"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Run authorization server or client",
	Run: func(cmd *cobra.Command, args []string) {
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			fmt.Println(err)
		}

		switch mode {
		case "server":
			authorizationServer.NewAuthorizationServer().RunServer()
		case "client":
			client := authorizationClient.NewAuthorizationClientService()
			client.RegisterUser()
		default:
			fmt.Println("No mode defined")
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().StringP("mode", "m", "server", "Run server mode")
}
