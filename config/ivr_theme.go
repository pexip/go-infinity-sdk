package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListIVRThemes retrieves a list of IVR themes
func (s *Service) ListIVRThemes(ctx context.Context, opts *ListOptions) (*IVRThemeListResponse, error) {
	endpoint := "configuration/v1/ivr_theme/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result IVRThemeListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetIVRTheme retrieves a specific IVR theme by ID
func (s *Service) GetIVRTheme(ctx context.Context, id int) (*IVRTheme, error) {
	endpoint := fmt.Sprintf("configuration/v1/ivr_theme/%d/", id)

	var result IVRTheme
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateIVRTheme creates a new IVR theme
func (s *Service) CreateIVRTheme(ctx context.Context, req *IVRThemeCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/ivr_theme/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateIVRTheme updates an existing IVR theme
func (s *Service) UpdateIVRTheme(ctx context.Context, id int, req *IVRThemeUpdateRequest) (*IVRTheme, error) {
	endpoint := fmt.Sprintf("configuration/v1/ivr_theme/%d/", id)

	var result IVRTheme
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteIVRTheme deletes an IVR theme
func (s *Service) DeleteIVRTheme(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/ivr_theme/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
