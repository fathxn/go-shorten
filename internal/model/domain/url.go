package domain

import (
	"context"
	"time"
)

type URL struct {
	Id        int       `db:"id" gorm:"primaryKey;autoIncrement;type:int"`
	UserId    string    `db:"user_id" gorm:"type:varchar(255)"`
	LongURL   string    `db:"long_url" gorm:"type:varchar(255)"`
	ShortCode string    `db:"short_code" gorm:"type:varchar(255)"`
	Visits    int       `db:"visits"`
	CreatedAt time.Time `db:"created_at"`
}

type URLRepository interface {
	Insert(ctx context.Context, url *URL) error
	FindByShortCode(ctx context.Context, shortCode string) (*URL, error)
	FindById(ctx context.Context, id int) (*URL, error)
	FindByUserId(ctx context.Context, userId string) ([]URL, error)
	Delete(ctx context.Context, id int) error
}

type URLService interface {
	Create(ctx context.Context, longURL string, userId string) (*URL, error)
	GetLongURL(ctx context.Context, shortCode string) (*URL, error)
	GetById(ctx context.Context, id int) (*URL, error)
	// GetByUserId(ctx context.Context, userId string) (*[]domain.URL, error)
	Delete(ctx context.Context, id int) error
}
