package service

import (
	"context"
	"go-short-url/internal/model/api"
	"go-short-url/internal/model/domain"
)

type UserService interface {
	RegisterUser(ctx context.Context, userInput *api.UserRegisterInput) error
	LoginUser(ctx context.Context, loginInput *api.UserLoginInput) (*domain.User, error)
	GetById(ctx context.Context, id string) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}
