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

// ListBreakInAllowListAddresses retrieves a list of break-in attempt IP allow list entries
func (s *Service) ListBreakInAllowListAddresses(ctx context.Context, opts *ListOptions) (*BreakInAllowListAddressListResponse, error) {
	endpoint := "configuration/v1/break_in_allow_list_address/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result BreakInAllowListAddressListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetBreakInAllowListAddress retrieves a specific break-in attempt IP allow list entry by ID
func (s *Service) GetBreakInAllowListAddress(ctx context.Context, id int) (*BreakInAllowListAddress, error) {
	endpoint := fmt.Sprintf("configuration/v1/break_in_allow_list_address/%d/", id)

	var result BreakInAllowListAddress
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateBreakInAllowListAddress creates a new break-in attempt IP allow list entry
func (s *Service) CreateBreakInAllowListAddress(ctx context.Context, req *BreakInAllowListAddressCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/break_in_allow_list_address/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateBreakInAllowListAddress updates an existing break-in attempt IP allow list entry
func (s *Service) UpdateBreakInAllowListAddress(ctx context.Context, id int, req *BreakInAllowListAddressUpdateRequest) (*BreakInAllowListAddress, error) {
	endpoint := fmt.Sprintf("configuration/v1/break_in_allow_list_address/%d/", id)

	var result BreakInAllowListAddress
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteBreakInAllowListAddress deletes a break-in attempt IP allow list entry
func (s *Service) DeleteBreakInAllowListAddress(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/break_in_allow_list_address/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
