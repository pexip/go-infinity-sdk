/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// Package status provides access to the Pexip Infinity status API.
// It allows monitoring of system status, active conferences, participants, worker nodes, and alarms
// with real-time status information and health monitoring capabilities.
package status

import (
	"context"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
)

// Service handles status API endpoints
type Service struct {
	client interfaces.HTTPClient
}

// New creates a new status API service
func New(client interfaces.HTTPClient) *Service {
	return &Service{
		client: client,
	}
}

// ListOptions contains options for listing resources
type ListOptions = options.BaseListOptions

func (s *Service) listEndpoint(ctx context.Context, endpoint string, opts *ListOptions, result interface{}) error {
	var params url.Values
	if opts != nil {
		params = opts.ToURLValues()
	}
	return s.client.GetJSON(ctx, endpoint, &params, result)
}
