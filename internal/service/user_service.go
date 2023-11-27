package service

import (
	"context"
	"go-short-url/internal/model/domain"
)

type UserService interface {
	Create(ctx context.Context, user *domain.User) error
	FindById(ctx context.Context, id string) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}
