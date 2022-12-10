// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"

	"github.com/gobuffalo/nulls"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, email, hashed_password, first_name, last_name, phone, role)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, hashed_password, username, email, first_name, last_name, phone, role, created_at, password_changed_at, updated_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Phone          string `json:"phone"`
	Role           string `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
		arg.Phone,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.HashedPassword,
		&i.Username,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Role,
		&i.CreatedAt,
		&i.PasswordChangedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, hashed_password, username, email, first_name, last_name, phone, role, created_at, password_changed_at, updated_at
FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.HashedPassword,
		&i.Username,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Role,
		&i.CreatedAt,
		&i.PasswordChangedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET hashed_password     = COALESCE($1, hashed_password),
    password_changed_at = COALESCE($2, password_changed_at),
    first_name          = COALESCE($3, first_name),
    last_name           = COALESCE($4, last_name),
    email               = COALESCE($5, email),
    phone               = COALESCE($6, phone),
    role                = COALESCE($7, role)
WHERE username = $8
RETURNING id, hashed_password, username, email, first_name, last_name, phone, role, created_at, password_changed_at, updated_at
`

type UpdateUserParams struct {
	HashedPassword    nulls.String `json:"hashed_password"`
	PasswordChangedAt nulls.Time   `json:"password_changed_at"`
	FirstName         nulls.String `json:"first_name"`
	LastName          nulls.String `json:"last_name"`
	Email             nulls.String `json:"email"`
	Phone             nulls.String `json:"phone"`
	Role              nulls.String `json:"role"`
	Username          string       `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChangedAt,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.Role,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.HashedPassword,
		&i.Username,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Role,
		&i.CreatedAt,
		&i.PasswordChangedAt,
		&i.UpdatedAt,
	)
	return i, err
}