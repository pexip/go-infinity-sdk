package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListStaticRoutes retrieves a list of static routes
func (s *Service) ListStaticRoutes(ctx context.Context, opts *ListOptions) (*StaticRouteListResponse, error) {
	endpoint := "configuration/v1/static_route/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result StaticRouteListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetStaticRoute retrieves a specific static route by ID
func (s *Service) GetStaticRoute(ctx context.Context, id int) (*StaticRoute, error) {
	endpoint := fmt.Sprintf("configuration/v1/static_route/%d/", id)

	var result StaticRoute
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateStaticRoute creates a new static route
func (s *Service) CreateStaticRoute(ctx context.Context, req *StaticRouteCreateRequest) (*StaticRoute, error) {
	endpoint := "configuration/v1/static_route/"

	var result StaticRoute
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// CreateStaticRouteWithResponse creates a new static route and returns both the response body and location header
func (s *Service) CreateStaticRouteWithResponse(ctx context.Context, req *StaticRouteCreateRequest) (*StaticRoute, *types.PostResponse, error) {
	endpoint := "configuration/v1/static_route/"

	var result StaticRoute
	postResp, err := s.client.PostWithResponse(ctx, endpoint, req, &result)
	if err != nil {
		return nil, postResp, err
	}
	return &result, postResp, nil
}

// UpdateStaticRoute updates an existing static route
func (s *Service) UpdateStaticRoute(ctx context.Context, id int, req *StaticRouteUpdateRequest) (*StaticRoute, error) {
	endpoint := fmt.Sprintf("configuration/v1/static_route/%d/", id)

	var result StaticRoute
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteStaticRoute deletes a static route
func (s *Service) DeleteStaticRoute(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/static_route/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
