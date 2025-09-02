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

// ListTeamsNodeCalls retrieves a list of Teams node call statuses
func (s *Service) ListTeamsNodeCalls(ctx context.Context, opts *ListOptions) (*TeamsNodeCallListResponse, error) {
	endpoint := "status/v1/teamsnode_call/"

	var result TeamsNodeCallListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetTeamsNodeCall retrieves a specific Teams node call status by ID
func (s *Service) GetTeamsNodeCall(ctx context.Context, id string) (*TeamsNodeCall, error) {
	endpoint := fmt.Sprintf("status/v1/teamsnode_call/%s/", id)

	var result TeamsNodeCall
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
