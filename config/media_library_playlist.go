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

// ListMediaLibraryPlaylists retrieves a list of media library playlists
func (s *Service) ListMediaLibraryPlaylists(ctx context.Context, opts *ListOptions) (*MediaLibraryPlaylistListResponse, error) {
	endpoint := "configuration/v1/media_library_playlist/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MediaLibraryPlaylistListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMediaLibraryPlaylist retrieves a specific media library playlist by ID
func (s *Service) GetMediaLibraryPlaylist(ctx context.Context, id int) (*MediaLibraryPlaylist, error) {
	endpoint := fmt.Sprintf("configuration/v1/media_library_playlist/%d/", id)

	var result MediaLibraryPlaylist
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateMediaLibraryPlaylist creates a new media library playlist
func (s *Service) CreateMediaLibraryPlaylist(ctx context.Context, req *MediaLibraryPlaylistCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/media_library_playlist/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMediaLibraryPlaylist updates an existing media library playlist
func (s *Service) UpdateMediaLibraryPlaylist(ctx context.Context, id int, req *MediaLibraryPlaylistUpdateRequest) (*MediaLibraryPlaylist, error) {
	endpoint := fmt.Sprintf("configuration/v1/media_library_playlist/%d/", id)

	var result MediaLibraryPlaylist
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMediaLibraryPlaylist deletes a media library playlist
func (s *Service) DeleteMediaLibraryPlaylist(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/media_library_playlist/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
