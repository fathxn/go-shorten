package domain

import "time"

type Subscription struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Price     int64     `db:"price"`
	Duration  time.Time `db:"duration"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type SubscriptionRepository interface{}

type SubscriptionService interface{}
