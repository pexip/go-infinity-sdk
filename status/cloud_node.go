package status

import (
	"context"
	"fmt"
)

// ListCloudNodes retrieves a list of cloud node statuses
func (s *Service) ListCloudNodes(ctx context.Context, opts *ListOptions) (*CloudNodeListResponse, error) {
	endpoint := "status/v1/cloud_node/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result CloudNodeListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetCloudNode retrieves a specific cloud node status by ID
func (s *Service) GetCloudNode(ctx context.Context, id string) (*CloudNode, error) {
	endpoint := fmt.Sprintf("status/v1/cloud_node/%s/", id)

	var result CloudNode
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
