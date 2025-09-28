package postgres

import (
	"context"
	"go-shorten/internal/domain"
	"go-shorten/internal/domain/model"
	"log/slog"

	"github.com/uptrace/bun"
)

type urlRepository struct {
	db *bun.DB
}

func NewUrlRepository(db *bun.DB) domain.UrlRepository {
	return &urlRepository{
		db: db,
	}
}

func (r *urlRepository) Insert(ctx context.Context, url model.Url) error {
	_, err := r.db.NewInsert().Model(&url).Exec(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "error while inserting data")
	}

	return nil
}
