/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"context"
	"fmt"
)

// ListMJXMeetings retrieves a list of MJX meeting statuses
func (s *Service) ListMJXMeetings(ctx context.Context, opts *ListOptions) (*MJXMeetingListResponse, error) {
	endpoint := "status/v1/mjx_meeting/"

	var result MJXMeetingListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetMJXMeeting retrieves a specific MJX meeting status by ID
func (s *Service) GetMJXMeeting(ctx context.Context, id string) (*MJXMeeting, error) {
	endpoint := fmt.Sprintf("status/v1/mjx_meeting/%s/", id)

	var result MJXMeeting
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
