package api

import "time"

type URLInputRequest struct {
	LongURL string `json:"long_url"`
}

type GetByUserIdRequest struct {
	UserId string `json:"user_id"`
}

type URLResponse struct {
	LongURL   string    `json:"long_url"`
	ShortURL  string    `json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
}
