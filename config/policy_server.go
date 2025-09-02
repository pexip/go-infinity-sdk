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

// ListPolicyServers retrieves a list of policy servers
func (s *Service) ListPolicyServers(ctx context.Context, opts *ListOptions) (*PolicyServerListResponse, error) {
	endpoint := "configuration/v1/policy_server/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result PolicyServerListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetPolicyServer retrieves a specific policy server by ID
func (s *Service) GetPolicyServer(ctx context.Context, id int) (*PolicyServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/policy_server/%d/", id)

	var result PolicyServer
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreatePolicyServer creates a new policy server
func (s *Service) CreatePolicyServer(ctx context.Context, req *PolicyServerCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/policy_server/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdatePolicyServer updates an existing policy server
func (s *Service) UpdatePolicyServer(ctx context.Context, id int, req *PolicyServerUpdateRequest) (*PolicyServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/policy_server/%d/", id)

	var result PolicyServer
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeletePolicyServer deletes a policy server
func (s *Service) DeletePolicyServer(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/policy_server/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
