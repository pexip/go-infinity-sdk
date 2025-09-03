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

// ListRoles retrieves a list of roles
func (s *Service) ListRoles(ctx context.Context, opts *ListOptions) (*RoleListResponse, error) {
	endpoint := "configuration/v1/role/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result RoleListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetRole retrieves a specific role by ID
func (s *Service) GetRole(ctx context.Context, id int) (*Role, error) {
	endpoint := fmt.Sprintf("configuration/v1/role/%d/", id)

	var result Role
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateRole creates a new role
func (s *Service) CreateRole(ctx context.Context, req *RoleCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/role/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateRole updates an existing role
func (s *Service) UpdateRole(ctx context.Context, id int, req *RoleUpdateRequest) (*Role, error) {
	endpoint := fmt.Sprintf("configuration/v1/role/%d/", id)

	var result Role
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteRole deletes a role
func (s *Service) DeleteRole(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/role/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
