package usecase

import "github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/domain"

type ListEventsOutputDTO struct {
	Events []EventDTO `json:"events"`
}

type ListEventsUseCase struct {
	repo domain.EventRepository
}

func NewListEventsUseCase(repo domain.EventRepository) *ListEventsUseCase {
	return &ListEventsUseCase{repo: repo}
}

func (useCase *ListEventsUseCase) Execute() (*ListEventsOutputDTO, error) {
	events, err := useCase.repo.ListEvents()
	if err != nil {
		return nil, err
	}

	eventDTOs := make([]EventDTO, len(events))
	for i, event := range events {
		eventDTOs[i] = EventDTO{
			Id:           event.Id,
			Name:         event.Name,
			Location:     event.Location,
			Organization: event.Organization,
			Rating:       string(event.Rating),
			Date:         event.Date.Format("2006-01-2 15:04:05"),
			ImageUrl:     event.ImageUrl,
			Capacity:     event.Capacity,
			Price:        event.Price,
			PartnerId:    event.PartnerId,
		}
	}

	return &ListEventsOutputDTO{Events: eventDTOs}, nil
}
