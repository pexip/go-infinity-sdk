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

// ListAutomaticParticipants retrieves a list of automatic participants
func (s *Service) ListAutomaticParticipants(ctx context.Context, opts *ListOptions) (*AutomaticParticipantListResponse, error) {
	endpoint := "configuration/v1/automatic_participant/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result AutomaticParticipantListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetAutomaticParticipant retrieves a specific automatic participant by ID
func (s *Service) GetAutomaticParticipant(ctx context.Context, id int) (*AutomaticParticipant, error) {
	endpoint := fmt.Sprintf("configuration/v1/automatic_participant/%d/", id)

	var result AutomaticParticipant
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateAutomaticParticipant creates a new automatic participant
func (s *Service) CreateAutomaticParticipant(ctx context.Context, req *AutomaticParticipantCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/automatic_participant/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateAutomaticParticipant updates an existing automatic participant
func (s *Service) UpdateAutomaticParticipant(ctx context.Context, id int, req *AutomaticParticipantUpdateRequest) (*AutomaticParticipant, error) {
	endpoint := fmt.Sprintf("configuration/v1/automatic_participant/%d/", id)

	var result AutomaticParticipant
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteAutomaticParticipant deletes an automatic participant
func (s *Service) DeleteAutomaticParticipant(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/automatic_participant/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
