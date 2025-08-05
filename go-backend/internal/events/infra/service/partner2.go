package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner2 struct {
	BaseURL string
}

type Partner2ReservationRequest struct {
	Spots      []string `json:"lugares"`
	TicketKind string   `json:"tipo_ingresso"`
	Email      string   `json:"email"`
}

type Partner2ReservationResponse struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"lugar"`
	TicketKind string `json:"tipo_ingresso"`
	Status     string `json:"estado"`
	EventId    string `json:"evento_id"`
}

func (partner *Partner2) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	partnerReq := Partner2ReservationRequest{
		Spots:      req.Spots,
		TicketKind: req.TicketKind,
		Email:      req.Email,
	}

	body, err := json.Marshal(partnerReq)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/eventos/%s/reservar", partner.BaseURL, req.EventId)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}

	var partnerResponse []Partner2ReservationResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&partnerResponse); err != nil {
		return nil, err
	}

	responses := make([]ReservationResponse, len(partnerResponse))
	for i, r := range partnerResponse {
		responses[i] = ReservationResponse{
			Id:     r.Id,
			Spot:   r.Spot,
			Status: r.Status,
		}
	}

	return responses, err
}
