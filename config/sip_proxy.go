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

// ListSIPProxies retrieves a list of SIP proxies
func (s *Service) ListSIPProxies(ctx context.Context, opts *ListOptions) (*SIPProxyListResponse, error) {
	endpoint := "configuration/v1/sip_proxy/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result SIPProxyListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetSIPProxy retrieves a specific SIP proxy by ID
func (s *Service) GetSIPProxy(ctx context.Context, id int) (*SIPProxy, error) {
	endpoint := fmt.Sprintf("configuration/v1/sip_proxy/%d/", id)

	var result SIPProxy
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateSIPProxy creates a new SIP proxy
func (s *Service) CreateSIPProxy(ctx context.Context, req *SIPProxyCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/sip_proxy/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateSIPProxy updates an existing SIP proxy
func (s *Service) UpdateSIPProxy(ctx context.Context, id int, req *SIPProxyUpdateRequest) (*SIPProxy, error) {
	endpoint := fmt.Sprintf("configuration/v1/sip_proxy/%d/", id)

	var result SIPProxy
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteSIPProxy deletes a SIP proxy
func (s *Service) DeleteSIPProxy(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/sip_proxy/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
