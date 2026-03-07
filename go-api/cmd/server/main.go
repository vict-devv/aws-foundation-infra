package main

import (
	"log"
	"net/http"
	"os"

	"github.com/victor-devv/cloud-native-api/internal/health"
)

func main() {
	registerRoutes()

	port := serverPort()
	log.Printf("Server starting on :%s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func registerRoutes() {
	http.HandleFunc("/health", health.Handler)
}

func serverPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return "8080"
	}

	return port
}
