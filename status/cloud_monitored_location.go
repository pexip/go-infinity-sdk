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

// ListCloudMonitoredLocations retrieves a list of cloud monitored location statuses
func (s *Service) ListCloudMonitoredLocations(ctx context.Context, opts *ListOptions) (*CloudMonitoredLocationListResponse, error) {
	endpoint := "status/v1/cloud_monitored_location/"

	var result CloudMonitoredLocationListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetCloudMonitoredLocation retrieves a specific cloud monitored location status by ID
func (s *Service) GetCloudMonitoredLocation(ctx context.Context, id int) (*CloudMonitoredLocation, error) {
	endpoint := fmt.Sprintf("status/v1/cloud_monitored_location/%d/", id)

	var result CloudMonitoredLocation
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
