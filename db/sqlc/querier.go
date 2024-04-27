// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	DeleteTable(ctx context.Context) ([]User, error)
	FindBYEmail(ctx context.Context, email string) (User, error)
	FindBYUsername(ctx context.Context, username string) (User, error)
	IsLoggedIn(ctx context.Context, username string) (Session, error)
	LoginUser(ctx context.Context, username string) (User, error)
	RegisterUser(ctx context.Context, arg RegisterUserParams) (User, error)
	UpdateSession(ctx context.Context, arg UpdateSessionParams) (Session, error)
}

var _ Querier = (*Queries)(nil)
