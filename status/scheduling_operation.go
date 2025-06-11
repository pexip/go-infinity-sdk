package status

import (
	"context"
	"fmt"
)

// ListSchedulingOperations retrieves a list of scheduling operation statuses
func (s *Service) ListSchedulingOperations(ctx context.Context, opts *ListOptions) (*SchedulingOperationListResponse, error) {
	endpoint := "status/v1/scheduling_operation/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result SchedulingOperationListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetSchedulingOperation retrieves a specific scheduling operation status by ID
func (s *Service) GetSchedulingOperation(ctx context.Context, id int) (*SchedulingOperation, error) {
	endpoint := fmt.Sprintf("status/v1/scheduling_operation/%d/", id)

	var result SchedulingOperation
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}