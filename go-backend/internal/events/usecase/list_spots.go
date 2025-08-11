package usecase

import "github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/domain"

type ListSpotsInputDTO struct {
	EventId string `json:"event_id"`
}

type ListSpotsOutputDTO struct {
	Event EventDTO  `json:"id"`
	Spots []SpotDTO `json:"spots"`
}

type ListSpotsUseCase struct {
	repo domain.EventRepository
}

func NewListSpotsUseCase(repo domain.EventRepository) *ListSpotsUseCase {
	return &ListSpotsUseCase{repo: repo}
}

func (useCase *ListSpotsUseCase) Execute(input ListSpotsInputDTO) (*ListSpotsOutputDTO, error) {
	event, err := useCase.repo.FindEventByID(input.EventId)
	if err != nil {
		return nil, err
	}

	spots, err := useCase.repo.FindSpotsByEventID(input.EventId)
	if err != nil {
		return nil, err
	}

	spotDTOs := make([]SpotDTO, len(spots))
	for i, spot := range spots {
		spotDTOs[i] = SpotDTO{
			Id:       spot.Id,
			EventId:  spot.EventId,
			Name:     spot.Name,
			Status:   string(spot.Status),
			TicketId: spot.TicketId,
		}
	}

	eventDTO := EventDTO{
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

	return &ListSpotsOutputDTO{Event: eventDTO, Spots: spotDTOs}, nil
}
