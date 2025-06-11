package history

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// ListMediaStreams retrieves a list of media stream history records
func (s *Service) ListMediaStreams(ctx context.Context, opts *ListOptions) (*MediaStreamListResponse, error) {
	endpoint := "history/v1/media_stream/"

	if opts != nil {
		params := opts.BaseListOptions.ToURLValues()
		if opts.StartTime != nil {
			params.Set("start_time__gte", opts.StartTime.Format(time.RFC3339))
		}
		if opts.EndTime != nil {
			params.Set("end_time__lt", opts.EndTime.Format(time.RFC3339))
		}
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MediaStreamListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMediaStream retrieves a specific media stream history record by ID
func (s *Service) GetMediaStream(ctx context.Context, id int) (*MediaStream, error) {
	endpoint := fmt.Sprintf("history/v1/media_stream/%d/", id)

	var result MediaStream
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// ListMediaStreamsByParticipant retrieves media stream history for a specific participant
func (s *Service) ListMediaStreamsByParticipant(ctx context.Context, participantID int, opts *ListOptions) (*MediaStreamListResponse, error) {
	endpoint := "history/v1/media_stream/"

	params := url.Values{}
	params.Set("participant_id", strconv.Itoa(participantID))

	if opts != nil {
		optParams := opts.BaseListOptions.ToURLValues()
		for key, values := range optParams {
			for _, value := range values {
				params.Set(key, value)
			}
		}
		if opts.StartTime != nil {
			params.Set("start_time__gte", opts.StartTime.Format(time.RFC3339))
		}
		if opts.EndTime != nil {
			params.Set("end_time__lt", opts.EndTime.Format(time.RFC3339))
		}
	}

	endpoint += "?" + params.Encode()

	var result MediaStreamListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
