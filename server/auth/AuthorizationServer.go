package server

import (
	"context"
	"github.com/StrikerSK/go-grpc/proto/auth"
	"github.com/StrikerSK/go-grpc/server/auth/ports"
	"github.com/StrikerSK/go-grpc/server/auth/service"
	"github.com/StrikerSK/go-grpc/src"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type AuthorizationServer struct {
	service ports.IUserService
	auth.UnimplementedAuthorizationServiceServer
}

func (c *AuthorizationServer) RegisterUser(ctx context.Context, in *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	_ = c.service.RegisterUser(in)
	return nil, nil
}

func CreateAuthorizationServer() {
	lis, err := net.Listen("tcp", src.ResolvePortNumber())
	if err != nil {
		log.Printf("Server init: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	auth.RegisterAuthorizationServiceServer(grpcServer, &AuthorizationServer{
		service: service.NewConsoleUserService(),
	})

	if err = grpcServer.Serve(lis); err != nil {
		log.Printf("Server init: %v\n", err)
		os.Exit(1)
	}
}
