package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListConferenceSyncTemplates retrieves a list of conference sync templates
func (s *Service) ListConferenceSyncTemplates(ctx context.Context, opts *ListOptions) (*ConferenceSyncTemplateListResponse, error) {
	endpoint := "configuration/v1/conference_sync_template/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ConferenceSyncTemplateListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetConferenceSyncTemplate retrieves a specific conference sync template by ID
func (s *Service) GetConferenceSyncTemplate(ctx context.Context, id int) (*ConferenceSyncTemplate, error) {
	endpoint := fmt.Sprintf("configuration/v1/conference_sync_template/%d/", id)

	var result ConferenceSyncTemplate
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateConferenceSyncTemplate creates a new conference sync template
func (s *Service) CreateConferenceSyncTemplate(ctx context.Context, req *ConferenceSyncTemplateCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/conference_sync_template/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateConferenceSyncTemplate updates an existing conference sync template
func (s *Service) UpdateConferenceSyncTemplate(ctx context.Context, id int, req *ConferenceSyncTemplateUpdateRequest) (*ConferenceSyncTemplate, error) {
	endpoint := fmt.Sprintf("configuration/v1/conference_sync_template/%d/", id)

	var result ConferenceSyncTemplate
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteConferenceSyncTemplate deletes a conference sync template
func (s *Service) DeleteConferenceSyncTemplate(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/conference_sync_template/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
