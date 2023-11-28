package service

import (
	"context"
	"go-short-url/internal/model/api"
	"go-short-url/internal/model/domain"
)

type UserService interface {
	Create(ctx context.Context, userInput *api.UserInput) error
	FindById(ctx context.Context, id string) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}
