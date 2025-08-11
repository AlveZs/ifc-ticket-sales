package usecase

type SpotDTO struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	EventId  string `json:"event_id"`
	Reserved bool   `json:"reserved"`
	Status   string `json:"status"`
	TicketId string `json:"ticket_id"`
}

type EventDTO struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	ImageUrl     string  `json:"image_url"`
	Capacity     int     `json:"capacity"`
	Price        float64 `json:"price"`
	PartnerId    int     `json:"partner_id"`
}

type TicketDTO struct {
	Id         string  `json:"id"`
	SpotId     string  `json:"spot_id"`
	TicketType string  `json:"ticket_type"`
	Price      float64 `json:"price"`
}
