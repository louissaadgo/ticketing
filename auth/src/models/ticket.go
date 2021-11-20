package models

type Ticket struct {
	TicketID string `json:"ticketid" bson:"ticketid"`
	Title    string `json:"title" bson:"title"`
	Price    string `json:"price" bson:"price"`
	UserID   string `json:"userid" bson:"userid"`
}
