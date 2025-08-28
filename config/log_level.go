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

// ListLogLevels retrieves a list of log levels
func (s *Service) ListLogLevels(ctx context.Context, opts *ListOptions) (*LogLevelListResponse, error) {
	endpoint := "configuration/v1/log_level/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result LogLevelListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetLogLevel retrieves a specific log level by ID
func (s *Service) GetLogLevel(ctx context.Context, id int) (*LogLevel, error) {
	endpoint := fmt.Sprintf("configuration/v1/log_level/%d/", id)

	var result LogLevel
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateLogLevel creates a new log level
func (s *Service) CreateLogLevel(ctx context.Context, req *LogLevelCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/log_level/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateLogLevel updates an existing log level
func (s *Service) UpdateLogLevel(ctx context.Context, id int, req *LogLevelUpdateRequest) (*LogLevel, error) {
	endpoint := fmt.Sprintf("configuration/v1/log_level/%d/", id)

	var result LogLevel
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteLogLevel deletes a log level
func (s *Service) DeleteLogLevel(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/log_level/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
