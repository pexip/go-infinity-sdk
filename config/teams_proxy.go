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

// ListTeamsProxies retrieves a list of Teams proxies
func (s *Service) ListTeamsProxies(ctx context.Context, opts *ListOptions) (*TeamsProxyListResponse, error) {
	endpoint := "configuration/v1/teams_proxy/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result TeamsProxyListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetTeamsProxy retrieves a specific Teams proxy by ID
func (s *Service) GetTeamsProxy(ctx context.Context, id int) (*TeamsProxy, error) {
	endpoint := fmt.Sprintf("configuration/v1/teams_proxy/%d/", id)

	var result TeamsProxy
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateTeamsProxy creates a new Teams proxy
func (s *Service) CreateTeamsProxy(ctx context.Context, req *TeamsProxyCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/teams_proxy/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateTeamsProxy updates an existing Teams proxy
func (s *Service) UpdateTeamsProxy(ctx context.Context, id int, req *TeamsProxyUpdateRequest) (*TeamsProxy, error) {
	endpoint := fmt.Sprintf("configuration/v1/teams_proxy/%d/", id)

	var result TeamsProxy
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteTeamsProxy deletes a Teams proxy
func (s *Service) DeleteTeamsProxy(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/teams_proxy/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
