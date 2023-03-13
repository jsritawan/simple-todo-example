package db

import (
	"context"
	"testing"
	"time"

	"github.com/jsritawan/simple-todo-example/util"
	"github.com/stretchr/testify/require"
)

func createRandomTodo(t *testing.T) Todo {
	// Arrange
	note := util.RandomTodoNote()

	// Act
	todo, err := testQuery.CreateTodo(context.Background(), note)

	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, note, todo.Note)

	require.NotZero(t, todo.ID)
	require.NotZero(t, todo.CreateAt)

	return todo
}

func TestCreateTodo(t *testing.T) {
	createRandomTodo(t)
}

func TestListTodo(t *testing.T) {
	// Arrange
	todo := createRandomTodo(t)

	// Act
	todos, err := testQuery.ListTodos(context.Background())

	// Assert
	require.NoError(t, err)
	require.Greater(t, len(todos), 0)

	latestTodo := todos[0]
	require.Equal(t, todo.ID, latestTodo.ID)
	require.Equal(t, todo.Note, latestTodo.Note)
	require.Equal(t, todo.Completed, latestTodo.Completed)
	require.WithinDuration(t, todo.CreateAt, latestTodo.CreateAt, time.Second*1)
	require.Equal(t, todo.UpdateAt, latestTodo.UpdateAt)
	require.Equal(t, todo.DeleteAt, latestTodo.DeleteAt)
}
