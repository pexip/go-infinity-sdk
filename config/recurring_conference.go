/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListRecurringConferences retrieves a list of recurring conferences
func (s *Service) ListRecurringConferences(ctx context.Context, opts *ListOptions) (*RecurringConferenceListResponse, error) {
	endpoint := "configuration/v1/recurring_conference/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result RecurringConferenceListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetRecurringConference retrieves a specific recurring conference by ID
func (s *Service) GetRecurringConference(ctx context.Context, id int) (*RecurringConference, error) {
	endpoint := fmt.Sprintf("configuration/v1/recurring_conference/%d/", id)

	var result RecurringConference
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateRecurringConference creates a new recurring conference
func (s *Service) CreateRecurringConference(ctx context.Context, req *RecurringConferenceCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/recurring_conference/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateRecurringConference updates an existing recurring conference
func (s *Service) UpdateRecurringConference(ctx context.Context, id int, req *RecurringConferenceUpdateRequest) (*RecurringConference, error) {
	endpoint := fmt.Sprintf("configuration/v1/recurring_conference/%d/", id)

	var result RecurringConference
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteRecurringConference deletes a recurring conference
func (s *Service) DeleteRecurringConference(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/recurring_conference/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
