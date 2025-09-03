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

// ListConferences retrieves a list of conference statuses
func (s *Service) ListConferences(ctx context.Context, opts *ListOptions) (*ConferenceListResponse, error) {
	endpoint := "status/v1/conference/"

	var result ConferenceListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetConference retrieves a specific conference status by ID
func (s *Service) GetConference(ctx context.Context, id string) (*ConferenceStatus, error) {
	endpoint := fmt.Sprintf("status/v1/conference/%s/", id)

	var result ConferenceStatus
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
