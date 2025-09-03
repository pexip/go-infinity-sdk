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

// ListScheduledAliases retrieves a list of scheduled aliases
func (s *Service) ListScheduledAliases(ctx context.Context, opts *ListOptions) (*ScheduledAliasListResponse, error) {
	endpoint := "configuration/v1/scheduled_alias/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result ScheduledAliasListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetScheduledAlias retrieves a specific scheduled alias by ID
func (s *Service) GetScheduledAlias(ctx context.Context, id int) (*ScheduledAlias, error) {
	endpoint := fmt.Sprintf("configuration/v1/scheduled_alias/%d/", id)

	var result ScheduledAlias
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateScheduledAlias creates a new scheduled alias
func (s *Service) CreateScheduledAlias(ctx context.Context, req *ScheduledAliasCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/scheduled_alias/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateScheduledAlias updates an existing scheduled alias
func (s *Service) UpdateScheduledAlias(ctx context.Context, id int, req *ScheduledAliasUpdateRequest) (*ScheduledAlias, error) {
	endpoint := fmt.Sprintf("configuration/v1/scheduled_alias/%d/", id)

	var result ScheduledAlias
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteScheduledAlias deletes a scheduled alias
func (s *Service) DeleteScheduledAlias(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/scheduled_alias/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
