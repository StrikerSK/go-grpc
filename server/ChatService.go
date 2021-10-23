package server

import (
	"github.com/StrikerSK/go-grpc/proto/chat"
	"github.com/StrikerSK/go-grpc/src"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func CreateServer() {
	lis, err := net.Listen("tcp", src.ResolvePortNumber())
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
