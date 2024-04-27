// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package db

import (
	"context"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
  username,
  refresh_token
) VALUES (
  $1, $2
) RETURNING id, username, refresh_token
`

type CreateSessionParams struct {
	Username     string `json:"username"`
	RefreshToken string `json:"refresh_token"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, createSession, arg.Username, arg.RefreshToken)
	var i Session
	err := row.Scan(&i.ID, &i.Username, &i.RefreshToken)
	return i, err
}

const deleteTable = `-- name: DeleteTable :many
DELETE FROM users RETURNING id, username, email, password
`

func (q *Queries) DeleteTable(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, deleteTable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findBYEmail = `-- name: FindBYEmail :one
SELECT id, username, email, password
FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) FindBYEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, findBYEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const findBYUsername = `-- name: FindBYUsername :one
SELECT id, username, email, password
FROM users
WHERE username = $1
LIMIT 1
`

func (q *Queries) FindBYUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, findBYUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const isLoggedIn = `-- name: IsLoggedIn :one
SELECT id, username, refresh_token
FROM sessions
WHERE username = $1
LIMIT 1
`

func (q *Queries) IsLoggedIn(ctx context.Context, username string) (Session, error) {
	row := q.db.QueryRow(ctx, isLoggedIn, username)
	var i Session
	err := row.Scan(&i.ID, &i.Username, &i.RefreshToken)
	return i, err
}

const loginUser = `-- name: LoginUser :one
SELECT id, username, email, password
FROM users
WHERE username = $1
LIMIT 1
`

func (q *Queries) LoginUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, loginUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const registerUser = `-- name: RegisterUser :one
INSERT INTO users (
  username,
  email,
  password
) VALUES (
  $1, $2, $3
) RETURNING id, username, email, password
`

type RegisterUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) RegisterUser(ctx context.Context, arg RegisterUserParams) (User, error) {
	row := q.db.QueryRow(ctx, registerUser, arg.Username, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const updateSession = `-- name: UpdateSession :one
UPDATE sessions
SET refresh_token = $1
WHERE username = $2
RETURNING id, username, refresh_token
`

type UpdateSessionParams struct {
	RefreshToken string `json:"refresh_token"`
	Username     string `json:"username"`
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, updateSession, arg.RefreshToken, arg.Username)
	var i Session
	err := row.Scan(&i.ID, &i.Username, &i.RefreshToken)
	return i, err
}
