-- name: CreateBook :one
INSERT INTO books (name, price)
VALUES ($1, $2)
    RETURNING *;

-- name: GetBookById :one
SELECT * FROM books WHERE id = $1;

-- name: ListBooks :many
SELECT * FROM books ORDER BY id;
