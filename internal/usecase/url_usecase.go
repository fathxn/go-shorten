package usecase

import (
	"context"
	"go-shorten/internal/domain"
	"go-shorten/internal/domain/dto"
	"go-shorten/internal/domain/model"
)

type urlUsecase struct {
	urlRepository domain.UrlRepository
	urlMapper     domain.UrlMapper
}

func NewUrlUseacse(
	urlRepository domain.UrlRepository,
	urlMapper domain.UrlMapper,
) domain.UrlUsecase {
	return &urlUsecase{
		urlRepository: urlRepository,
		urlMapper:     urlMapper,
	}
}

func (u *urlUsecase) CreateShortUrl(ctx context.Context, request dto.RequestCreateShortUrl) (*model.Url, error) {

	return nil, nil
}
