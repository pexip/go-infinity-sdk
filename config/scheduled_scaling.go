package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListScheduledScalings retrieves a list of scheduled scaling policies
func (s *Service) ListScheduledScalings(ctx context.Context, opts *ListOptions) (*ScheduledScalingListResponse, error) {
	endpoint := "configuration/v1/scheduled_scaling/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ScheduledScalingListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetScheduledScaling retrieves a specific scheduled scaling policy by ID
func (s *Service) GetScheduledScaling(ctx context.Context, id int) (*ScheduledScaling, error) {
	endpoint := fmt.Sprintf("configuration/v1/scheduled_scaling/%d/", id)

	var result ScheduledScaling
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateScheduledScaling creates a new scheduled scaling policy
func (s *Service) CreateScheduledScaling(ctx context.Context, req *ScheduledScalingCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/scheduled_scaling/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateScheduledScaling updates an existing scheduled scaling policy
func (s *Service) UpdateScheduledScaling(ctx context.Context, id int, req *ScheduledScalingUpdateRequest) (*ScheduledScaling, error) {
	endpoint := fmt.Sprintf("configuration/v1/scheduled_scaling/%d/", id)

	var result ScheduledScaling
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteScheduledScaling deletes a scheduled scaling policy
func (s *Service) DeleteScheduledScaling(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/scheduled_scaling/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
