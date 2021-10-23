package client

import (
	"context"
	"github.com/StrikerSK/go-grpc/proto/chat"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
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
