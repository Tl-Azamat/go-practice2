package main

import (
	"fmt"
	"go-practice2/internal/handlers"
	"go-practice2/internal/middleware"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Apply middleware to routes
	mux.Handle("/user", middleware.APIKey(http.HandlerFunc(handlers.UserHandler)))

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
