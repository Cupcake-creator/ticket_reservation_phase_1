package model

import "time"

type Ticket struct {
	ID          int64           `json:"id"`
	OwnerId     string          `json:"owner_id"`
	Airline     string          `json:"airline"`
	From        string          `json:"from"`
	Destination string          `json:"destination"`
	Promotion   TicketPromotion `json:"ticket_promotion"`
	Status      string          `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
