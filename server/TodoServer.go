package server

import (
	"github.com/StrikerSK/go-grpc/proto/todo"
	todoRepository "github.com/StrikerSK/go-grpc/server/Repository"
	todoService "github.com/StrikerSK/go-grpc/server/service"
	"github.com/StrikerSK/go-grpc/src"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type TodoGrpcServer struct {
	server *grpc.Server
}

func NewTodoGrpcServer() *TodoGrpcServer {
	repository := todoRepository.NewLocalTodoRepository()
	service := todoService.NewTodoService(&repository)

	grpcServer := grpc.NewServer()
	todo.RegisterTodoServiceServer(grpcServer, service)

	return &TodoGrpcServer{
		server: grpcServer,
	}
}

func (r TodoGrpcServer) RunServer() {
	lis, err := net.Listen("tcp", src.ResolvePortNumber())
	if err != nil {
		log.Printf("Server init: %v\n", err)
		os.Exit(1)
	}

	if err = r.server.Serve(lis); err != nil {
		log.Printf("Server init: %v\n", err)
		os.Exit(1)
	}
}
