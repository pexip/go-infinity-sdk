package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListUserGroups retrieves a list of user groups
func (s *Service) ListUserGroups(ctx context.Context, opts *ListOptions) (*UserGroupListResponse, error) {
	endpoint := "configuration/v1/user_group/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result UserGroupListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetUserGroup retrieves a specific user group by ID
func (s *Service) GetUserGroup(ctx context.Context, id int) (*UserGroup, error) {
	endpoint := fmt.Sprintf("configuration/v1/user_group/%d/", id)

	var result UserGroup
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateUserGroup creates a new user group
func (s *Service) CreateUserGroup(ctx context.Context, req *UserGroupCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/user_group/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateUserGroup updates an existing user group
func (s *Service) UpdateUserGroup(ctx context.Context, id int, req *UserGroupUpdateRequest) (*UserGroup, error) {
	endpoint := fmt.Sprintf("configuration/v1/user_group/%d/", id)

	var result UserGroup
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteUserGroup deletes a user group
func (s *Service) DeleteUserGroup(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/user_group/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
