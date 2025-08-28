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

// ListTelehealthProfiles retrieves a list of telehealth profiles
func (s *Service) ListTelehealthProfiles(ctx context.Context, opts *ListOptions) (*TelehealthProfileListResponse, error) {
	endpoint := "configuration/v1/telehealth_profile/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result TelehealthProfileListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetTelehealthProfile retrieves a specific telehealth profile by ID
func (s *Service) GetTelehealthProfile(ctx context.Context, id int) (*TelehealthProfile, error) {
	endpoint := fmt.Sprintf("configuration/v1/telehealth_profile/%d/", id)

	var result TelehealthProfile
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateTelehealthProfile creates a new telehealth profile
func (s *Service) CreateTelehealthProfile(ctx context.Context, req *TelehealthProfileCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/telehealth_profile/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateTelehealthProfile updates an existing telehealth profile
func (s *Service) UpdateTelehealthProfile(ctx context.Context, id int, req *TelehealthProfileUpdateRequest) (*TelehealthProfile, error) {
	endpoint := fmt.Sprintf("configuration/v1/telehealth_profile/%d/", id)

	var result TelehealthProfile
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteTelehealthProfile deletes a telehealth profile
func (s *Service) DeleteTelehealthProfile(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/telehealth_profile/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
