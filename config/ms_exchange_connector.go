/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListMsExchangeConnectors retrieves a list of Microsoft Exchange connectors
func (s *Service) ListMsExchangeConnectors(ctx context.Context, opts *ListOptions) (*MsExchangeConnectorListResponse, error) {
	endpoint := "configuration/v1/ms_exchange_connector/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result MsExchangeConnectorListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetMsExchangeConnector retrieves a specific Microsoft Exchange connector by ID
func (s *Service) GetMsExchangeConnector(ctx context.Context, id int) (*MsExchangeConnector, error) {
	endpoint := fmt.Sprintf("configuration/v1/ms_exchange_connector/%d/", id)

	var result MsExchangeConnector
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
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
