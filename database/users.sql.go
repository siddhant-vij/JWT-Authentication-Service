// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, email, password_hash, is_admin FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.PasswordHash,
		&i.IsAdmin,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, created_at, updated_at, email, password_hash, is_admin FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.PasswordHash,
		&i.IsAdmin,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users
  (id, created_at, updated_at, email, password_hash, is_admin)
VALUES
  ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, email, password_hash, is_admin
`

type InsertUserParams struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string
	PasswordHash string
	IsAdmin      bool
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Email,
		arg.PasswordHash,
		arg.IsAdmin,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.PasswordHash,
		&i.IsAdmin,
	)
	return i, err
}
