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

// ListMjxGoogleDeployments retrieves a list of MJX Google deployments
func (s *Service) ListMjxGoogleDeployments(ctx context.Context, opts *ListOptions) (*MjxGoogleDeploymentListResponse, error) {
	endpoint := "configuration/v1/mjx_google_deployment/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result MjxGoogleDeploymentListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetMjxGoogleDeployment retrieves a specific MJX Google deployment by ID
func (s *Service) GetMjxGoogleDeployment(ctx context.Context, id int) (*MjxGoogleDeployment, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_google_deployment/%d/", id)

	var result MjxGoogleDeployment
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateMjxGoogleDeployment creates a new MJX Google deployment
func (s *Service) CreateMjxGoogleDeployment(ctx context.Context, req *MjxGoogleDeploymentCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/mjx_google_deployment/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMjxGoogleDeployment updates an existing MJX Google deployment
func (s *Service) UpdateMjxGoogleDeployment(ctx context.Context, id int, req *MjxGoogleDeploymentUpdateRequest) (*MjxGoogleDeployment, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_google_deployment/%d/", id)

	var result MjxGoogleDeployment
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMjxGoogleDeployment deletes a MJX Google deployment
func (s *Service) DeleteMjxGoogleDeployment(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/mjx_google_deployment/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
