package models

type Ticket struct {
	Title  string `json:"title"`
	Price  string `json:"price"`
	UserID string `json:"userid"`
}

func (ticket *Ticket) Validate() bool {
	return true
}
