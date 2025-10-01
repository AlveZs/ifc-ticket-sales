package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpHandler "github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/infra/http"
	"github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/infra/repository"
	"github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/infra/service"
	"github.com/alvezs/ifc-ticket-sales/go-backend/internal/events/usecase"
	"github.com/joho/godotenv"
)

type DBInfos struct {
	driverName   string
	host         string
	username     string
	password     string
	port         string
	databaseName string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbInfos := DBInfos{
		driverName:   os.Getenv("DB_DRIVER_NAME"),
		host:         os.Getenv("DB_HOST"),
		username:     os.Getenv("DB_USER"),
		password:     os.Getenv("DB_PASSWORD"),
		port:         os.Getenv("DB_PORT"),
		databaseName: os.Getenv("DB_NAME"),
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbInfos.username,
		dbInfos.password,
		dbInfos.host,
		dbInfos.port,
		dbInfos.databaseName,
	)

	db, err := sql.Open(dbInfos.driverName, connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventRepo, err := repository.NewPgEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseUrls := map[int]string{
		1: os.Getenv("PARTNER1_BASE_URL"),
		2: os.Getenv("PARTNER2_BASE_URL"),
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

	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		log.Println("Received signal interruption, starting graceful shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Error in graceful shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	log.Printf("Server HTTP listening on port %s\n", port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Error starting HTTP server: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("HTTP server closed")
}
