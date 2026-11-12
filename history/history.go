/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// Package history provides access to the Pexip Infinity history API.
// It allows retrieval of historical data for conferences, participants, and media streams
// with support for time-based filtering and search capabilities.
package history

import (
	"context"
	"net/url"
	"time"
)

// HTTPClient is the subset of HTTP operations required by the history service.
type HTTPClient interface {
	// GetJSON performs a GET request and unmarshals the JSON response
	GetJSON(ctx context.Context, endpoint string, queryParams *url.Values, result interface{}) error
}

// Service handles history API endpoints
type Service struct {
	client HTTPClient
}

// New creates a new history API service
func New(client HTTPClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) listEndpoint(ctx context.Context, endpoint string, opts *ListOptions, result interface{}) error {
	var params url.Values
	if opts != nil {
		params = opts.ToURLValues()
	}
	return s.client.GetJSON(ctx, endpoint, &params, result)
}

func (s *Service) listEndpointWithSearchField(ctx context.Context, endpoint string, opts *ListOptions, searchField string, result interface{}) error {
	var params url.Values
	if opts != nil {
		if searchField != "" {
			params = opts.ToURLValuesWithSearchField(searchField)
		} else {
			params = opts.ToURLValues()
		}
	}
	return s.client.GetJSON(ctx, endpoint, &params, result)
}

func (s *Service) listEndpointWithTimeFilter(ctx context.Context, endpoint string, opts *ListOptions, result interface{}) error {
	var params url.Values
	if opts != nil {
		params = opts.ToURLValues()
		if opts.StartTime != nil {
			params.Set("start_time__gte", opts.StartTime.Format(time.RFC3339))
		}
		if opts.EndTime != nil {
			params.Set("end_time__lt", opts.EndTime.Format(time.RFC3339))
		}
	}
	return s.client.GetJSON(ctx, endpoint, &params, result)
}
