package status

import (
	"context"
	"fmt"
)

// ListMJXEndpoints retrieves a list of MJX endpoint statuses
func (s *Service) ListMJXEndpoints(ctx context.Context, opts *ListOptions) (*MJXEndpointListResponse, error) {
	endpoint := "status/v1/mjx_endpoint/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MJXEndpointListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMJXEndpoint retrieves a specific MJX endpoint status by ID
func (s *Service) GetMJXEndpoint(ctx context.Context, id int) (*MJXEndpoint, error) {
	endpoint := fmt.Sprintf("status/v1/mjx_endpoint/%d/", id)

	var result MJXEndpoint
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
