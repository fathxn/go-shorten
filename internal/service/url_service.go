package service

import (
	"context"
	"go-short-url/internal/model/domain"
)

type URLService interface {
	Create(ctx context.Context, longURL string) (*domain.URL, error)
	GetLongURL(ctx context.Context, shortCode string) (*domain.URL, error)
	GetById(ctx context.Context, id int) (*domain.URL, error)
	Delete(ctx context.Context, id int) error
}
