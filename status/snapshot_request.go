package status

import (
	"context"
	"fmt"
)

// ListSnapshotRequests retrieves a list of snapshot request statuses
func (s *Service) ListSnapshotRequests(ctx context.Context, opts *ListOptions) (*SnapshotRequestListResponse, error) {
	endpoint := "status/v1/snapshot_request/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result SnapshotRequestListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetSnapshotRequest retrieves a specific snapshot request status by ID
func (s *Service) GetSnapshotRequest(ctx context.Context, id int) (*SnapshotRequest, error) {
	endpoint := fmt.Sprintf("status/v1/snapshot_request/%d/", id)

	var result SnapshotRequest
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
