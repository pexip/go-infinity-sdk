package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListRoleMappings retrieves a list of role mappings
func (s *Service) ListRoleMappings(ctx context.Context, opts *ListOptions) (*RoleMappingListResponse, error) {
	endpoint := "configuration/v1/role_mapping/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result RoleMappingListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetRoleMapping retrieves a specific role mapping by ID
func (s *Service) GetRoleMapping(ctx context.Context, id int) (*RoleMapping, error) {
	endpoint := fmt.Sprintf("configuration/v1/role_mapping/%d/", id)

	var result RoleMapping
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateRoleMapping creates a new role mapping
func (s *Service) CreateRoleMapping(ctx context.Context, req *RoleMappingCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/role_mapping/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateRoleMapping updates an existing role mapping
func (s *Service) UpdateRoleMapping(ctx context.Context, id int, req *RoleMappingUpdateRequest) (*RoleMapping, error) {
	endpoint := fmt.Sprintf("configuration/v1/role_mapping/%d/", id)

	var result RoleMapping
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteRoleMapping deletes a role mapping
func (s *Service) DeleteRoleMapping(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/role_mapping/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
