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

// ListSTUNServers retrieves a list of STUN servers
func (s *Service) ListSTUNServers(ctx context.Context, opts *ListOptions) (*STUNServerListResponse, error) {
	endpoint := "configuration/v1/stun_server/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result STUNServerListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetSTUNServer retrieves a specific STUN server by ID
func (s *Service) GetSTUNServer(ctx context.Context, id int) (*STUNServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/stun_server/%d/", id)

	var result STUNServer
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateSTUNServer creates a new STUN server
func (s *Service) CreateSTUNServer(ctx context.Context, req *STUNServerCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/stun_server/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateSTUNServer updates an existing STUN server
func (s *Service) UpdateSTUNServer(ctx context.Context, id int, req *STUNServerUpdateRequest) (*STUNServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/stun_server/%d/", id)

	var result STUNServer
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteSTUNServer deletes a STUN server
func (s *Service) DeleteSTUNServer(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/stun_server/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
