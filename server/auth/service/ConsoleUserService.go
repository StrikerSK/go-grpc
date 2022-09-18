package service

import (
	"github.com/StrikerSK/go-grpc/proto/auth"
	"github.com/StrikerSK/go-grpc/server/auth/domain"
	"log"
)

type ConsoleUserService struct{}

func NewConsoleUserService() *ConsoleUserService {
	return &ConsoleUserService{}
}

func (c *ConsoleUserService) RegisterUser(request *auth.RegisterRequest) error {
	domain.NewUser().FromRegisterRequest(request)
	log.Printf("User %s registered\n", request.Username)
	return nil
}
