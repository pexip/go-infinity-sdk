/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListWorkerVMs retrieves a list of worker VMs
func (s *Service) ListWorkerVMs(ctx context.Context, opts *ListOptions) (*WorkerVMListResponse, error) {
	endpoint := "configuration/v1/worker_vm/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result WorkerVMListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetWorkerVM retrieves a specific worker VM by ID
func (s *Service) GetWorkerVM(ctx context.Context, id int) (*WorkerVM, error) {
	endpoint := fmt.Sprintf("configuration/v1/worker_vm/%d/", id)

	var result WorkerVM
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateWorkerVM creates a new worker VM
func (s *Service) CreateWorkerVM(ctx context.Context, req *WorkerVMCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/worker_vm/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateWorkerVM updates an existing worker VM
func (s *Service) UpdateWorkerVM(ctx context.Context, id int, req *WorkerVMUpdateRequest) (*WorkerVM, error) {
	endpoint := fmt.Sprintf("configuration/v1/worker_vm/%d/", id)

	var result WorkerVM
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteWorkerVM deletes a worker VM
func (s *Service) DeleteWorkerVM(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/worker_vm/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
