package domain

import (
	"context"
	"time"
)

type User struct {
	Id                         string    `db:"id"`
	Name                       string    `db:"name"`
	Email                      string    `db:"email"`
	PasswordHash               string    `db:"password_hash"`
	IsVerified                 bool      `db:"is_verified"`
	VerificationToken          string    `db:"verification_token"`
	VerificationTokenExpiresAt time.Time `db:"verification_token_expires_at"`
	CreatedAt                  time.Time `db:"created_at"`
	UpdatedAt                  time.Time `db:"updated_at"`
}

type UserRepository interface {
	Insert(ctx context.Context, user *User) error
	FindById(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Delete(ctx context.Context, id string) error
}

type UserService interface {
	GetById(ctx context.Context, id string) (*User, error)
	GetURLsByUserId(ctx context.Context, userId string) (*[]URL, error)
	Delete(ctx context.Context, id string) error
}
