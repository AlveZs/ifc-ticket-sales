package usecase

import (
	"github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/domain"
	"github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/infra/service"
)

type BuyTicketsInputDTO struct {
	EventId    string   `json:"event_id"`
	Spots      []string `json:"spots"`
	TicketType string   `json:"ticket_type"`
	CardHash   string   `json:"card_hash"`
	Email      string   `json:"email"`
}

type BuyTicketsOutputDTO struct {
	Tickets []TicketDTO `json:"tickets"`
}

type BuyTicketsUseCase struct {
	repo           domain.EventRepository
	partnerFactory service.PartnerFactory
}

func NewBuyTicketsUseCase(repo domain.EventRepository, partnerFactory service.PartnerFactory) *BuyTicketsUseCase {
	return &BuyTicketsUseCase{repo: repo, partnerFactory: partnerFactory}
}

func (useCase *BuyTicketsUseCase) Execute(input BuyTicketsInputDTO) (*BuyTicketsOutputDTO, error) {
	event, err := useCase.repo.FindEventByID(input.EventId)
	if err != nil {
		return nil, err
	}

	request := &service.ReservationRequest{
		EventId:    input.EventId,
		Spots:      input.Spots,
		TicketType: input.TicketType,
		CardHash:   input.CardHash,
		Email:      input.Email,
	}

	partnerService, err := useCase.partnerFactory.CreatePartner(event.PartnerId)
	if err != nil {
		return nil, err
	}

	reservationResponse, err := partnerService.MakeReservation(request)
	if err != nil {
		return nil, err
	}

	tickets := make([]domain.Ticket, len(reservationResponse))
	for i, reservation := range reservationResponse {
		spot, err := useCase.repo.FindSpotByName(event.Id, reservation.Spot)
		if err != nil {
			return nil, err
		}

		ticket, err := domain.NewTicket(event, spot, domain.TicketType(reservation.TicketType))
		if err != nil {
			return nil, err
		}

		err = useCase.repo.CreateTicket(ticket)
		if err != nil {
			return nil, err
		}

		spot.Reserve(ticket.Id)
		err = useCase.repo.ReserveSpot(spot.Id, ticket.Id)
		if err != nil {
			return nil, err
		}

		tickets[i] = *ticket
	}

	ticketDTOs := make([]TicketDTO, len(tickets))
	for i, ticket := range tickets {
		ticketDTOs[i] = TicketDTO{
			Id:         ticket.Id,
			SpotId:     ticket.Spot.Id,
			TicketType: string(ticket.TicketType),
			Price:      ticket.Price,
		}
	}

	return &BuyTicketsOutputDTO{Tickets: ticketDTOs}, nil
}
