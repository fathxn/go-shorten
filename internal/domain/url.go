package domain

import (
	"context"
	"go-shorten/internal/domain/dto"
	"go-shorten/internal/domain/model"
)

type UrlRepository interface {
	Insert(ctx context.Context, url model.Url) error
}

type UrlUsecase interface {
	CreateShortUrl(ctx context.Context, request dto.RequestCreateShortUrl) (*model.Url, error)
}

type UrlMapper interface {
	ToCreateShortUrlMapper(ctx context.Context, request dto.RequestCreateShortUrl) (*model.Url, error)
	ToResponseShortUrlMapper(ctx context.Context, url model.Url) dto.Response
}
