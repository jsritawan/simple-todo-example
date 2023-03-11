// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: todo.sql

package db

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (
    title
) VALUES (
    $1
) RETURNING id, title, completed, create_at, update_at, delete_at
`

func (q *Queries) CreateTodo(ctx context.Context, title string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, title)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.CreateAt,
		&i.UpdateAt,
		&i.DeleteAt,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
UPDATE todos
SET delete_at = $2
WHERE id = $1
`

type DeleteTodoParams struct {
	ID       int64        `json:"id"`
	DeleteAt sql.NullTime `json:"delete_at"`
}

func (q *Queries) DeleteTodo(ctx context.Context, arg DeleteTodoParams) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, arg.ID, arg.DeleteAt)
	return err
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, completed, create_at, update_at, delete_at FROM todos
ORDER BY id
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Completed,
			&i.CreateAt,
			&i.UpdateAt,
			&i.DeleteAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET title = $2, update_at = $3
WHERE id = $1
RETURNING id, title, completed, create_at, update_at, delete_at
`

type UpdateTodoParams struct {
	ID       int64        `json:"id"`
	Title    string       `json:"title"`
	UpdateAt sql.NullTime `json:"update_at"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo, arg.ID, arg.Title, arg.UpdateAt)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.CreateAt,
		&i.UpdateAt,
		&i.DeleteAt,
	)
	return i, err
}
