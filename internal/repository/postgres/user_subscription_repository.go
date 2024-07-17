package postgres

import (
	"context"
	"go-short-url/internal/model/domain"

	"github.com/jmoiron/sqlx"
)

type userSubscriptionRepository struct {
	db *sqlx.DB
}

func NewUserSubscriptionRepository(db *sqlx.DB) domain.UserSubscriptionRepository {
	return &userSubscriptionRepository{db: db}
}

// Create implements domain.UserSubscriptionRepository.
func (r *userSubscriptionRepository) Create(ctx context.Context, userSubscription *domain.UserSubscription) error {
	query := `
		INSERT INTO	user_subscriptions (user_id, subscription_id, start_date, end_date, is_active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowxContext(ctx, query,
		userSubscription.UserId,
		userSubscription.SubscriptionId,
		userSubscription.StartDate,
		userSubscription.EndDate,
		userSubscription.IsActive,
	).Scan(&userSubscription.Id, &userSubscription.CreatedAt, &userSubscription.UpdatedAt)
}

// Delete implements domain.UserSubscriptionRepository.
func (r *userSubscriptionRepository) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetById implements domain.UserSubscriptionRepository.
func (r *userSubscriptionRepository) GetById(ctx context.Context, id int) (*domain.UserSubscription, error) {
	panic("unimplemented")
}

// GetByUserId implements domain.UserSubscriptionRepository.
func (r *userSubscriptionRepository) GetByUserId(ctx context.Context, userId string) ([]*domain.UserSubscription, error) {
	panic("unimplemented")
}

// Update implements domain.UserSubscriptionRepository.
func (r *userSubscriptionRepository) Update(ctx context.Context, userSubscription *domain.UserSubscription) error {
	panic("unimplemented")
}
