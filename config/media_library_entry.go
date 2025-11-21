/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListMediaLibraryEntries retrieves a list of media library entries
func (s *Service) ListMediaLibraryEntries(ctx context.Context, opts *ListOptions) (*MediaLibraryEntryListResponse, error) {
	endpoint := "configuration/v1/media_library_entry/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result MediaLibraryEntryListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetMediaLibraryEntry retrieves a specific media library entry by ID
func (s *Service) GetMediaLibraryEntry(ctx context.Context, id int) (*MediaLibraryEntry, error) {
	endpoint := fmt.Sprintf("configuration/v1/media_library_entry/%d/", id)

	var result MediaLibraryEntry
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateMediaLibraryEntry creates a new media library entry
func (s *Service) CreateMediaLibraryEntry(ctx context.Context, req *MediaLibraryEntryCreateRequest, filename string, file io.Reader) (*types.PostResponse, error) {
	endpoint := "configuration/v1/media_library_entry/"

	// Create form fields from request
	fields := map[string]string{
		"name":        req.Name,
		"description": req.Description,
	}

	// Use PostMultipartFormWithFieldsAndResponse to send all data in a single request
	return s.client.PostMultipartFormWithFieldsAndResponse(ctx, endpoint, fields, "media_file", filename, file, nil)
}

// UpdateMediaLibraryEntry updates an existing media library entry
func (s *Service) UpdateMediaLibraryEntry(ctx context.Context, id int, req *MediaLibraryEntryUpdateRequest, filename string, file io.Reader) (*MediaLibraryEntry, error) {
	endpoint := fmt.Sprintf("configuration/v1/media_library_entry/%d/", id)

	// Create form fields from request
	fields := map[string]string{
		"name":        req.Name,
		"uuid":        req.UUID,
		"description": req.Description,
	}

	var result MediaLibraryEntry
	// Use PatchMultipartFormWithFieldsAndResponse to send all data in a single request
	_, err := s.client.PatchMultipartFormWithFieldsAndResponse(ctx, endpoint, fields, "media_file", filename, file, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteMediaLibraryEntry deletes a media library entry
func (s *Service) DeleteMediaLibraryEntry(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/media_library_entry/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
