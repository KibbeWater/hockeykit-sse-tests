package main

import (
	"hockeykit-tester/internal/server"
	"log"
)

func main() {
	if err := server.StartServer(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
