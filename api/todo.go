package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type TodoResponse struct {
	ID        int64     `json:"id"`
	Note      string    `json:"note"`
	Completed bool      `json:"completed"`
	CreateAt  time.Time `json:"createAt"`
}

type createTodoRequest struct {
	Note string `json:"note" validate:"required"`
}

type createTodoRequestValidator struct {
	validator *validator.Validate
}

func (v *createTodoRequestValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (s *Server) createTodo(c echo.Context) (err error) {
	var req createTodoRequest

	if err = c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	s.router.Validator = &createTodoRequestValidator{validator: validator.New()}
	if err = c.Validate(&req); err != nil {
		return err
	}

	todo, err := s.store.CreateTodo(context.Background(), req.Note)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, TodoResponse{
		ID:        todo.ID,
		Note:      todo.Note,
		Completed: todo.Completed,
		CreateAt:  todo.CreateAt,
	})
}

func (s *Server) listTodo(c echo.Context) (err error) {
	var todos []TodoResponse

	results, err := s.store.ListTodos(context.Background())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for _, todo := range results {
		todos = append(todos, TodoResponse{
			ID:        todo.ID,
			Note:      todo.Note,
			Completed: todo.Completed,
			CreateAt:  todo.CreateAt,
		})
	}

	return c.JSON(http.StatusOK, todos)
}
