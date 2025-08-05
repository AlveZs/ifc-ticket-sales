package domain

import "errors"

type TicketType string

var ErrTicketPriceZero = errors.New("ticket price must be greater than zero")

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

func isValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (ticket *Ticket) CalculatedPrice() {
	if ticket.TicketType == TicketTypeHalf {
		ticket.Price /= 2
	}
}

func (ticket *Ticket) Validate() error {
	if ticket.Price <= 0 {
		return ErrTicketPriceZero
	}

	return nil
}
