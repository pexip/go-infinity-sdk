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

// ListH323Gatekeepers retrieves a list of H.323 gatekeepers
func (s *Service) ListH323Gatekeepers(ctx context.Context, opts *ListOptions) (*H323GatekeeperListResponse, error) {
	endpoint := "configuration/v1/h323_gatekeeper/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result H323GatekeeperListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetH323Gatekeeper retrieves a specific H.323 gatekeeper by ID
func (s *Service) GetH323Gatekeeper(ctx context.Context, id int) (*H323Gatekeeper, error) {
	endpoint := fmt.Sprintf("configuration/v1/h323_gatekeeper/%d/", id)

	var result H323Gatekeeper
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateH323Gatekeeper creates a new H.323 gatekeeper
func (s *Service) CreateH323Gatekeeper(ctx context.Context, req *H323GatekeeperCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/h323_gatekeeper/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateH323Gatekeeper updates an existing H.323 gatekeeper
func (s *Service) UpdateH323Gatekeeper(ctx context.Context, id int, req *H323GatekeeperUpdateRequest) (*H323Gatekeeper, error) {
	endpoint := fmt.Sprintf("configuration/v1/h323_gatekeeper/%d/", id)

	var result H323Gatekeeper
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteH323Gatekeeper deletes an H.323 gatekeeper
func (s *Service) DeleteH323Gatekeeper(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/h323_gatekeeper/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
