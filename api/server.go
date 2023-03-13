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
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register routes
	gTodo := e.Group("/todos")
	gTodo.POST("", server.createTodo)
	gTodo.GET("", server.listTodo)
	gTodo.PUT("/:id", server.updateTodo)
	gTodo.DELETE("/:id", server.deleteTodo)

	server.router = e
	return server
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Start(address)
}
