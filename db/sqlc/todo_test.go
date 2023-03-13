package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T) {
	// Arrange
	note := "Test Create Todo"

	// Act
	todo, err := testQuery.CreateTodo(context.Background(), note)

	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, note, todo.Note)

	require.NotZero(t, todo.ID)
	require.NotZero(t, todo.CreateAt)
}
