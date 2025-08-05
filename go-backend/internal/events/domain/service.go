package domain

import (
	"errors"
	"fmt"
)

type spotService struct{}

var ErrInvalidQuantity = errors.New("quantity must be greater than zero")

func NewSpotService() *spotService {
	return &spotService{}
}

func (spotService *spotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	for i := range quantity {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)
		spot, spotCreationError := NewSpot(event, spotName)
		if spotCreationError != nil {
			return spotCreationError
		}
		event.Spots = append(event.Spots, *spot)
	}

	return nil
}
