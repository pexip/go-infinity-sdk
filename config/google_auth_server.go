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

// ListGoogleAuthServers retrieves a list of Google OAuth 2.0 Credentials
func (s *Service) ListGoogleAuthServers(ctx context.Context, opts *ListOptions) (*GoogleAuthServerListResponse, error) {
	endpoint := "configuration/v1/google_auth_server/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result GoogleAuthServerListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetGoogleAuthServer retrieves a specific Google OAuth 2.0 Credential by ID
func (s *Service) GetGoogleAuthServer(ctx context.Context, id int) (*GoogleAuthServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/google_auth_server/%d/", id)

	var result GoogleAuthServer
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateGoogleAuthServer creates a new Google OAuth 2.0 Credential
func (s *Service) CreateGoogleAuthServer(ctx context.Context, req *GoogleAuthServerCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/google_auth_server/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateGoogleAuthServer updates an existing Google OAuth 2.0 Credential
func (s *Service) UpdateGoogleAuthServer(ctx context.Context, id int, req *GoogleAuthServerUpdateRequest) (*GoogleAuthServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/google_auth_server/%d/", id)

	var result GoogleAuthServer
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteGoogleAuthServer deletes a Google OAuth 2.0 Credential
func (s *Service) DeleteGoogleAuthServer(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/google_auth_server/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
