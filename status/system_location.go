package status

import (
	"context"
	"fmt"
)

// ListSystemLocations retrieves a list of system location statuses
func (s *Service) ListSystemLocations(ctx context.Context, opts *ListOptions) (*SystemLocationListResponse, error) {
	endpoint := "status/v1/system_location/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result SystemLocationListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetSystemLocation retrieves a specific system location status by ID
func (s *Service) GetSystemLocation(ctx context.Context, id int) (*SystemLocation, error) {
	endpoint := fmt.Sprintf("status/v1/system_location/%d/", id)

	var result SystemLocation
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
