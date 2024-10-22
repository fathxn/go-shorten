package domain

import (
	"context"
	"time"
)

type UserSubscription struct {
	Id             int       `db:"id"`
	UserId         string    `db:"user_id"`
	SubscriptionId int       `db:"subscription_id"`
	StartDate      time.Time `db:"start_date"`
	EndDate        time.Time `db:"end_date"`
	IsActive       int       `db:"is_active"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type UserSubscriptionRepository interface {
	Create(ctx context.Context, userSubscription *UserSubscription) error
	GetById(ctx context.Context, id int) (*UserSubscription, error)
	GetByUserId(ctx context.Context, userId string) ([]*UserSubscription, error)
	Update(ctx context.Context, userSubscription *UserSubscription) error
	Delete(ctx context.Context, id int) error
}

type UserSubscriptionService interface {
}
