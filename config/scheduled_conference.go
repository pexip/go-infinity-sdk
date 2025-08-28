/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListScheduledConferences retrieves a list of scheduled conferences
func (s *Service) ListScheduledConferences(ctx context.Context, opts *ListOptions) (*ScheduledConferenceListResponse, error) {
	endpoint := "configuration/v1/scheduled_conference/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ScheduledConferenceListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetScheduledConference retrieves a specific scheduled conference by ID
func (s *Service) GetScheduledConference(ctx context.Context, id int) (*ScheduledConference, error) {
	endpoint := fmt.Sprintf("configuration/v1/scheduled_conference/%d/", id)

	var result ScheduledConference
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateScheduledConference creates a new scheduled conference
func (s *Service) CreateScheduledConference(ctx context.Context, req *ScheduledConferenceCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/scheduled_conference/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateScheduledConference updates an existing scheduled conference
func (s *Service) UpdateScheduledConference(ctx context.Context, id int, req *ScheduledConferenceUpdateRequest) (*ScheduledConference, error) {
	endpoint := fmt.Sprintf("configuration/v1/scheduled_conference/%d/", id)

	var result ScheduledConference
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteScheduledConference deletes a scheduled conference
func (s *Service) DeleteScheduledConference(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/scheduled_conference/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
