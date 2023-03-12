package api

import (
	db "github.com/jsritawan/simple-todo-example/db/sqlc"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server serves HTTP requests for service.
type Server struct {
	store  *db.Store
	router *echo.Echo
}

// NewServer creates a new HTTP server and registers routes.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := echo.New()

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// Register routes
	router.POST("/todos", server.createTodo)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Start(address)
}
