package domain

type TicketType string

const (
	TicketTypeHalf TicketType = "half"
	TicketTypeFull TicketType = "full"
)

type Ticket struct {
	Id         string
	EventId    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}
