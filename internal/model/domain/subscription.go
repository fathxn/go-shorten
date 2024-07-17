package domain

import (
	"context"
	"time"
)

type Subscription struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Price     int64     `db:"price"`
	Duration  time.Time `db:"duration"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

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

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *Subscription) error
}

type SubscriptionService interface{}
