package domain

import "time"

type User struct {
	Id           string    `gorm:"primaryKey" db:"id"`
	Name         string    `gorm:"type:varchar(255);not null" db:"name"`
	Email        string    `gorm:"type:varchar(255);unique;not null" db:"email"`
	PasswordHash string    `gorm:"varchar(255);not null" db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}
