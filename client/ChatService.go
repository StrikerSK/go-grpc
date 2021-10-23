package client

import (
	"context"
	"github.com/StrikerSK/go-grpc/proto/chat"
	"github.com/StrikerSK/go-grpc/src"
	"google.golang.org/grpc"
	"log"
)

func SendMessage() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(src.ResolvePortNumber(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{Body: "Hello server"}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error calling method: %v\n", err)
	}

	log.Printf("Server response: %s\n", response.Body)
}
