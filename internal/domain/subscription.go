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
	GetById(ctx context.Context, id int) (*Subscription, error)
	Update(ctx context.Context, subscription *Subscription) error
	Delete(ctx context.Context, id int) error
}

type SubscriptionService interface{}
