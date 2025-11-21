/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"encoding/json"
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

	// Create form fields from request
	fields := map[string]string{
		"name":            req.Name,
		"uuid":            req.UUID,
		"custom_layouts":  req.CustomLayouts,
		"pinning_configs": req.PinningConfigs,
	}

	// Convert Conference slice to JSON string if not empty
	if len(req.Conference) > 0 {
		conferenceJSON, err := json.Marshal(req.Conference)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal conference field: %w", err)
		}
		fields["conference"] = string(conferenceJSON)
	}

	return s.client.PostMultipartFormWithFieldsAndResponse(ctx, endpoint, fields, "package", filename, file, nil)
}

// UpdateIVRTheme updates an existing IVR theme
func (s *Service) UpdateIVRTheme(ctx context.Context, id int, req *IVRThemeUpdateRequest, filename string, file io.Reader) (*IVRTheme, error) {
	endpoint := fmt.Sprintf("configuration/v1/ivr_theme/%d/", id)

	// Create form fields from request
	fields := map[string]string{
		"name":            req.Name,
		"uuid":            req.UUID,
		"custom_layouts":  req.CustomLayouts,
		"pinning_configs": req.PinningConfigs,
	}

	// Convert Conference slice to JSON string if not empty
	if len(req.Conference) > 0 {
		conferenceJSON, err := json.Marshal(req.Conference)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal conference field: %w", err)
		}
		fields["conference"] = string(conferenceJSON)
	}

	var result IVRTheme
	_, err := s.client.PatchMultipartFormWithFieldsAndResponse(ctx, endpoint, fields, "package", filename, file, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteIVRTheme deletes an IVR theme
func (s *Service) DeleteIVRTheme(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/ivr_theme/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
