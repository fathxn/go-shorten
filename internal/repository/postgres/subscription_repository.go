package postgres

import (
	"context"
	"go-short-url/internal/model/domain"

	"github.com/jmoiron/sqlx"
)

type subscriptionRepository struct {
	db *sqlx.DB
}

func NewSubscriptionRepository(db *sqlx.DB) domain.SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

// Create implements domain.SubscriptionRepository.
func (r *subscriptionRepository) Create(ctx context.Context, subscription *domain.Subscription) error {
	query := `
		INSERT INTO subscription (name, price, duration, max_urls)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowxContext(ctx, query,
		subscription.Name,
		subscription.Price,
		subscription.Duration,
		subscription.MaxURLs,
	).Scan(&subscription.Id, &subscription.CreatedAt, &subscription.UpdatedAt)
}

// GetById implements domain.SubscriptionRepository.
func (r *subscriptionRepository) GetById(ctx context.Context, id int) (*domain.Subscription, error) {
	panic("unimplemented")
}

// Update implements domain.SubscriptionRepository.
func (r *subscriptionRepository) Update(ctx context.Context, subscription *domain.Subscription) error {
	panic("unimplemented")
}

// Delete implements domain.SubscriptionRepository.
func (r *subscriptionRepository) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}
