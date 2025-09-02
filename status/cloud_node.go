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

// ListCloudNodes retrieves a list of cloud node statuses
func (s *Service) ListCloudNodes(ctx context.Context, opts *ListOptions) (*CloudNodeListResponse, error) {
	endpoint := "status/v1/cloud_node/"

	var result CloudNodeListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetCloudNode retrieves a specific cloud node status by ID
func (s *Service) GetCloudNode(ctx context.Context, id string) (*CloudNode, error) {
	endpoint := fmt.Sprintf("status/v1/cloud_node/%s/", id)

	var result CloudNode
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
