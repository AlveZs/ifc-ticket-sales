package events

import (
	"database/sql"
	"net/http"

	httpHandler "github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/infra/http"
	"github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/infra/repository"
	"github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/infra/service"
	"github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/usecase"
)

func main() {
	connStr := "postgres://username:password@localhost/eventsdb"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventRepo, err := repository.NewPgEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseUrls := map[int]string{
		1: "http://łocalhost:9080/api1",
		2: "http://localhost:9080/api2",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseUrls)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	buyTicketUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	eventsHandler := httpHandler.NewEventHandler(
		listEventsUseCase,
		listSpotsUseCase,
		getEventUseCase,
		buyTicketUseCase,
	)

	router := http.NewServeMux()
	router.HandleFunc("/events", eventsHandler.ListEvents)
	router.HandleFunc("/events/{eventId}/spots", eventsHandler.ListSpots)
	router.HandleFunc("/events/{eventId}", eventsHandler.GetEvent)
	router.HandleFunc("POST /checkout", eventsHandler.BuyTickets)

	http.ListenAndServe(":8080", router)
}
