package server

import (
	"context"
	"github.com/StrikerSK/go-grpc/proto/chat"
	"log"
)

type Server struct {
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *chat.UserMessage) (*chat.UserMessage, error) {
	log.Printf("Client message: %s\n", message.Body)

	return &chat.UserMessage{
		Body: "Server is greeting you!",
	}, nil
}
