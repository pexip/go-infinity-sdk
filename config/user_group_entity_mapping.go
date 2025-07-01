package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListUserGroupEntityMappings retrieves a list of user group entity mappings
func (s *Service) ListUserGroupEntityMappings(ctx context.Context, opts *ListOptions) (*UserGroupEntityMappingListResponse, error) {
	endpoint := "configuration/v1/user_group_entity_mapping/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result UserGroupEntityMappingListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetUserGroupEntityMapping retrieves a specific user group entity mapping by ID
func (s *Service) GetUserGroupEntityMapping(ctx context.Context, id int) (*UserGroupEntityMapping, error) {
	endpoint := fmt.Sprintf("configuration/v1/user_group_entity_mapping/%d/", id)

	var result UserGroupEntityMapping
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateUserGroupEntityMapping creates a new user group entity mapping
func (s *Service) CreateUserGroupEntityMapping(ctx context.Context, req *UserGroupEntityMappingCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/user_group_entity_mapping/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateUserGroupEntityMapping updates an existing user group entity mapping
func (s *Service) UpdateUserGroupEntityMapping(ctx context.Context, id int, req *UserGroupEntityMappingUpdateRequest) (*UserGroupEntityMapping, error) {
	endpoint := fmt.Sprintf("configuration/v1/user_group_entity_mapping/%d/", id)

	var result UserGroupEntityMapping
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteUserGroupEntityMapping deletes a user group entity mapping
func (s *Service) DeleteUserGroupEntityMapping(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/user_group_entity_mapping/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
