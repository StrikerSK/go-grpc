package server

import (
	todoProto "github.com/StrikerSK/go-grpc/commons/proto/todo"
	"github.com/StrikerSK/go-grpc/commons/src"
	todoRepository "github.com/StrikerSK/go-grpc/server/Repository"
	todoService "github.com/StrikerSK/go-grpc/server/service"
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
	service := todoService.NewTodoLocalService(repository)
	grpcService := todoService.NewTodoGrpcService(service)

	grpcServer := grpc.NewServer()
	todoProto.RegisterTodoServiceServer(grpcServer, grpcService)

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
