package domain

import (
	"context"
	"go-shorten/internal/domain/dto"
	"go-shorten/internal/domain/model"
)

type UrlRepository interface {
	Insert(ctx context.Context, url model.Url) error
	FindByShortCode(ctx context.Context, shortCode string) (model.Url, error)
	FindById(ctx context.Context, id int) (model.Url, error)
	FindByUserId(ctx context.Context, userId string) ([]model.Url, error)
	Delete(ctx context.Context, id int) error
}

type UrlUsecase interface {
	Create(ctx context.Context, longURL string, userId string) (*model.Url, error)
	GetLongURL(ctx context.Context, shortCode string) (*model.Url, error)
	GetById(ctx context.Context, id int) (*model.Url, error)
	// GetByUserId(ctx context.Context, userId string) (*[]domain.URL, error)
	Delete(ctx context.Context, id int) error
}

type UrlMapper interface {
	ToCreateShortUrlMapper(ctx context.Context, request dto.RequestCreateShortUrl) (*model.Url, error)
	ToResponseShortUrlMapper(ctx context.Context, url model.Url) dto.Response
}
