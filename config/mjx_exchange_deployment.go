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

// ListMjxExchangeDeployments retrieves a list of MJX Exchange deployments
func (s *Service) ListMjxExchangeDeployments(ctx context.Context, opts *ListOptions) (*MjxExchangeDeploymentListResponse, error) {
	endpoint := "configuration/v1/mjx_exchange_deployment/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MjxExchangeDeploymentListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMjxExchangeDeployment retrieves a specific MJX Exchange deployment by ID
func (s *Service) GetMjxExchangeDeployment(ctx context.Context, id int) (*MjxExchangeDeployment, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_exchange_deployment/%d/", id)

	var result MjxExchangeDeployment
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateMjxExchangeDeployment creates a new MJX Exchange deployment
func (s *Service) CreateMjxExchangeDeployment(ctx context.Context, req *MjxExchangeDeploymentCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/mjx_exchange_deployment/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMjxExchangeDeployment updates an existing MJX Exchange deployment
func (s *Service) UpdateMjxExchangeDeployment(ctx context.Context, id int, req *MjxExchangeDeploymentUpdateRequest) (*MjxExchangeDeployment, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_exchange_deployment/%d/", id)

	var result MjxExchangeDeployment
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMjxExchangeDeployment deletes a MJX Exchange deployment
func (s *Service) DeleteMjxExchangeDeployment(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/mjx_exchange_deployment/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
