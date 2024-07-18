package postgres

import (
	"context"
	"go-short-url/internal/model/domain"

	"github.com/jmoiron/sqlx"
)

type urlRepository struct {
	db *sqlx.DB
}

func NewURLRepository(db *sqlx.DB) domain.URLRepository {
	return &urlRepository{db: db}
}

// Insert implements domain.URLRepository.
func (r *urlRepository) Insert(ctx context.Context, url *domain.URL) error {
	query := `
		INSERT INTO url (user_id, long_url, short_code)
		VALUES ($1, $2, $3)
	`

	return r.db.QueryRowxContext(ctx, query, url.UserId, url.LongURL, url.ShortCode).Scan()
}

// FindByShortCode implements domain.URLRepository.
func (r *urlRepository) FindByShortCode(ctx context.Context, shortCode string) (*domain.URL, error) {
	panic("unimplemented")
}

// FindById implements domain.URLRepository.
func (r *urlRepository) FindById(ctx context.Context, id int) (*domain.URL, error) {
	panic("unimplemented")
}

// FindByUserId implements domain.URLRepository.
func (r *urlRepository) FindByUserId(ctx context.Context, userId string) ([]domain.URL, error) {
	panic("unimplemented")
}

// Delete implements domain.URLRepository.
func (r *urlRepository) Delete(ctx context.Context, id int) error {
	execQuery := "DELETE FROM urls WHERE id = $1;"
	_, err := r.db.ExecContext(ctx, execQuery, id)
	if err != nil {
		return err
	}

	return nil
}
