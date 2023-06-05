package model

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type UserService interface {
	Get(uid uuid.UUID) (*User, error)
	Signup(ctx context.Context, u *User) error
	Signin(ctx context.Context, u *User) error
	UpdateDetails(ctx context.Context, u *User) error
}

type TokenService interface {
	Signout(ctx context.Context, uid uuid.UUID) error
}

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	FindById(uid uuid.UUID) (*User, error)
}

type TokenRepository interface {
	SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userID string, prevTokenID string) error
	DeleteUserRefreshToken(ctx context.Context, userID string) error
}
