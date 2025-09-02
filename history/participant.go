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
	"strconv"
)

// ListParticipants retrieves a list of participant history records
func (s *Service) ListParticipants(ctx context.Context, opts *ListOptions) (*ParticipantListResponse, error) {
	endpoint := "history/v1/participant/"

	var result ParticipantListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetParticipant retrieves a specific participant history record by ID
func (s *Service) GetParticipant(ctx context.Context, id int) (*Participant, error) {
	endpoint := fmt.Sprintf("history/v1/participant/%d/", id)

	var result Participant
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// ListParticipantsByConference retrieves participant history for a specific conference
func (s *Service) ListParticipantsByConference(ctx context.Context, conferenceID int, opts *ListOptions) (*ParticipantListResponse, error) {
	endpoint := "history/v1/participant/"

	params := url.Values{}
	params.Set("conference_id", strconv.Itoa(conferenceID))

	if opts != nil {
		optParams := opts.ToURLValuesWithSearchField("display_name__icontains")
		for key, values := range optParams {
			for _, value := range values {
				params.Set(key, value)
			}
		}
	}

	var result ParticipantListResponse
	err := s.client.GetJSON(ctx, endpoint, &params, &result)
	return &result, err
}
