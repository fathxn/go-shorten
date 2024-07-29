package domain

import (
	"context"
	"time"
)

type User struct {
	Id           string    `gorm:"primaryKey" db:"id"`
	Name         string    `gorm:"type:varchar(255);not null" db:"name"`
	Email        string    `gorm:"type:varchar(255);unique;not null" db:"email"`
	PasswordHash string    `gorm:"varchar(255);not null" db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

type UserRepository interface {
	Insert(ctx context.Context, user *User) error
	FindById(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Delete(ctx context.Context, id string) error
}

type UserService interface {
	GetById(ctx context.Context, id string) (*User, error)
	GetURLsByUserId(ctx context.Context, userId string) (*[]URL, error)
	Delete(ctx context.Context, id string) error
}
