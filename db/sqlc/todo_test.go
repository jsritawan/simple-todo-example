package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T) {
	// Arrange
	title := "Test Create Todo"

	// Act
	todo, err := testQuery.CreateTodo(context.Background(), title)

	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, title, todo.Title)

	require.NotZero(t, todo.ID)
	require.NotZero(t, todo.CreateAt)
}
