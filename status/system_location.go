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

// ListSystemLocations retrieves a list of system location statuses
func (s *Service) ListSystemLocations(ctx context.Context, opts *ListOptions) (*SystemLocationListResponse, error) {
	endpoint := "status/v1/system_location/"

	var result SystemLocationListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetSystemLocation retrieves a specific system location status by ID
func (s *Service) GetSystemLocation(ctx context.Context, id int) (*SystemLocation, error) {
	endpoint := fmt.Sprintf("status/v1/system_location/%d/", id)

	var result SystemLocation
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
