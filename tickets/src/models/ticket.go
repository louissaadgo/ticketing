package models

import (
	"github.com/google/uuid"
)

type Ticket struct {
	TicketID uuid.UUID `json:"ticketid" bson:"ticketid"`
	Title    string    `json:"title" bson:"title"`
	Price    string    `json:"price" bson:"price"`
	UserID   string    `json:"userid" bson:"userid"`
}

func (ticket *Ticket) Validate() bool {
	return true
}
