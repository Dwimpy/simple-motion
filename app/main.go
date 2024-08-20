package main

import (
	"github.com/Dwimpy/simple-motion/handlers"
	"github.com/joho/godotenv"
	"log/slog"
)

func main() {
	// Run your server.

	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}

	server := handlers.NewServer()
}
