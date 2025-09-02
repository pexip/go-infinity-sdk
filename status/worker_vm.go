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

// ListWorkerVMs retrieves a list of worker VM statuses
func (s *Service) ListWorkerVMs(ctx context.Context, opts *ListOptions) (*WorkerVMListResponse, error) {
	endpoint := "status/v1/worker_vm/"

	var result WorkerVMListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetWorkerVM retrieves a specific worker VM status by ID
func (s *Service) GetWorkerVM(ctx context.Context, id int) (*WorkerVM, error) {
	endpoint := fmt.Sprintf("status/v1/worker_vm/%d/", id)

	var result WorkerVM
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
