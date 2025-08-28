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

// ListMjxExchangeAutodiscoverURLs retrieves a list of MJX Exchange autodiscover URLs
func (s *Service) ListMjxExchangeAutodiscoverURLs(ctx context.Context, opts *ListOptions) (*MjxExchangeAutodiscoverURLListResponse, error) {
	endpoint := "configuration/v1/mjx_exchange_autodiscover_url/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MjxExchangeAutodiscoverURLListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMjxExchangeAutodiscoverURL retrieves a specific MJX Exchange autodiscover URL by ID
func (s *Service) GetMjxExchangeAutodiscoverURL(ctx context.Context, id int) (*MjxExchangeAutodiscoverURL, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_exchange_autodiscover_url/%d/", id)

	var result MjxExchangeAutodiscoverURL
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateMjxExchangeAutodiscoverURL creates a new MJX Exchange autodiscover URL
func (s *Service) CreateMjxExchangeAutodiscoverURL(ctx context.Context, req *MjxExchangeAutodiscoverURLCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/mjx_exchange_autodiscover_url/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMjxExchangeAutodiscoverURL updates an existing MJX Exchange autodiscover URL
func (s *Service) UpdateMjxExchangeAutodiscoverURL(ctx context.Context, id int, req *MjxExchangeAutodiscoverURLUpdateRequest) (*MjxExchangeAutodiscoverURL, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_exchange_autodiscover_url/%d/", id)

	var result MjxExchangeAutodiscoverURL
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMjxExchangeAutodiscoverURL deletes a MJX Exchange autodiscover URL
func (s *Service) DeleteMjxExchangeAutodiscoverURL(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/mjx_exchange_autodiscover_url/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
