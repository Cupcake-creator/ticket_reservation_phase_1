package model

import "time"

type TicketPromotion struct {
	PromoCode   string    `json:"promocode"`
	Price       float64   `json:"price"`
	ValidUnitil time.Time `json:"validunitil"`
	Status      string    `json:"status"`
	Limit       int64     `json:"limit"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
