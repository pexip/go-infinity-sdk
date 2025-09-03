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

// ListTURNServers retrieves a list of TURN servers
func (s *Service) ListTURNServers(ctx context.Context, opts *ListOptions) (*TURNServerListResponse, error) {
	endpoint := "configuration/v1/turn_server/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result TURNServerListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetTURNServer retrieves a specific TURN server by ID
func (s *Service) GetTURNServer(ctx context.Context, id int) (*TURNServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/turn_server/%d/", id)

	var result TURNServer
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateTURNServer creates a new TURN server
func (s *Service) CreateTURNServer(ctx context.Context, req *TURNServerCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/turn_server/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateTURNServer updates an existing TURN server
func (s *Service) UpdateTURNServer(ctx context.Context, id int, req *TURNServerUpdateRequest) (*TURNServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/turn_server/%d/", id)

	var result TURNServer
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteTURNServer deletes a TURN server
func (s *Service) DeleteTURNServer(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/turn_server/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
