package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNameRequired = errors.New("Event name is required")
	ErrEventPastDate     = errors.New("Event date must be in the future")
	ErrEventCapacityZero = errors.New("Event capacity must be greater than zero")
	ErrEventPriceZero    = errors.New("Event price must be greater than zero")
)

type Rating string

const (
	RatingG    Rating = "G"
	RatingPG   Rating = "PG"
	RatingPG13 Rating = "PG-13"
	RatingR    Rating = "R"
	RatingNC17 Rating = "NC-17"
)

type Event struct {
	Id           string
	Name         string
	Location     string
	Organization string
	Rating       string
	Date         time.Time
	ImageUrl     string
	Capacity     int
	Price        float64
	PartnerId    int
	Spots        []Spot
	Tickets      []Ticket
}

func (event Event) Validate() error {
	if event.Name == "" {
		return ErrEventNameRequired
	}

	if event.Date.Before(time.Now()) {
		return ErrEventPastDate
	}

	if event.Capacity <= 0 {
		return ErrEventCapacityZero
	}

	if event.Price <= 0 {
		return ErrEventPriceZero
	}

	return nil
}

func (event *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(event, name)
	if err != nil {
		return nil, err
	}
	event.Spots = append(event.Spots, *spot)
	return spot, nil
}
