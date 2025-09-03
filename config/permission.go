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
)

// ListPermissions retrieves a list of permissions (read-only)
func (s *Service) ListPermissions(ctx context.Context, opts *ListOptions) (*PermissionListResponse, error) {
	endpoint := "configuration/v1/permission/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result PermissionListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetPermission retrieves a specific permission by ID (read-only)
func (s *Service) GetPermission(ctx context.Context, id int) (*Permission, error) {
	endpoint := fmt.Sprintf("configuration/v1/permission/%d/", id)

	var result Permission
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
