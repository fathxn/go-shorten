package repository

import (
	"context"
	"go-short-url/internal/model/domain"
)

type URLRepository interface {
	Insert(ctx context.Context, url *domain.URL) error
	FindByShortCode(ctx context.Context, shortCode string) (*domain.URL, error)
	FindById(ctx context.Context, id int) (*domain.URL, error)
	FindByUserId(ctx context.Context, userId string) ([]domain.URL, error)
	Delete(ctx context.Context, id int) error
}
