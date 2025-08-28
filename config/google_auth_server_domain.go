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

// ListGoogleAuthServerDomains retrieves a list of Google OAuth 2.0 Credential domains
func (s *Service) ListGoogleAuthServerDomains(ctx context.Context, opts *ListOptions) (*GoogleAuthServerDomainListResponse, error) {
	endpoint := "configuration/v1/google_auth_server_domain/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result GoogleAuthServerDomainListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetGoogleAuthServerDomain retrieves a specific Google OAuth 2.0 Credential domain by ID
func (s *Service) GetGoogleAuthServerDomain(ctx context.Context, id int) (*GoogleAuthServerDomain, error) {
	endpoint := fmt.Sprintf("configuration/v1/google_auth_server_domain/%d/", id)

	var result GoogleAuthServerDomain
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateGoogleAuthServerDomain creates a new Google OAuth 2.0 Credential domain
func (s *Service) CreateGoogleAuthServerDomain(ctx context.Context, req *GoogleAuthServerDomainCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/google_auth_server_domain/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateGoogleAuthServerDomain updates an existing Google OAuth 2.0 Credential domain
func (s *Service) UpdateGoogleAuthServerDomain(ctx context.Context, id int, req *GoogleAuthServerDomainUpdateRequest) (*GoogleAuthServerDomain, error) {
	endpoint := fmt.Sprintf("configuration/v1/google_auth_server_domain/%d/", id)

	var result GoogleAuthServerDomain
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteGoogleAuthServerDomain deletes a Google OAuth 2.0 Credential domain
func (s *Service) DeleteGoogleAuthServerDomain(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/google_auth_server_domain/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
