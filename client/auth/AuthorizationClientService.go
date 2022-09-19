package auth

import (
	"context"
	"fmt"
	"github.com/StrikerSK/go-grpc/proto/auth"
	"github.com/StrikerSK/go-grpc/server/auth/domain"
	"github.com/StrikerSK/go-grpc/src"
	"google.golang.org/grpc"
	"log"
)

type AuthorizationClientService struct {
	client auth.AuthorizationServiceClient
}

func NewAuthorizationClientService() *AuthorizationClientService {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(src.ResolvePortNumber(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	return &AuthorizationClientService{
		client: auth.NewAuthorizationServiceClient(conn),
	}
}

func (c *AuthorizationClientService) RegisterUser() {
	user := domain.User{
		Username: "tester",
		Password: "123",
		Email:    "tester@test.com",
	}

	res, _ := c.client.RegisterUser(context.Background(), user.ToRegisterRequest())
	fmt.Println(res.Status)
	return
}
