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

// ListSchedulingOperations retrieves a list of scheduling operation statuses
func (s *Service) ListSchedulingOperations(ctx context.Context, opts *ListOptions) (*SchedulingOperationListResponse, error) {
	endpoint := "status/v1/scheduling_operation/"

	var result SchedulingOperationListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetSchedulingOperation retrieves a specific scheduling operation status by ID
func (s *Service) GetSchedulingOperation(ctx context.Context, id int) (*SchedulingOperation, error) {
	endpoint := fmt.Sprintf("status/v1/scheduling_operation/%d/", id)

	var result SchedulingOperation
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
