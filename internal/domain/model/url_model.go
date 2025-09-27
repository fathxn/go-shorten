package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Url struct {
	bun.BaseModel `bun:"table:url.urls"`

	ID        string    `bun:"id,pk,notnull"`
	LongUrl   string    `bun:"long_url"`
	ShortCode string    `bun:"short_code"`
	CreatedBy string    `bun:"created_by"`
	CreatedAt time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,default:current_timestamp"`
	DeletedAt time.Time `bun:"deleted_at,soft_delete,nullzero"`
}
