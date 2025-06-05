-- name: CreateBook :one
INSERT INTO books (id, title, author, pages)
VALUES ($1, $2, $3, $4)
RETURNING id, title, author, pages;

-- name: GetBook :one
SELECT id, title, author, pages
FROM books
WHERE id = $1;

-- name: GetAllBooks :many
SELECT id, title, author, pages
FROM books
ORDER BY id;

-- name: UpdateBook :one
UPDATE books
SET title = $2, author = $3, pages = $4
WHERE id = $1
RETURNING id, title, author, pages;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;