package domain

import (
	"context"
	"time"
)

type User struct {
	Id                         string     `db:"id"`
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
	GetById(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByVerificationToken(ctx context.Context, token string) (*User, error)
	UpdateVerificationStatus(ctx context.Context, userId string, verifiedAt *time.Time) error
	Delete(ctx context.Context, id string) error
}

type UserUsecase interface {
	RegisterUser(ctx context.Context, registerInput *UserRegisterInput) (string, error)
	VerifyEmail(ctx context.Context, token string) error
	LoginUser(ctx context.Context, loginInput *UserLoginInput) (*User, error)
	GetById(ctx context.Context, id string) (*User, error)
	GetURLsByUserId(ctx context.Context, userId string) (*[]URL, error)
	Delete(ctx context.Context, id string) error
}

type UserRegisterInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserAuthResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type Auth struct {
	Id string
}
