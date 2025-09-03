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

// ListSnapshotRequests retrieves a list of snapshot request statuses
func (s *Service) ListSnapshotRequests(ctx context.Context, opts *ListOptions) (*SnapshotRequestListResponse, error) {
	endpoint := "status/v1/snapshot_request/"

	var result SnapshotRequestListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetSnapshotRequest retrieves a specific snapshot request status by ID
func (s *Service) GetSnapshotRequest(ctx context.Context, id int) (*SnapshotRequest, error) {
	endpoint := fmt.Sprintf("status/v1/snapshot_request/%d/", id)

	var result SnapshotRequest
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
