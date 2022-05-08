package client

import (
	"fmt"
	"github.com/StrikerSK/go-grpc/proto/todo"
	"github.com/StrikerSK/go-grpc/src"
	"google.golang.org/grpc"
	"log"
	"sync"
)

var lock = &sync.Mutex{}

type todoClientSingleton struct {
	client todo.TodoServiceClient
}

var singleInstance *todoClientSingleton

func getInstance() *todoClientSingleton {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &todoClientSingleton{client: createClient()}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		//fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func GetClient() todo.TodoServiceClient {
	return getInstance().client
}

func createClient() todo.TodoServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(src.ResolvePortNumber(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	return todo.NewTodoServiceClient(conn)
}
