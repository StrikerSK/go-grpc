package service

import (
	"github.com/StrikerSK/go-grpc/proto/auth"
	"log"
)

type ConsoleUserService struct{}

func NewConsoleUserService() *ConsoleUserService {
	return &ConsoleUserService{}
}

func (c *ConsoleUserService) RegisterUser(request *auth.RegisterRequest) error {
	log.Printf("User %s registered\n", request.Username)
	return nil
}
