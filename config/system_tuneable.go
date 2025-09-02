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

// ListSystemTuneables retrieves a list of system tuneables
func (s *Service) ListSystemTuneables(ctx context.Context, opts *ListOptions) (*SystemTuneableListResponse, error) {
	endpoint := "configuration/v1/system_tuneable/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result SystemTuneableListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetSystemTuneable retrieves a specific system tuneable by ID
func (s *Service) GetSystemTuneable(ctx context.Context, id int) (*SystemTuneable, error) {
	endpoint := fmt.Sprintf("configuration/v1/system_tuneable/%d/", id)

	var result SystemTuneable
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateSystemTuneable creates a new system tuneable
func (s *Service) CreateSystemTuneable(ctx context.Context, req *SystemTuneableCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/system_tuneable/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateSystemTuneable updates an existing system tuneable
func (s *Service) UpdateSystemTuneable(ctx context.Context, id int, req *SystemTuneableUpdateRequest) (*SystemTuneable, error) {
	endpoint := fmt.Sprintf("configuration/v1/system_tuneable/%d/", id)

	var result SystemTuneable
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteSystemTuneable deletes a system tuneable
func (s *Service) DeleteSystemTuneable(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/system_tuneable/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
