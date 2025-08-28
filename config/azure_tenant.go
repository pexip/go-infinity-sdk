/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListAzureTenants retrieves a list of Microsoft Teams tenants
func (s *Service) ListAzureTenants(ctx context.Context, opts *ListOptions) (*AzureTenantListResponse, error) {
	endpoint := "configuration/v1/azure_tenant/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result AzureTenantListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetAzureTenant retrieves a specific Microsoft Teams tenant by ID
func (s *Service) GetAzureTenant(ctx context.Context, id int) (*AzureTenant, error) {
	endpoint := fmt.Sprintf("configuration/v1/azure_tenant/%d/", id)

	var result AzureTenant
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateAzureTenant creates a new Microsoft Teams tenant
func (s *Service) CreateAzureTenant(ctx context.Context, req *AzureTenantCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/azure_tenant/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateAzureTenant updates an existing Microsoft Teams tenant
func (s *Service) UpdateAzureTenant(ctx context.Context, id int, req *AzureTenantUpdateRequest) (*AzureTenant, error) {
	endpoint := fmt.Sprintf("configuration/v1/azure_tenant/%d/", id)

	var result AzureTenant
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteAzureTenant deletes a Microsoft Teams tenant
func (s *Service) DeleteAzureTenant(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/azure_tenant/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
