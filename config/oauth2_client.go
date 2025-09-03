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

// ListOAuth2Clients retrieves a list of OAuth2 clients
func (s *Service) ListOAuth2Clients(ctx context.Context, opts *ListOptions) (*OAuth2ClientListResponse, error) {
	endpoint := "configuration/v1/oauth2_client/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result OAuth2ClientListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetOAuth2Client retrieves a specific OAuth2 client by client ID
func (s *Service) GetOAuth2Client(ctx context.Context, clientID string) (*OAuth2Client, error) {
	endpoint := fmt.Sprintf("configuration/v1/oauth2_client/%s/", clientID)

	var result OAuth2Client
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateOAuth2Client creates a new OAuth2 client
func (s *Service) CreateOAuth2Client(ctx context.Context, req *OAuth2ClientCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/oauth2_client/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateOAuth2Client updates an existing OAuth2 client
func (s *Service) UpdateOAuth2Client(ctx context.Context, clientID string, req *OAuth2ClientUpdateRequest) (*OAuth2Client, error) {
	endpoint := fmt.Sprintf("configuration/v1/oauth2_client/%s/", clientID)

	var result OAuth2Client
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteOAuth2Client deletes an OAuth2 client
func (s *Service) DeleteOAuth2Client(ctx context.Context, clientID string) error {
	endpoint := fmt.Sprintf("configuration/v1/oauth2_client/%s/", clientID)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
