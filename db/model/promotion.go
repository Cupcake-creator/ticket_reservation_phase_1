package model

import "time"

type TicketPromotion struct {
	ID        string    `json:"id"`
	Discount  int64     `json:"discount"`
	Quota     int64     `json:"quota"`
	Airline   string    `json:"airline"`
	Status    string    `json:"status"`
	ExpireAt  time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
