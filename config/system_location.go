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

// ListSystemLocations retrieves a list of system locations
func (s *Service) ListSystemLocations(ctx context.Context, opts *ListOptions) (*SystemLocationListResponse, error) {
	endpoint := "configuration/v1/system_location/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result SystemLocationListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetSystemLocation retrieves a specific system location by ID
func (s *Service) GetSystemLocation(ctx context.Context, id int) (*SystemLocation, error) {
	endpoint := fmt.Sprintf("configuration/v1/system_location/%d/", id)

	var result SystemLocation
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateSystemLocation creates a new system location
func (s *Service) CreateSystemLocation(ctx context.Context, req *SystemLocationCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/system_location/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateSystemLocation updates an existing system location
func (s *Service) UpdateSystemLocation(ctx context.Context, id int, req *SystemLocationUpdateRequest) (*SystemLocation, error) {
	endpoint := fmt.Sprintf("configuration/v1/system_location/%d/", id)

	var result SystemLocation
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteSystemLocation deletes a system location
func (s *Service) DeleteSystemLocation(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/system_location/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
