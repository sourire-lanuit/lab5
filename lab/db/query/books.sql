-- name: CreateBook :one
INSERT INTO books (id, title, author, pages, year)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetBook :one
SELECT * FROM books WHERE id = $1;

-- name: GetAllBooks :many
SELECT * FROM books;

-- name: UpdateBook :one
UPDATE books
SET title = $2, author = $3, pages = $4, year = $5
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books WHERE id = $1;