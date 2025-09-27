package model

import (
	"time"

	"github.com/uptrace/bun"
)

type UrlStatistic struct {
	bun.BaseModel `bun:"table:url.url_statistics"`

	ID         string    `bun:"id,pk,notnull"`
	UrlID      string    `bun:"url_id"`
	VisitCount int64     `bun:"visit_count"`
	CreatedAt  time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt  time.Time `bun:"updated_at,default:current_timestamp"`

	Url *Url `bun:"rel:belongs-to,join:url_id=id"`
}
