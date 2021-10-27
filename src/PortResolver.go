package src

import (
	"fmt"
	"log"
	"os"
)

func ResolvePortNumber() string {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("Default port [9000] has been set!")
		port = "9000"
	}

	return fmt.Sprintf("localhost:%s", port)
}
