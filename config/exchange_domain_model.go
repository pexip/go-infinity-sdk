package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// ExchangeDomain represents an Exchange Metadata Domain configuration
type ExchangeDomain struct {
	ID                int               `json:"id,omitempty"`
	Domain            string            `json:"domain"`
	ExchangeConnector string            `json:"exchange_connector"`
	CreationTime      util.InfinityTime `json:"creation_time,omitempty"`
	ResourceURI       string            `json:"resource_uri,omitempty"`
}

// ExchangeDomainCreateRequest represents a request to create an Exchange Metadata Domain
type ExchangeDomainCreateRequest struct {
	Domain            string `json:"domain"`
	ExchangeConnector string `json:"exchange_connector"`
}

// ExchangeDomainUpdateRequest represents a request to update an Exchange Metadata Domain
type ExchangeDomainUpdateRequest struct {
	Domain            string `json:"domain,omitempty"`
	ExchangeConnector string `json:"exchange_connector,omitempty"`
}

// ExchangeDomainListResponse represents the response from listing Exchange Metadata Domains
type ExchangeDomainListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ExchangeDomain `json:"objects"`
}
