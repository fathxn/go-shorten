package dto

import (
	"time"
)

type (
	URLInputRequest struct {
		LongURL string `json:"long_url" validate:"required" binding:"required"`
	}

	GetByUserIdRequest struct {
		UserId string `json:"user_id"`
	}

	URLResponse struct {
		LongURL   string    `json:"long_url"`
		ShortURL  string    `json:"short_url"`
		CreatedAt time.Time `json:"created_at"`
	}

	URLResponseByUserId struct {
		Id        int       `json:"id"`
		UserId    string    `json:"user_id"`
		LongURL   string    `json:"long_url"`
		ShortURL  string    `json:"short_url"`
		CreatedAt time.Time `json:"created_at"`
	}
)
