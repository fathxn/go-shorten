package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id                         uuid.UUID  `db:"id"`
	Name                       string     `db:"name"`
	Email                      string     `db:"email"`
	PasswordHash               string     `db:"password_hash"`
	IsVerified                 bool       `db:"is_verified"`
	VerifiedAt                 *time.Time `db:"verified_at"`
	VerificationToken          string     `db:"verification_token"`
	VerificationTokenExpiresAt time.Time  `db:"verification_token_expires_at"`
	CreatedAt                  time.Time  `db:"created_at"`
	UpdatedAt                  time.Time  `db:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByVerificationToken(ctx context.Context, token string) (*User, error)
	UpdateVerificationStatus(ctx context.Context, userId uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserUsecase interface {
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	GetURLsByUserId(ctx context.Context, userId string) (*[]URL, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
