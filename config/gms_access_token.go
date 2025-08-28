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

// ListGMSAccessTokens retrieves a list of Google Meet access tokens
func (s *Service) ListGMSAccessTokens(ctx context.Context, opts *ListOptions) (*GMSAccessTokenListResponse, error) {
	endpoint := "configuration/v1/gms_access_token/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result GMSAccessTokenListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetGMSAccessToken retrieves a specific Google Meet access token by ID
func (s *Service) GetGMSAccessToken(ctx context.Context, id int) (*GMSAccessToken, error) {
	endpoint := fmt.Sprintf("configuration/v1/gms_access_token/%d/", id)

	var result GMSAccessToken
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateGMSAccessToken creates a new Google Meet access token
func (s *Service) CreateGMSAccessToken(ctx context.Context, req *GMSAccessTokenCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/gms_access_token/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateGMSAccessToken updates an existing Google Meet access token
func (s *Service) UpdateGMSAccessToken(ctx context.Context, id int, req *GMSAccessTokenUpdateRequest) (*GMSAccessToken, error) {
	endpoint := fmt.Sprintf("configuration/v1/gms_access_token/%d/", id)

	var result GMSAccessToken
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteGMSAccessToken deletes a Google Meet access token
func (s *Service) DeleteGMSAccessToken(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/gms_access_token/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
