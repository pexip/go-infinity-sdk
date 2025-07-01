package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListIdentityProviderGroups retrieves a list of identity provider groups
func (s *Service) ListIdentityProviderGroups(ctx context.Context, opts *ListOptions) (*IdentityProviderGroupListResponse, error) {
	endpoint := "configuration/v1/identity_provider_group/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result IdentityProviderGroupListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetIdentityProviderGroup retrieves a specific identity provider group by ID
func (s *Service) GetIdentityProviderGroup(ctx context.Context, id int) (*IdentityProviderGroup, error) {
	endpoint := fmt.Sprintf("configuration/v1/identity_provider_group/%d/", id)

	var result IdentityProviderGroup
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateIdentityProviderGroup creates a new identity provider group
func (s *Service) CreateIdentityProviderGroup(ctx context.Context, req *IdentityProviderGroupCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/identity_provider_group/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateIdentityProviderGroup updates an existing identity provider group
func (s *Service) UpdateIdentityProviderGroup(ctx context.Context, id int, req *IdentityProviderGroupUpdateRequest) (*IdentityProviderGroup, error) {
	endpoint := fmt.Sprintf("configuration/v1/identity_provider_group/%d/", id)

	var result IdentityProviderGroup
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteIdentityProviderGroup deletes an identity provider group
func (s *Service) DeleteIdentityProviderGroup(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/identity_provider_group/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
