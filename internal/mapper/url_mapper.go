package mapper

import (
	"context"
	"go-shorten/internal/domain"
	"go-shorten/internal/domain/dto"
	"go-shorten/internal/domain/model"
	"time"

	"github.com/google/uuid"
)

type urlMapper struct {
}

func NewUrlMapper() domain.UrlMapper {
	return &urlMapper{}
}

func (m *urlMapper) ToCreateShortUrlMapper(ctx context.Context, request dto.RequestCreateShortUrl) (*model.Url, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	url := model.Url{
		ID:        uuid.String(),
		LongUrl:   request.LongUrl,
		ShortCode: *request.ShortCode,
		CreatedBy: "00000000-00000000-00000000-00000000",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &url, nil
}
func (m *urlMapper) ToResponseShortUrlMapper(ctx context.Context, url model.Url) dto.Response {
	response := dto.Response{}

	return response
}
