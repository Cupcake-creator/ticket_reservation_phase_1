package model

import "time"

type Ticket struct {
	ID           int64           `json:"id"`
	OwnerId      string          `json:"owner_id"`
	Airline      string          `json:"airline"`
	FromLocation string          `json:"from_location"`
	Destination  string          `json:"destination"`
	Promotion    TicketPromotion `json:"ticket_promotion"` // TODO how to get promotion json
	Price        float64         `json:"price"`
	Status       string          `json:"status"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}
