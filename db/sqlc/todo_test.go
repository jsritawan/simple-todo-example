package db

import (
	"context"
	"database/sql"
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
	numOfTodos := 10

	for i := 0; i < numOfTodos; i++ {
		createRandomTodo(t)
	}

	// Act
	todos, err := testQuery.ListTodos(context.Background())

	// Assert
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(todos), numOfTodos)

	for _, todo := range todos {
		require.NotEmpty(t, todo)
	}
}

func TestUpdateTodo(t *testing.T) {

	todo1 := createRandomTodo(t)

	arg := UpdateTodoParams{
		ID:        todo1.ID,
		Note:      util.RandomTodoNote(),
		Completed: true,
		UpdateAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	todo2, err := testQuery.UpdateTodo(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, arg.Note, todo2.Note)
	require.Equal(t, arg.Completed, todo2.Completed)
	require.WithinDuration(t, todo1.CreateAt, todo2.CreateAt, time.Second*1)
	require.WithinDuration(t, arg.UpdateAt.Time, todo2.UpdateAt.Time, time.Second*1)
	require.Equal(t, todo1.DeleteAt, todo2.DeleteAt)
}

func TestDeleteTodo(t *testing.T) {
	todo1 := createRandomTodo(t)

	arg := DeleteTodoParams{
		ID: todo1.ID,
		DeleteAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := testQuery.DeleteTodo(context.Background(), arg)

	require.NoError(t, err)
}
