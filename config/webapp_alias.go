package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListWebappAliases retrieves a list of web app aliases
func (s *Service) ListWebappAliases(ctx context.Context, opts *ListOptions) (*WebappAliasListResponse, error) {
	endpoint := "configuration/v1/webapp_alias/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result WebappAliasListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetWebappAlias retrieves a specific web app alias by ID
func (s *Service) GetWebappAlias(ctx context.Context, id int) (*WebappAlias, error) {
	endpoint := fmt.Sprintf("configuration/v1/webapp_alias/%d/", id)

	var result WebappAlias
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateWebappAlias creates a new web app alias
func (s *Service) CreateWebappAlias(ctx context.Context, req *WebappAliasCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/webapp_alias/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateWebappAlias updates an existing web app alias
func (s *Service) UpdateWebappAlias(ctx context.Context, id int, req *WebappAliasUpdateRequest) (*WebappAlias, error) {
	endpoint := fmt.Sprintf("configuration/v1/webapp_alias/%d/", id)

	var result WebappAlias
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteWebappAlias deletes a web app alias
func (s *Service) DeleteWebappAlias(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/webapp_alias/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
