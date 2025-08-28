/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListExchangeDomains retrieves a list of Exchange Metadata Domains
func (s *Service) ListExchangeDomains(ctx context.Context, opts *ListOptions) (*ExchangeDomainListResponse, error) {
	endpoint := "configuration/v1/exchange_domain/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ExchangeDomainListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetExchangeDomain retrieves a specific Exchange Metadata Domain by ID
func (s *Service) GetExchangeDomain(ctx context.Context, id int) (*ExchangeDomain, error) {
	endpoint := fmt.Sprintf("configuration/v1/exchange_domain/%d/", id)

	var result ExchangeDomain
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateExchangeDomain creates a new Exchange Metadata Domain
func (s *Service) CreateExchangeDomain(ctx context.Context, req *ExchangeDomainCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/exchange_domain/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateExchangeDomain updates an existing Exchange Metadata Domain
func (s *Service) UpdateExchangeDomain(ctx context.Context, id int, req *ExchangeDomainUpdateRequest) (*ExchangeDomain, error) {
	endpoint := fmt.Sprintf("configuration/v1/exchange_domain/%d/", id)

	var result ExchangeDomain
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteExchangeDomain deletes an Exchange Metadata Domain
func (s *Service) DeleteExchangeDomain(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/exchange_domain/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
