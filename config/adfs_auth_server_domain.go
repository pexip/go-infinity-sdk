/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListADFSAuthServerDomains retrieves a list of AD FS OAuth 2.0 Client domains
func (s *Service) ListADFSAuthServerDomains(ctx context.Context, opts *ListOptions) (*ADFSAuthServerDomainListResponse, error) {
	endpoint := "configuration/v1/adfs_auth_server_domain/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result ADFSAuthServerDomainListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetADFSAuthServerDomain retrieves a specific AD FS OAuth 2.0 Client domain by ID
func (s *Service) GetADFSAuthServerDomain(ctx context.Context, id int) (*ADFSAuthServerDomain, error) {
	endpoint := fmt.Sprintf("configuration/v1/adfs_auth_server_domain/%d/", id)

	var result ADFSAuthServerDomain
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateADFSAuthServerDomain creates a new AD FS OAuth 2.0 Client domain
func (s *Service) CreateADFSAuthServerDomain(ctx context.Context, req *ADFSAuthServerDomainCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/adfs_auth_server_domain/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateADFSAuthServerDomain updates an existing AD FS OAuth 2.0 Client domain
func (s *Service) UpdateADFSAuthServerDomain(ctx context.Context, id int, req *ADFSAuthServerDomainUpdateRequest) (*ADFSAuthServerDomain, error) {
	endpoint := fmt.Sprintf("configuration/v1/adfs_auth_server_domain/%d/", id)

	var result ADFSAuthServerDomain
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteADFSAuthServerDomain deletes an AD FS OAuth 2.0 Client domain
func (s *Service) DeleteADFSAuthServerDomain(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/adfs_auth_server_domain/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
