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

// ListTeamsNodes retrieves a list of Teams node statuses
func (s *Service) ListTeamsNodes(ctx context.Context, opts *ListOptions) (*TeamsNodeListResponse, error) {
	endpoint := "status/v1/teamsnode/"

	var result TeamsNodeListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetTeamsNode retrieves a specific Teams node status by ID
func (s *Service) GetTeamsNode(ctx context.Context, id int) (*TeamsNode, error) {
	endpoint := fmt.Sprintf("status/v1/teamsnode/%d/", id)

	var result TeamsNode
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
