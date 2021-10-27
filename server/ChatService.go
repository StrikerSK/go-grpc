package server

import (
	"github.com/StrikerSK/go-grpc/proto/chat"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func CreateChatServer() {
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Printf("Server init: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &Server{})

	if err = grpcServer.Serve(lis); err != nil {
		log.Printf("Server init: %v\n", err)
		os.Exit(1)
	}
}
