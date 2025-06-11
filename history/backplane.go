package history

import (
	"context"
	"fmt"
)

// ListBackplanes retrieves a list of backplane history records
func (s *Service) ListBackplanes(ctx context.Context, opts *ListOptions) (*BackplaneListResponse, error) {
	endpoint := "history/v1/backplane/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result BackplaneListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetBackplane retrieves a specific backplane history record by ID
func (s *Service) GetBackplane(ctx context.Context, id string) (*Backplane, error) {
	endpoint := fmt.Sprintf("history/v1/backplane/%s/", id)

	var result Backplane
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
