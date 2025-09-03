/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListWebappBrandings retrieves a list of webapp brandings
func (s *Service) ListWebappBrandings(ctx context.Context, opts *ListOptions) (*WebappBrandingListResponse, error) {
	endpoint := "configuration/v1/webapp_branding/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result WebappBrandingListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetWebappBranding retrieves a specific webapp branding by name
func (s *Service) GetWebappBranding(ctx context.Context, name string) (*WebappBranding, error) {
	endpoint := "configuration/v1/webapp_branding/" + name + "/"

	var result WebappBranding
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateWebappBranding creates a new webapp branding
func (s *Service) CreateWebappBranding(ctx context.Context, req *WebappBrandingCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/webapp_branding/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateWebappBranding updates an existing webapp branding
func (s *Service) UpdateWebappBranding(ctx context.Context, name string, req *WebappBrandingUpdateRequest) (*WebappBranding, error) {
	endpoint := "configuration/v1/webapp_branding/" + name + "/"

	var result WebappBranding
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteWebappBranding deletes a webapp branding
func (s *Service) DeleteWebappBranding(ctx context.Context, name string) error {
	endpoint := "configuration/v1/webapp_branding/" + name + "/"
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
