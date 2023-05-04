package model

import "time"

type TicketEventPromotion struct {
	ID        string    `json:"promocode"`
	UserId    int64     `json:"user_id"`
	TicketId  int64     `json:"ticket_id"`
	Method    string    `json:"method"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
