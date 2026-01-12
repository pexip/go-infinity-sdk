/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"io"
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
func (s *Service) GetWebappBranding(ctx context.Context, uuid string) (*WebappBranding, error) {
	endpoint := "configuration/v1/webapp_branding/" + uuid + "/"

	var result WebappBranding
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

func toBoolStr(val bool) string {
	if val {
		return "True"
	}
	return "False"
}

// CreateWebappBranding creates a new webapp branding
func (s *Service) CreateWebappBranding(ctx context.Context, req *WebappBrandingCreateRequest, filename string, file io.Reader) (*types.PostResponseWithUUID, error) {
	endpoint := "configuration/v1/webapp_branding/"

	// Create form fields from request
	fields := map[string]string{
		"name":        req.Name,
		"description": req.Description,
		"webapp_type": req.WebappType,
	}

	var err error
	var resp *types.PostResponseWithUUID
	var result WebappBranding
	if resp, err = s.client.PostMultipartFormWithFieldsAndResponseUUID(ctx, endpoint, fields, "branding_file", filename, file, &result); err != nil {
		return resp, err
	}

	return resp, err
}

// UpdateWebappBranding updates an existing webapp branding
func (s *Service) UpdateWebappBranding(ctx context.Context, req *WebappBrandingUpdateRequest, uuid string) (*WebappBranding, error) {
	endpoint := "configuration/v1/webapp_branding/" + uuid + "/"

	// Create form fields from request
	fields := map[string]string{
		"name":        req.Name,
		"description": req.Description,
		"webapp_type": req.WebappType,
	}

	var result WebappBranding
	err := s.client.PatchJSON(ctx, endpoint, fields, &result)

	return &result, err
}

// DeleteWebappBranding deletes a webapp branding
func (s *Service) DeleteWebappBranding(ctx context.Context, uuid string) error {
	endpoint := "configuration/v1/webapp_branding/" + uuid + "/"
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
