package history

import (
	"context"
	"fmt"
)

// ListConferences retrieves a list of conference history records
func (s *Service) ListConferences(ctx context.Context, opts *ListOptions) (*ConferenceListResponse, error) {
	endpoint := "history/v1/conference/"

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

// GetConference retrieves a specific conference history record by ID
func (s *Service) GetConference(ctx context.Context, id int) (*Conference, error) {
	endpoint := fmt.Sprintf("history/v1/conference/%d/", id)

	var result Conference
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
