/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package history

import (
	"context"
	"fmt"
)

// ListWorkerVMStatusEvents retrieves a list of worker VM status event history records
func (s *Service) ListWorkerVMStatusEvents(ctx context.Context, opts *ListOptions) (*WorkerVMStatusEventListResponse, error) {
	endpoint := "history/v1/workervm_status_event/"

	var result WorkerVMStatusEventListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetWorkerVMStatusEvent retrieves a specific worker VM status event history record by ID
func (s *Service) GetWorkerVMStatusEvent(ctx context.Context, id int) (*WorkerVMStatusEvent, error) {
	endpoint := fmt.Sprintf("history/v1/workervm_status_event/%d/", id)

	var result WorkerVMStatusEvent
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
