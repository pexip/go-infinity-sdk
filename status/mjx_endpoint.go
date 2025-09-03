/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"context"
	"fmt"
)

// ListMJXEndpoints retrieves a list of MJX endpoint statuses
func (s *Service) ListMJXEndpoints(ctx context.Context, opts *ListOptions) (*MJXEndpointListResponse, error) {
	endpoint := "status/v1/mjx_endpoint/"

	var result MJXEndpointListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetMJXEndpoint retrieves a specific MJX endpoint status by ID
func (s *Service) GetMJXEndpoint(ctx context.Context, id int) (*MJXEndpoint, error) {
	endpoint := fmt.Sprintf("status/v1/mjx_endpoint/%d/", id)

	var result MJXEndpoint
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
