package api

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type createTodoRequest struct {
	Title string `json:"title" validate:"required"`
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

	todo, err := s.store.CreateTodo(context.Background(), req.Title)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, todo)
}