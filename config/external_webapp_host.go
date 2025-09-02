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

// ListExternalWebappHosts retrieves a list of external web app hosts
func (s *Service) ListExternalWebappHosts(ctx context.Context, opts *ListOptions) (*ExternalWebappHostListResponse, error) {
	endpoint := "configuration/v1/external_webapp_host/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result ExternalWebappHostListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetExternalWebappHost retrieves a specific external web app host by ID
func (s *Service) GetExternalWebappHost(ctx context.Context, id int) (*ExternalWebappHost, error) {
	endpoint := fmt.Sprintf("configuration/v1/external_webapp_host/%d/", id)

	var result ExternalWebappHost
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateExternalWebappHost creates a new external web app host
func (s *Service) CreateExternalWebappHost(ctx context.Context, req *ExternalWebappHostCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/external_webapp_host/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateExternalWebappHost updates an existing external web app host
func (s *Service) UpdateExternalWebappHost(ctx context.Context, id int, req *ExternalWebappHostUpdateRequest) (*ExternalWebappHost, error) {
	endpoint := fmt.Sprintf("configuration/v1/external_webapp_host/%d/", id)

	var result ExternalWebappHost
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteExternalWebappHost deletes an external web app host
func (s *Service) DeleteExternalWebappHost(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/external_webapp_host/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
