package status

import (
	"context"
	"fmt"
)

// ListConferences retrieves a list of conference statuses
func (s *Service) ListConferences(ctx context.Context, opts *ListOptions) (*ConferenceListResponse, error) {
	endpoint := "status/v1/conference/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ConferenceListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetConference retrieves a specific conference status by ID
func (s *Service) GetConference(ctx context.Context, id string) (*ConferenceStatus, error) {
	endpoint := fmt.Sprintf("status/v1/conference/%s/", id)

	var result ConferenceStatus
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}