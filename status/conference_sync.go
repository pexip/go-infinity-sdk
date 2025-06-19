package status

import (
	"context"
	"fmt"
)

// ListConferenceSyncs retrieves a list of conference sync statuses
func (s *Service) ListConferenceSyncs(ctx context.Context, opts *ListOptions) (*ConferenceSyncListResponse, error) {
	endpoint := "status/v1/conference_sync/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ConferenceSyncListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetConferenceSync retrieves a specific conference sync status by ID
func (s *Service) GetConferenceSync(ctx context.Context, id int) (*ConferenceSync, error) {
	endpoint := fmt.Sprintf("status/v1/conference_sync/%d/", id)

	var result ConferenceSync
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
