package status

import (
	"context"
	"fmt"
)

// ListMJXMeetings retrieves a list of MJX meeting statuses
func (s *Service) ListMJXMeetings(ctx context.Context, opts *ListOptions) (*MJXMeetingListResponse, error) {
	endpoint := "status/v1/mjx_meeting/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MJXMeetingListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMJXMeeting retrieves a specific MJX meeting status by ID
func (s *Service) GetMJXMeeting(ctx context.Context, id string) (*MJXMeeting, error) {
	endpoint := fmt.Sprintf("status/v1/mjx_meeting/%s/", id)

	var result MJXMeeting
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}