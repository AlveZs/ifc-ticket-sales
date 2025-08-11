package http

import (
	"encoding/json"
	"net/http"

	"github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/usecase"
)

type EventsHandler struct {
	listEventsUseCase *usecase.ListEventsUseCase
	ListSpotsUseCase  *usecase.ListSpotsUseCase
	getEventUseCase   *usecase.GetEventUseCase
	buyTicketsUseCase *usecase.BuyTicketsUseCase
}

func NewEventHandler(
	listEventsUseCase *usecase.ListEventsUseCase,
	listSpotsUseCase *usecase.ListSpotsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	buyTicketsUseCase *usecase.BuyTicketsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase: listEventsUseCase,
		ListSpotsUseCase:  listSpotsUseCase,
		buyTicketsUseCase: buyTicketsUseCase,
		getEventUseCase:   getEventUseCase,
	}
}

func (handler *EventsHandler) ListEvents(writer http.ResponseWriter, request *http.Request) {
	output, err := handler.listEventsUseCase.Execute()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(output)
}

func (handler *EventsHandler) GetEvent(writer http.ResponseWriter, request *http.Request) {
	eventId := request.PathValue("eventId")
	input := usecase.GetEventInputDTO{Id: eventId}

	output, err := handler.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(output)
}

func (handler *EventsHandler) ListSpots(writer http.ResponseWriter, request *http.Request) {
	eventId := request.PathValue("eventId")
	input := usecase.ListSpotsInputDTO{EventId: eventId}

	output, err := handler.ListSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(output)
}

func (handler *EventsHandler) BuyTickets(writer http.ResponseWriter, request *http.Request) {
	var input usecase.BuyTicketsInputDTO
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := handler.buyTicketsUseCase.Execute(input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(output)
}
