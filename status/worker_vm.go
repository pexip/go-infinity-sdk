package status

import (
	"context"
	"fmt"
)

// ListWorkerVMs retrieves a list of worker VM statuses
func (s *Service) ListWorkerVMs(ctx context.Context, opts *ListOptions) (*WorkerVMListResponse, error) {
	endpoint := "status/v1/worker_vm/"

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

// GetWorkerVM retrieves a specific worker VM status by ID
func (s *Service) GetWorkerVM(ctx context.Context, id int) (*WorkerVM, error) {
	endpoint := fmt.Sprintf("status/v1/worker_vm/%d/", id)

	var result WorkerVM
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
