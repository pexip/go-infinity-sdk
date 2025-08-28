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

// ListMjxIntegrations retrieves a list of MJX integrations
func (s *Service) ListMjxIntegrations(ctx context.Context, opts *ListOptions) (*MjxIntegrationListResponse, error) {
	endpoint := "configuration/v1/mjx_integration/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MjxIntegrationListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMjxIntegration retrieves a specific MJX integration by ID
func (s *Service) GetMjxIntegration(ctx context.Context, id int) (*MjxIntegration, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_integration/%d/", id)

	var result MjxIntegration
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateMjxIntegration creates a new MJX integration
func (s *Service) CreateMjxIntegration(ctx context.Context, req *MjxIntegrationCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/mjx_integration/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMjxIntegration updates an existing MJX integration
func (s *Service) UpdateMjxIntegration(ctx context.Context, id int, req *MjxIntegrationUpdateRequest) (*MjxIntegration, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_integration/%d/", id)

	var result MjxIntegration
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMjxIntegration deletes a MJX integration
func (s *Service) DeleteMjxIntegration(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/mjx_integration/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
