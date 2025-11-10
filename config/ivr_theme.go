/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListIVRThemes retrieves a list of IVR themes
func (s *Service) ListIVRThemes(ctx context.Context, opts *ListOptions) (*IVRThemeListResponse, error) {
	endpoint := "configuration/v1/ivr_theme/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result IVRThemeListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetIVRTheme retrieves a specific IVR theme by ID
func (s *Service) GetIVRTheme(ctx context.Context, id int) (*IVRTheme, error) {
	endpoint := fmt.Sprintf("configuration/v1/ivr_theme/%d/", id)

	var result IVRTheme
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateIVRTheme creates a new IVR theme
func (s *Service) CreateIVRTheme(ctx context.Context, req *IVRThemeCreateRequest, filename string, file io.Reader) (*types.PostResponse, error) {
	endpoint := "configuration/v1/ivr_theme/"
	resp, err := s.client.PostWithResponse(ctx, endpoint, req, nil)
	if err != nil {
		return nil, err
	}

	// Extract ID from resource URI (e.g., "/api/admin/configuration/v1/ivr_theme/123/" -> 123)
	id, err := resp.ResourceID()
	if err != nil {
		return resp, fmt.Errorf("failed to parse ID from resource URI: %w", err)
	}

	// Upload the package file
	if err = s.client.PatchFile(ctx, fmt.Sprintf("%s/%d/", endpoint, id), "package", filename, file, nil); err != nil {
		return resp, fmt.Errorf("failed to upload package file: %w", err)
	}
	return resp, nil
}

// UpdateIVRTheme updates an existing IVR theme
func (s *Service) UpdateIVRTheme(ctx context.Context, id int, req *IVRThemeUpdateRequest, filename string, file io.Reader) (*IVRTheme, error) {
	endpoint := fmt.Sprintf("configuration/v1/ivr_theme/%d/", id)

	var result IVRTheme
	err := s.client.PatchJSON(ctx, endpoint, req, &result)
	if err != nil {
		return nil, err
	}

	// Upload the package file
	if err = s.client.PatchFile(ctx, endpoint, "package", filename, file, nil); err != nil {
		return &result, fmt.Errorf("failed to upload package file: %w", err)
	}
	return &result, err
}

// DeleteIVRTheme deletes an IVR theme
func (s *Service) DeleteIVRTheme(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/ivr_theme/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
