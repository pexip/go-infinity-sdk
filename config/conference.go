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

// ListConferences retrieves a list of conferences
func (s *Service) ListConferences(ctx context.Context, opts *ListOptions) (*ConferenceListResponse, error) {
	endpoint := "configuration/v1/conference/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result ConferenceListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetConference retrieves a specific conference by ID
func (s *Service) GetConference(ctx context.Context, id int) (*Conference, error) {
	endpoint := fmt.Sprintf("configuration/v1/conference/%d/", id)

	var result Conference
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateConference creates a new conference
func (s *Service) CreateConference(ctx context.Context, req *ConferenceCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/conference/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateConference updates an existing conference
func (s *Service) UpdateConference(ctx context.Context, id int, req *ConferenceUpdateRequest) (*Conference, error) {
	endpoint := fmt.Sprintf("configuration/v1/conference/%d/", id)

	var result Conference
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteConference deletes a conference
func (s *Service) DeleteConference(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/conference/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
