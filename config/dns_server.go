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

// ListDNSServers retrieves a list of DNS servers
func (s *Service) ListDNSServers(ctx context.Context, opts *ListOptions) (*DNSServerListResponse, error) {
	endpoint := "configuration/v1/dns_server/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result DNSServerListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetDNSServer retrieves a specific DNS server by ID
func (s *Service) GetDNSServer(ctx context.Context, id int) (*DNSServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/dns_server/%d/", id)

	var result DNSServer
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateDNSServer creates a new DNS server
func (s *Service) CreateDNSServer(ctx context.Context, req *DNSServerCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/dns_server/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateDNSServer updates an existing DNS server
func (s *Service) UpdateDNSServer(ctx context.Context, id int, req *DNSServerUpdateRequest) (*DNSServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/dns_server/%d/", id)

	var result DNSServer
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteDNSServer deletes a DNS server
func (s *Service) DeleteDNSServer(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/dns_server/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
