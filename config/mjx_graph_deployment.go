package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListMjxGraphDeployments retrieves a list of MJX Graph deployments
func (s *Service) ListMjxGraphDeployments(ctx context.Context, opts *ListOptions) (*MjxGraphDeploymentListResponse, error) {
	endpoint := "configuration/v1/mjx_graph_deployment/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MjxGraphDeploymentListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMjxGraphDeployment retrieves a specific MJX Graph deployment by ID
func (s *Service) GetMjxGraphDeployment(ctx context.Context, id int) (*MjxGraphDeployment, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_graph_deployment/%d/", id)

	var result MjxGraphDeployment
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateMjxGraphDeployment creates a new MJX Graph deployment
func (s *Service) CreateMjxGraphDeployment(ctx context.Context, req *MjxGraphDeploymentCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/mjx_graph_deployment/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMjxGraphDeployment updates an existing MJX Graph deployment
func (s *Service) UpdateMjxGraphDeployment(ctx context.Context, id int, req *MjxGraphDeploymentUpdateRequest) (*MjxGraphDeployment, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_graph_deployment/%d/", id)

	var result MjxGraphDeployment
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMjxGraphDeployment deletes a MJX Graph deployment
func (s *Service) DeleteMjxGraphDeployment(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/mjx_graph_deployment/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
