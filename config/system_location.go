package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListSystemLocations retrieves a list of system locations
func (s *Service) ListSystemLocations(ctx context.Context, opts *ListOptions) (*SystemLocationListResponse, error) {
	endpoint := "configuration/v1/system_location/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result SystemLocationListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetSystemLocation retrieves a specific system location by ID
func (s *Service) GetSystemLocation(ctx context.Context, id int) (*SystemLocation, error) {
	endpoint := fmt.Sprintf("configuration/v1/system_location/%d/", id)

	var result SystemLocation
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateSystemLocation creates a new system location
func (s *Service) CreateSystemLocation(ctx context.Context, req *SystemLocationCreateRequest) (*SystemLocation, error) {
	endpoint := "configuration/v1/system_location/"

	var result SystemLocation
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// CreateSystemLocationWithResponse creates a new system location and returns both the response body and location header
func (s *Service) CreateSystemLocationWithResponse(ctx context.Context, req *SystemLocationCreateRequest) (*SystemLocation, *types.PostResponse, error) {
	endpoint := "configuration/v1/system_location/"

	var result SystemLocation
	postResp, err := s.client.PostWithResponse(ctx, endpoint, req, &result)
	if err != nil {
		return nil, postResp, err
	}
	return &result, postResp, nil
}

// UpdateSystemLocation updates an existing system location
func (s *Service) UpdateSystemLocation(ctx context.Context, id int, req *SystemLocationUpdateRequest) (*SystemLocation, error) {
	endpoint := fmt.Sprintf("configuration/v1/system_location/%d/", id)

	var result SystemLocation
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteSystemLocation deletes a system location
func (s *Service) DeleteSystemLocation(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/system_location/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
