package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListMsExchangeConnectors retrieves a list of Microsoft Exchange connectors
func (s *Service) ListMsExchangeConnectors(ctx context.Context, opts *ListOptions) (*MsExchangeConnectorListResponse, error) {
	endpoint := "configuration/v1/ms_exchange_connector/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MsExchangeConnectorListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMsExchangeConnector retrieves a specific Microsoft Exchange connector by ID
func (s *Service) GetMsExchangeConnector(ctx context.Context, id int) (*MsExchangeConnector, error) {
	endpoint := fmt.Sprintf("configuration/v1/ms_exchange_connector/%d/", id)

	var result MsExchangeConnector
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateMsExchangeConnector creates a new Microsoft Exchange connector
func (s *Service) CreateMsExchangeConnector(ctx context.Context, req *MsExchangeConnectorCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/ms_exchange_connector/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMsExchangeConnector updates an existing Microsoft Exchange connector
func (s *Service) UpdateMsExchangeConnector(ctx context.Context, id int, req *MsExchangeConnectorUpdateRequest) (*MsExchangeConnector, error) {
	endpoint := fmt.Sprintf("configuration/v1/ms_exchange_connector/%d/", id)

	var result MsExchangeConnector
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMsExchangeConnector deletes a Microsoft Exchange connector
func (s *Service) DeleteMsExchangeConnector(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/ms_exchange_connector/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
