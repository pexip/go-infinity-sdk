package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListIdentityProviderAttributes retrieves a list of identity provider attributes
func (s *Service) ListIdentityProviderAttributes(ctx context.Context, opts *ListOptions) (*IdentityProviderAttributeListResponse, error) {
	endpoint := "configuration/v1/identity_provider_attribute/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result IdentityProviderAttributeListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetIdentityProviderAttribute retrieves a specific identity provider attribute by ID
func (s *Service) GetIdentityProviderAttribute(ctx context.Context, id int) (*IdentityProviderAttribute, error) {
	endpoint := fmt.Sprintf("configuration/v1/identity_provider_attribute/%d/", id)

	var result IdentityProviderAttribute
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateIdentityProviderAttribute creates a new identity provider attribute
func (s *Service) CreateIdentityProviderAttribute(ctx context.Context, req *IdentityProviderAttributeCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/identity_provider_attribute/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateIdentityProviderAttribute updates an existing identity provider attribute
func (s *Service) UpdateIdentityProviderAttribute(ctx context.Context, id int, req *IdentityProviderAttributeUpdateRequest) (*IdentityProviderAttribute, error) {
	endpoint := fmt.Sprintf("configuration/v1/identity_provider_attribute/%d/", id)

	var result IdentityProviderAttribute
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteIdentityProviderAttribute deletes an identity provider attribute
func (s *Service) DeleteIdentityProviderAttribute(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/identity_provider_attribute/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
