-- name: InsertUser :one
INSERT INTO users
  (id, created_at, updated_at, email, password_hash, is_admin)
VALUES
  ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;