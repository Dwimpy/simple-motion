package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"

	"github.com/Dwimpy/simple-motion/handlers"
)

func main() {
	// Run your server.

	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}

	if err := handlers.RunServer(); err != nil {
		slog.Error("Failed to start server!", "details", err.Error())
		os.Exit(1)
	}
}
