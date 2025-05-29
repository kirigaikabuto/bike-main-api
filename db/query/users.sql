-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2)
    RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id;
