package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListEndUsers retrieves a list of end users
func (s *Service) ListEndUsers(ctx context.Context, opts *ListOptions) (*EndUserListResponse, error) {
	endpoint := "configuration/v1/end_user/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result EndUserListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetEndUser retrieves a specific end user by ID
func (s *Service) GetEndUser(ctx context.Context, id int) (*EndUser, error) {
	endpoint := fmt.Sprintf("configuration/v1/end_user/%d/", id)

	var result EndUser
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateEndUser creates a new end user
func (s *Service) CreateEndUser(ctx context.Context, req *EndUserCreateRequest) (*EndUser, error) {
	endpoint := "configuration/v1/end_user/"

	var result EndUser
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// CreateEndUserWithResponse creates a new end user and returns both the response body and location header
func (s *Service) CreateEndUserWithResponse(ctx context.Context, req *EndUserCreateRequest) (*EndUser, *types.PostResponse, error) {
	endpoint := "configuration/v1/end_user/"

	var result EndUser
	postResp, err := s.client.PostWithResponse(ctx, endpoint, req, &result)
	if err != nil {
		return nil, postResp, err
	}
	return &result, postResp, nil
}

// UpdateEndUser updates an existing end user
func (s *Service) UpdateEndUser(ctx context.Context, id int, req *EndUserUpdateRequest) (*EndUser, error) {
	endpoint := fmt.Sprintf("configuration/v1/end_user/%d/", id)

	var result EndUser
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteEndUser deletes an end user
func (s *Service) DeleteEndUser(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/end_user/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
