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

// ListLocations retrieves a list of locations
func (s *Service) ListLocations(ctx context.Context, opts *ListOptions) (*LocationListResponse, error) {
	endpoint := "configuration/v1/location/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result LocationListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetLocation retrieves a specific location by ID
func (s *Service) GetLocation(ctx context.Context, id int) (*Location, error) {
	endpoint := fmt.Sprintf("configuration/v1/location/%d/", id)

	var result Location
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateLocation creates a new location
func (s *Service) CreateLocation(ctx context.Context, req *LocationCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/location/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateLocation updates an existing location
func (s *Service) UpdateLocation(ctx context.Context, id int, req *LocationUpdateRequest) (*Location, error) {
	endpoint := fmt.Sprintf("configuration/v1/location/%d/", id)

	var result Location
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteLocation deletes a location
func (s *Service) DeleteLocation(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/location/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
