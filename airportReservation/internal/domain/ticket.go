package domain

type Ticket struct {
	TicketID string `json:"ticket_id"`
	Chair    Chair  `json:"chair"`
	UserID   int    `json:"user_id"`
	Payment  string `json:"payment"`
}
