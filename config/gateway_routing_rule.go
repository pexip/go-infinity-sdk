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

// ListGatewayRoutingRules retrieves a list of gateway routing rules
func (s *Service) ListGatewayRoutingRules(ctx context.Context, opts *ListOptions) (*GatewayRoutingRuleListResponse, error) {
	endpoint := "configuration/v1/gateway_routing_rule/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result GatewayRoutingRuleListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetGatewayRoutingRule retrieves a specific gateway routing rule by ID
func (s *Service) GetGatewayRoutingRule(ctx context.Context, id int) (*GatewayRoutingRule, error) {
	endpoint := fmt.Sprintf("configuration/v1/gateway_routing_rule/%d/", id)

	var result GatewayRoutingRule
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateGatewayRoutingRule creates a new gateway routing rule
func (s *Service) CreateGatewayRoutingRule(ctx context.Context, req *GatewayRoutingRuleCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/gateway_routing_rule/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateGatewayRoutingRule updates an existing gateway routing rule
func (s *Service) UpdateGatewayRoutingRule(ctx context.Context, id int, req *GatewayRoutingRuleUpdateRequest) (*GatewayRoutingRule, error) {
	endpoint := fmt.Sprintf("configuration/v1/gateway_routing_rule/%d/", id)

	var result GatewayRoutingRule
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteGatewayRoutingRule deletes a gateway routing rule
func (s *Service) DeleteGatewayRoutingRule(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/gateway_routing_rule/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
