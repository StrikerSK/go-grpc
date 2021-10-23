package server

import (
	"github.com/StrikerSK/go-grpc/proto/chat"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func CreateServer() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Printf("Server init: %v\n", err)
		os.Exit(1)
	}

	s := Server{}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err = grpcServer.Serve(lis); err != nil {
		log.Printf("Server init: %v\n", err)
		os.Exit(1)
	}
}
