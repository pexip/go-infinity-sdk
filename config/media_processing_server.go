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

// ListMediaProcessingServers retrieves a list of media processing servers
func (s *Service) ListMediaProcessingServers(ctx context.Context, opts *ListOptions) (*MediaProcessingServerListResponse, error) {
	endpoint := "configuration/v1/media_processing_server/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MediaProcessingServerListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMediaProcessingServer retrieves a specific media processing server by ID
func (s *Service) GetMediaProcessingServer(ctx context.Context, id int) (*MediaProcessingServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/media_processing_server/%d/", id)

	var result MediaProcessingServer
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateMediaProcessingServer creates a new media processing server
func (s *Service) CreateMediaProcessingServer(ctx context.Context, req *MediaProcessingServerCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/media_processing_server/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMediaProcessingServer updates an existing media processing server
func (s *Service) UpdateMediaProcessingServer(ctx context.Context, id int, req *MediaProcessingServerUpdateRequest) (*MediaProcessingServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/media_processing_server/%d/", id)

	var result MediaProcessingServer
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMediaProcessingServer deletes a media processing server
func (s *Service) DeleteMediaProcessingServer(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/media_processing_server/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
