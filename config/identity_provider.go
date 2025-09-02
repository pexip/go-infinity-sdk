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

// ListIdentityProviders retrieves a list of identity providers
func (s *Service) ListIdentityProviders(ctx context.Context, opts *ListOptions) (*IdentityProviderListResponse, error) {
	endpoint := "configuration/v1/identity_provider/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result IdentityProviderListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetIdentityProvider retrieves a specific identity provider by ID
func (s *Service) GetIdentityProvider(ctx context.Context, id int) (*IdentityProvider, error) {
	endpoint := fmt.Sprintf("configuration/v1/identity_provider/%d/", id)

	var result IdentityProvider
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateIdentityProvider creates a new identity provider
func (s *Service) CreateIdentityProvider(ctx context.Context, req *IdentityProviderCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/identity_provider/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateIdentityProvider updates an existing identity provider
func (s *Service) UpdateIdentityProvider(ctx context.Context, id int, req *IdentityProviderUpdateRequest) (*IdentityProvider, error) {
	endpoint := fmt.Sprintf("configuration/v1/identity_provider/%d/", id)

	var result IdentityProvider
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteIdentityProvider deletes an identity provider
func (s *Service) DeleteIdentityProvider(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/identity_provider/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
