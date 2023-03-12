-- name: CreateTodo :one
INSERT INTO todos (
    title
) VALUES (
    $1
) RETURNING *;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id;

-- name: UpdateTodo :one
UPDATE todos
SET title = $2, update_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
UPDATE todos
SET delete_at = $2
WHERE id = $1;