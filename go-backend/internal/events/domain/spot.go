package domain

import (
	"errors"

	"github.com/google/uuid"
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

var (
	ErrSpotMinLength       = errors.New("Spot name must be at least 2 characters long")
	ErrSpotStartNotAlpha   = errors.New("Spot name must start with a letter")
	ErrSpotEndNotNumber    = errors.New("Spot name must end with a number")
	ErrInvalidSpotNumber   = errors.New("Invalid spot number")
	ErrSpotNotFound        = errors.New("Spot not found")
	ErrSpotAlreadyReserved = errors.New("Spot already reserved")
)

type Spot struct {
	Id       string
	EventId  string
	Name     string
	Status   SpotStatus
	TicketId string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		Id:      uuid.NewString(),
		EventId: event.Id,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if spotValidationError := spot.Validate(); spotValidationError != nil {
		return nil, spotValidationError
	}

	return spot, nil
}

func (spot Spot) Validate() error {
	if spot.Name == "" {
		return ErrEventNameRequired
	}

	if len(spot.Name) < 2 {
		return ErrSpotMinLength
	}

	if spot.Name[0] < 'A' || spot.Name[0] > 'Z' {
		return ErrSpotStartNotAlpha
	}

	if spot.Name[1] < '0' || spot.Name[1] > '9' {
		return ErrSpotEndNotNumber
	}

	return nil
}
