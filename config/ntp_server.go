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

// ListNTPServers retrieves a list of NTP servers
func (s *Service) ListNTPServers(ctx context.Context, opts *ListOptions) (*NTPServerListResponse, error) {
	endpoint := "configuration/v1/ntp_server/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result NTPServerListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetNTPServer retrieves a specific NTP server by ID
func (s *Service) GetNTPServer(ctx context.Context, id int) (*NTPServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/ntp_server/%d/", id)

	var result NTPServer
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateNTPServer creates a new NTP server
func (s *Service) CreateNTPServer(ctx context.Context, req *NTPServerCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/ntp_server/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateNTPServer updates an existing NTP server
func (s *Service) UpdateNTPServer(ctx context.Context, id int, req *NTPServerUpdateRequest) (*NTPServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/ntp_server/%d/", id)

	var result NTPServer
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteNTPServer deletes an NTP server
func (s *Service) DeleteNTPServer(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/ntp_server/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
