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

// ListMediaLibraryPlaylistEntries retrieves a list of media library playlist entries
func (s *Service) ListMediaLibraryPlaylistEntries(ctx context.Context, opts *ListOptions) (*MediaLibraryPlaylistEntryListResponse, error) {
	endpoint := "configuration/v1/media_library_playlist_entry/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MediaLibraryPlaylistEntryListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMediaLibraryPlaylistEntry retrieves a specific media library playlist entry by ID
func (s *Service) GetMediaLibraryPlaylistEntry(ctx context.Context, id int) (*MediaLibraryPlaylistEntry, error) {
	endpoint := fmt.Sprintf("configuration/v1/media_library_playlist_entry/%d/", id)

	var result MediaLibraryPlaylistEntry
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateMediaLibraryPlaylistEntry creates a new media library playlist entry
func (s *Service) CreateMediaLibraryPlaylistEntry(ctx context.Context, req *MediaLibraryPlaylistEntryCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/media_library_playlist_entry/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMediaLibraryPlaylistEntry updates an existing media library playlist entry
func (s *Service) UpdateMediaLibraryPlaylistEntry(ctx context.Context, id int, req *MediaLibraryPlaylistEntryUpdateRequest) (*MediaLibraryPlaylistEntry, error) {
	endpoint := fmt.Sprintf("configuration/v1/media_library_playlist_entry/%d/", id)

	var result MediaLibraryPlaylistEntry
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMediaLibraryPlaylistEntry deletes a media library playlist entry
func (s *Service) DeleteMediaLibraryPlaylistEntry(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/media_library_playlist_entry/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
