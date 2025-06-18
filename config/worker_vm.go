package config

import (
	"context"
	"fmt"
)

// ListWorkerVMs retrieves a list of worker VMs
func (s *Service) ListWorkerVMs(ctx context.Context, opts *ListOptions) (*WorkerVMListResponse, error) {
	endpoint := "configuration/v1/worker_vm/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result WorkerVMListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetWorkerVM retrieves a specific worker VM by ID
func (s *Service) GetWorkerVM(ctx context.Context, id int) (*WorkerVM, error) {
	endpoint := fmt.Sprintf("configuration/v1/worker_vm/%d/", id)

	var result WorkerVM
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateWorkerVM creates a new worker VM
func (s *Service) CreateWorkerVM(ctx context.Context, req *WorkerVMCreateRequest) (*ResourceCreateResponse, error) {
	endpoint := "configuration/v1/worker_vm/"

	resp, err := s.client.PostJSONRawBodyResponse(ctx, endpoint, req)
	return resp, err
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
