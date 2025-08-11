package usecase

import "github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/domain"

type GetEventInputDTO struct {
	Id string `json:"id"`
}

type GetEventOutputDTO struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	Capacity     int     `json:"capacity"`
	Price        float64 `json:"price"`
	PartnerId    int     `json:"partner_id"`
}

type GetEventUseCase struct {
	repo domain.EventRepository
}

func NewGetEventUseCase(repo domain.EventRepository) *GetEventUseCase {
	return &GetEventUseCase{repo: repo}
}

func (useCase *GetEventUseCase) Execute(input GetEventInputDTO) (*GetEventOutputDTO, error) {
	event, err := useCase.repo.FindEventByID(input.Id)
	if err != nil {
		return nil, err
	}

	return &GetEventOutputDTO{
		Id:           event.Id,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date.Format("2006-01-2 15:04:05"),
		Capacity:     event.Capacity,
		Price:        event.Price,
		PartnerId:    event.PartnerId,
	}, nil
}
