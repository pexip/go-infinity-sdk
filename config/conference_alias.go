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

// ListConferenceAliases retrieves a list of conference aliases
func (s *Service) ListConferenceAliases(ctx context.Context, opts *ListOptions) (*ConferenceAliasListResponse, error) {
	endpoint := "configuration/v1/conference_alias/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ConferenceAliasListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetConferenceAlias retrieves a specific conference alias by ID
func (s *Service) GetConferenceAlias(ctx context.Context, id int) (*ConferenceAlias, error) {
	endpoint := fmt.Sprintf("configuration/v1/conference_alias/%d/", id)

	var result ConferenceAlias
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateConferenceAlias creates a new conference alias
func (s *Service) CreateConferenceAlias(ctx context.Context, req *ConferenceAliasCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/conference_alias/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateConferenceAlias updates an existing conference alias
func (s *Service) UpdateConferenceAlias(ctx context.Context, id int, req *ConferenceAliasUpdateRequest) (*ConferenceAlias, error) {
	endpoint := fmt.Sprintf("configuration/v1/conference_alias/%d/", id)

	var result ConferenceAlias
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteConferenceAlias deletes a conference alias
func (s *Service) DeleteConferenceAlias(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/conference_alias/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
