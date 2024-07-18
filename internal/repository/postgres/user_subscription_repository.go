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
		RETURNING id, created_at, updated_at;
	`

	return r.db.QueryRowxContext(ctx, query,
		userSubscription.UserId,
		userSubscription.SubscriptionId,
		userSubscription.StartDate,
		userSubscription.EndDate,
		userSubscription.IsActive,
	).Scan(&userSubscription.Id, &userSubscription.CreatedAt, &userSubscription.UpdatedAt)
}

// GetById implements domain.UserSubscriptionRepository.
func (r *userSubscriptionRepository) GetById(ctx context.Context, id int) (*domain.UserSubscription, error) {
	var userSubscription domain.UserSubscription
	query := `
		SELECT id, user_id, subscription_id, start_date, end_date, is_active, created_at, updated_at
		WHERE id = $1
		LIMIT 1;
	`
	err := r.db.GetContext(ctx, &userSubscription, query, id)
	if err != nil {
		return nil, err
	}
	return &userSubscription, nil
}

// GetByUserId implements domain.UserSubscriptionRepository.
func (r *userSubscriptionRepository) GetByUserId(ctx context.Context, userId string) ([]*domain.UserSubscription, error) {
	var userSubscriptions []*domain.UserSubscription
	query := `
		SELECT id, user_id, subscription_id, start_date, end_date, is_active, created_at, updated_at
		WHERE user_id = $1
		LIMIT 1;
	`
	err := r.db.SelectContext(ctx, &userSubscriptions, query, userId)
	if err != nil {
		return nil, err
	}
	return userSubscriptions, nil
}

// Update implements domain.UserSubscriptionRepository.
func (r *userSubscriptionRepository) Update(ctx context.Context, userSubscription *domain.UserSubscription) error {
	query := `
		UPDATE user_subscriptions
		SET user_id = $1, subscription_id = $2, start_date = $3, end_date = $4, is_active = $5, updated_at = NOW()
		WHERE id = $6
		RETURNING updated_at
	`
	return r.db.QueryRowxContext(
		ctx,
		query,
		userSubscription.UserId,
		userSubscription.SubscriptionId,
		userSubscription.StartDate,
		userSubscription.EndDate,
		userSubscription.IsActive,
		userSubscription.Id,
	).Scan(&userSubscription.UpdatedAt)
}

// Delete implements domain.UserSubscriptionRepository.
func (r *userSubscriptionRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM user_subscriptions WHERE id = $1;`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
