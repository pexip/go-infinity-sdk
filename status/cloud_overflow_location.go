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

// ListCloudOverflowLocations retrieves a list of cloud overflow location statuses
func (s *Service) ListCloudOverflowLocations(ctx context.Context, opts *ListOptions) (*CloudOverflowLocationListResponse, error) {
	endpoint := "status/v1/cloud_overflow_location/"

	var result CloudOverflowLocationListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetCloudOverflowLocation retrieves a specific cloud overflow location status by ID
func (s *Service) GetCloudOverflowLocation(ctx context.Context, id int) (*CloudOverflowLocation, error) {
	endpoint := fmt.Sprintf("status/v1/cloud_overflow_location/%d/", id)

	var result CloudOverflowLocation
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
