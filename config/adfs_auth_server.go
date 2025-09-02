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

// ListADFSAuthServers retrieves a list of AD FS OAuth 2.0 Clients
func (s *Service) ListADFSAuthServers(ctx context.Context, opts *ListOptions) (*ADFSAuthServerListResponse, error) {
	endpoint := "configuration/v1/adfs_auth_server/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result ADFSAuthServerListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetADFSAuthServer retrieves a specific AD FS OAuth 2.0 Client by ID
func (s *Service) GetADFSAuthServer(ctx context.Context, id int) (*ADFSAuthServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/adfs_auth_server/%d/", id)

	var result ADFSAuthServer
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateADFSAuthServer creates a new AD FS OAuth 2.0 Client
func (s *Service) CreateADFSAuthServer(ctx context.Context, req *ADFSAuthServerCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/adfs_auth_server/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateADFSAuthServer updates an existing AD FS OAuth 2.0 Client
func (s *Service) UpdateADFSAuthServer(ctx context.Context, id int, req *ADFSAuthServerUpdateRequest) (*ADFSAuthServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/adfs_auth_server/%d/", id)

	var result ADFSAuthServer
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteADFSAuthServer deletes an AD FS OAuth 2.0 Client
func (s *Service) DeleteADFSAuthServer(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/adfs_auth_server/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
