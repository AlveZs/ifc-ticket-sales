package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TicketType string

var (
	ErrTicketPriceZero = errors.New("ticket price must be greater than zero")
	ErrInvalidType     = errors.New("invalid ticket type")
)

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

func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !isValidTicketType(ticketType) {
		return nil, ErrInvalidType
	}
	ticket := &Ticket{
		Id:         uuid.NewString(),
		EventId:    event.Id,
		Spot:       spot,
		TicketType: ticketType,
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}

	return ticket, nil
}

func isValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (ticket *Ticket) CalculatePrice() {
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
