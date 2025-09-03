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

// ListBackupRequests retrieves a list of backup request statuses
func (s *Service) ListBackupRequests(ctx context.Context, opts *ListOptions) (*BackupRequestListResponse, error) {
	endpoint := "status/v1/backup_request/"

	var result BackupRequestListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetBackupRequest retrieves a specific backup request status by ID
func (s *Service) GetBackupRequest(ctx context.Context, id int) (*BackupRequest, error) {
	endpoint := fmt.Sprintf("status/v1/backup_request/%d/", id)

	var result BackupRequest
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
