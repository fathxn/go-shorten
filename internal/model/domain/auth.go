package domain

import (
	"context"
	"go-shorten/internal/model/dto"
)

type AuthUsecase interface {
	RegisterUser(ctx context.Context, registerInput *dto.UserRegisterInput) (string, error)
	VerifyEmail(ctx context.Context, token string) error
	LoginUser(ctx context.Context, loginInput *dto.UserLoginInput) (*User, error)
}
