package handlers

import (
	"database/sql"
	"fmt"
	db "github.com/Dwimpy/simple-motion/db/sqlc"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	gowebly "github.com/gowebly/helpers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// runServer runs a new HTTP server with the loaded environment variables.

package handlers

import (
"database/sql"
db "github.com/Dwimpy/simple-motion/db/sqlc"
gowebly "github.com/gowebly/helpers"
"github.com/labstack/echo/v4"
)

// runServer runs a new HTTP server with the loaded environment variables.

type Server struct {
	queries *db.Queries
	router  *echo.Echo
}

func NewServer() *Server {

	conn, err := sql.Open(gowebly.Getenv("DB_DRIVER", "postgres"), gowebly.Getenv("DB_SOURCE", "secret"))
	if err != nil {
		panic("Failed to connect to database")
	}
	// Create a new Echo server.
	queries := db.New(conn)
	router := echo.New()
	// Add Echo middlewares.

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/

	return &Server{
		queries: queries,
		router:  router,
	}
}

func (server *Server) Start(port int) {
	server.router.Use(middleware.Logger())

	// Handle static files.
	server.router.Static("/static", "./static")

	// Handle index page view.
	server.router.GET("/", indexViewHandler)

	// Handle API endpoints.
	server.router.GET("/api/hello-world", showContentAPIHandler)

	serv := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      server.router, // handle all Echo routes
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	slog.Info("Starting server...", "port", port)
	err := serv.ListenAndServe()
	if err != nil {
		panic("Failed to start server")
	}
}

//func RunServer() *Server {
//	// Validate environment variables.
//	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "8080"))
//	if err != nil {
//		return err
//	}
//
//	conn, err := sql.Open(gowebly.Getenv("DB_DRIVER", "postgres"), gowebly.Getenv("DB_SOURCE", "secret"))
//	// Create a new Echo server.
//	queries := db.New(conn)
//	router := echo.New()
//	// Add Echo middlewares.
//	router.Use(middleware.Logger())
//
//	// Handle static files.
//	router.Static("/static", "./static")
//
//	// Handle index page view.
//	router.GET("/", indexViewHandler)
//
//	// Handle API endpoints.
//	router.GET("/api/hello-world", showContentAPIHandler)

	//Create a new server instance with options from environment variables.
	//For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	//server := http.Server{
	//	Addr:         fmt.Sprintf(":%d", port),
	//	Handler:      router, // handle all Echo routes
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//}

//	// Send log message.
//	slog.Info("Starting server...", "port", port)
//
//	return Server {
//		queries: queries,
//		router:
//	}
//}
