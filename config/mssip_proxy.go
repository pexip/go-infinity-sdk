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

// ListMSSIPProxies retrieves a list of MS-SIP proxies
func (s *Service) ListMSSIPProxies(ctx context.Context, opts *ListOptions) (*MSSIPProxyListResponse, error) {
	endpoint := "configuration/v1/mssip_proxy/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result MSSIPProxyListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetMSSIPProxy retrieves a specific MS-SIP proxy by ID
func (s *Service) GetMSSIPProxy(ctx context.Context, id int) (*MSSIPProxy, error) {
	endpoint := fmt.Sprintf("configuration/v1/mssip_proxy/%d/", id)

	var result MSSIPProxy
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateMSSIPProxy creates a new MS-SIP proxy
func (s *Service) CreateMSSIPProxy(ctx context.Context, req *MSSIPProxyCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/mssip_proxy/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMSSIPProxy updates an existing MS-SIP proxy
func (s *Service) UpdateMSSIPProxy(ctx context.Context, id int, req *MSSIPProxyUpdateRequest) (*MSSIPProxy, error) {
	endpoint := fmt.Sprintf("configuration/v1/mssip_proxy/%d/", id)

	var result MSSIPProxy
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMSSIPProxy deletes an MS-SIP proxy
func (s *Service) DeleteMSSIPProxy(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/mssip_proxy/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
