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

// ListSnmpNetworkManagementSystems retrieves a list of SNMP network management systems
func (s *Service) ListSnmpNetworkManagementSystems(ctx context.Context, opts *ListOptions) (*SnmpNetworkManagementSystemListResponse, error) {
	endpoint := "configuration/v1/snmp_network_management_system/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result SnmpNetworkManagementSystemListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetSnmpNetworkManagementSystem retrieves a specific SNMP network management system by ID
func (s *Service) GetSnmpNetworkManagementSystem(ctx context.Context, id int) (*SnmpNetworkManagementSystem, error) {
	endpoint := fmt.Sprintf("configuration/v1/snmp_network_management_system/%d/", id)

	var result SnmpNetworkManagementSystem
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateSnmpNetworkManagementSystem creates a new SNMP network management system
func (s *Service) CreateSnmpNetworkManagementSystem(ctx context.Context, req *SnmpNetworkManagementSystemCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/snmp_network_management_system/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateSnmpNetworkManagementSystem updates an existing SNMP network management system
func (s *Service) UpdateSnmpNetworkManagementSystem(ctx context.Context, id int, req *SnmpNetworkManagementSystemUpdateRequest) (*SnmpNetworkManagementSystem, error) {
	endpoint := fmt.Sprintf("configuration/v1/snmp_network_management_system/%d/", id)

	var result SnmpNetworkManagementSystem
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteSnmpNetworkManagementSystem deletes an SNMP network management system
func (s *Service) DeleteSnmpNetworkManagementSystem(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/snmp_network_management_system/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
