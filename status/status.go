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

	"github.com/pexip/go-infinity-sdk/v42/options"
)

// HTTPClient is the subset of HTTP operations required by the status service.
type HTTPClient interface {
	// GetJSON performs a GET request and unmarshals the JSON response
	GetJSON(ctx context.Context, endpoint string, queryParams *url.Values, result interface{}) error
}

// Service handles status API endpoints
type Service struct {
	client HTTPClient
}

// New creates a new status API service
func New(client HTTPClient) *Service {
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
