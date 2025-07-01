package config

import (
	"context"
	"fmt"
)

// ListPermissions retrieves a list of permissions (read-only)
func (s *Service) ListPermissions(ctx context.Context, opts *ListOptions) (*PermissionListResponse, error) {
	endpoint := "configuration/v1/permission/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result PermissionListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetPermission retrieves a specific permission by ID (read-only)
func (s *Service) GetPermission(ctx context.Context, id int) (*Permission, error) {
	endpoint := fmt.Sprintf("configuration/v1/permission/%d/", id)

	var result Permission
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
