package server

import (
	taskProto "github.com/StrikerSK/go-grpc/commons/proto/task"
	portResolver "github.com/StrikerSK/go-grpc/commons/src"
	taskRepository "github.com/StrikerSK/go-grpc/server/Repository"
	taskService "github.com/StrikerSK/go-grpc/server/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type TaskGrpcServer struct {
	server *grpc.Server
}

func NewTaskGrpcServer() *TaskGrpcServer {
	repository := taskRepository.NewLocalTaskRepository()
	service := taskService.NewTaskLocalService(repository)
	grpcService := taskService.NewTaskGrpcService(service)

	grpcServer := grpc.NewServer()
	taskProto.RegisterTaskServiceServer(grpcServer, grpcService)

	return &TaskGrpcServer{
		server: grpcServer,
	}
}

func (r TaskGrpcServer) RunServer() {
	lis, err := net.Listen("tcp", portResolver.ResolvePortNumber())
	if err != nil {
		log.Printf("Server init: %v\n", err)
		os.Exit(1)
	}

	if err = r.server.Serve(lis); err != nil {
		log.Printf("Server init: %v\n", err)
		os.Exit(1)
	}
}
