package domain

import "time"

type Transaction struct {
	Id                string    `json:"id"`
	UserId            string    `json:"user_id"`
	SubscriptionId    int       `json:"subscription_id"`
	Amount            float64   `json:"amount"`
	Status            string    `json:"status"`
	PaymentType       string    `json:"payment_type"`
	MidtransSnapToken string    `json:"midtrans_snap_token"`
	MidtransSnapURL   string    `json:"midtrans_snap_url"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
