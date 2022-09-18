package main

import (
	"github.com/StrikerSK/go-grpc/client/auth"
)

func main() {
	client := auth.NewAuthorizationClient()
	client.RegisterUser()
}
