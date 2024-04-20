-- name: InsertUser :one
INSERT INTO users
  (id, created_at, updated_at, email, password_hash)
VALUES
  ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByCredentials :one
SELECT * FROM users
WHERE email = $1 AND password_hash = $2
LIMIT 1;