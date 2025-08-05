package service

import "fmt"

type PartnerFactory interface {
	CreatePartner(partnerId int) (Partner, error)
}

type DefaultPartnerFactory struct {
	partnerBaseURLs map[int]string
}

func NewPartnerFactory(partnerBaseURLs map[int]string) PartnerFactory {
	return &DefaultPartnerFactory{partnerBaseURLs: partnerBaseURLs}
}

func (factory *DefaultPartnerFactory) CreatePartner(partnerId int) (Partner, error) {
	baseURL, ok := factory.partnerBaseURLs[partnerId]
	if !ok {
		return nil, fmt.Errorf("partner with ID %d not found", partnerId)
	}

	switch partnerId {
	case 1:
		return &Partner1{BaseURL: baseURL}, nil
	case 2:
		return &Partner2{BaseURL: baseURL}, nil
	default:
		return nil, fmt.Errorf("partner with ID %d does not exists", partnerId)
	}
}
