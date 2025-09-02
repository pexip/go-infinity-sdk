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

// ListMjxEndpoints retrieves a list of MJX endpoints
func (s *Service) ListMjxEndpoints(ctx context.Context, opts *ListOptions) (*MjxEndpointListResponse, error) {
	endpoint := "configuration/v1/mjx_endpoint/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result MjxEndpointListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetMjxEndpoint retrieves a specific MJX endpoint by ID
func (s *Service) GetMjxEndpoint(ctx context.Context, id int) (*MjxEndpoint, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_endpoint/%d/", id)

	var result MjxEndpoint
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateMjxEndpoint creates a new MJX endpoint
func (s *Service) CreateMjxEndpoint(ctx context.Context, req *MjxEndpointCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/mjx_endpoint/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMjxEndpoint updates an existing MJX endpoint
func (s *Service) UpdateMjxEndpoint(ctx context.Context, id int, req *MjxEndpointUpdateRequest) (*MjxEndpoint, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_endpoint/%d/", id)

	var result MjxEndpoint
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMjxEndpoint deletes a MJX endpoint
func (s *Service) DeleteMjxEndpoint(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/mjx_endpoint/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
