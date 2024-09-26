package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	portVar := "PORT"
	port := os.Getenv(portVar)
	if port == "" {
		log.Fatalf("Environment variable %s is not set", portVar)
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Starting cocktail server on %v\n", port)

	log.Fatal(server.ListenAndServe())
}
