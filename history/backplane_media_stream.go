/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package history

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// ListBackplaneMediaStreams retrieves a list of backplane media stream history records
func (s *Service) ListBackplaneMediaStreams(ctx context.Context, opts *ListOptions) (*BackplaneMediaStreamListResponse, error) {
	endpoint := "history/v1/backplane_media_stream/"

	var result BackplaneMediaStreamListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetBackplaneMediaStream retrieves a specific backplane media stream history record by ID
func (s *Service) GetBackplaneMediaStream(ctx context.Context, id int) (*BackplaneMediaStream, error) {
	endpoint := fmt.Sprintf("history/v1/backplane_media_stream/%d/", id)

	var result BackplaneMediaStream
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// ListBackplaneMediaStreamsByBackplane retrieves backplane media stream history for a specific backplane
func (s *Service) ListBackplaneMediaStreamsByBackplane(ctx context.Context, backplaneID string, opts *ListOptions) (*BackplaneMediaStreamListResponse, error) {
	endpoint := "history/v1/backplane_media_stream/"

	params := url.Values{}
	params.Set("backplane", backplaneID)

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

	var result BackplaneMediaStreamListResponse
	err := s.client.GetJSON(ctx, endpoint, &params, &result)
	return &result, err
}
