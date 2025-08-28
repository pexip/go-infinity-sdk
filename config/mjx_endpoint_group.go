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

// ListMjxEndpointGroups retrieves a list of MJX endpoint groups
func (s *Service) ListMjxEndpointGroups(ctx context.Context, opts *ListOptions) (*MjxEndpointGroupListResponse, error) {
	endpoint := "configuration/v1/mjx_endpoint_group/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MjxEndpointGroupListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMjxEndpointGroup retrieves a specific MJX endpoint group by ID
func (s *Service) GetMjxEndpointGroup(ctx context.Context, id int) (*MjxEndpointGroup, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_endpoint_group/%d/", id)

	var result MjxEndpointGroup
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateMjxEndpointGroup creates a new MJX endpoint group
func (s *Service) CreateMjxEndpointGroup(ctx context.Context, req *MjxEndpointGroupCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/mjx_endpoint_group/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMjxEndpointGroup updates an existing MJX endpoint group
func (s *Service) UpdateMjxEndpointGroup(ctx context.Context, id int, req *MjxEndpointGroupUpdateRequest) (*MjxEndpointGroup, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_endpoint_group/%d/", id)

	var result MjxEndpointGroup
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMjxEndpointGroup deletes a MJX endpoint group
func (s *Service) DeleteMjxEndpointGroup(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/mjx_endpoint_group/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
