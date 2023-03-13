-- name: CreateTodo :one
INSERT INTO todos (
    note
) VALUES (
    $1
) RETURNING *;

-- name: ListTodos :many
SELECT * FROM todos
WHERE delete_at IS NULL
ORDER BY id DESC;

-- name: UpdateTodo :one
UPDATE todos
SET note = $2, completed = $3, update_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
UPDATE todos
SET delete_at = $2
WHERE id = $1;