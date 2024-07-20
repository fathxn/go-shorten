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
	MaxURLs   int       `db:"max_urls"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *Subscription) error
}

type SubscriptionService interface{}
