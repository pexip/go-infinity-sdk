package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListMediaLibraryEntries retrieves a list of media library entries
func (s *Service) ListMediaLibraryEntries(ctx context.Context, opts *ListOptions) (*MediaLibraryEntryListResponse, error) {
	endpoint := "configuration/v1/media_library_entry/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MediaLibraryEntryListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMediaLibraryEntry retrieves a specific media library entry by ID
func (s *Service) GetMediaLibraryEntry(ctx context.Context, id int) (*MediaLibraryEntry, error) {
	endpoint := fmt.Sprintf("configuration/v1/media_library_entry/%d/", id)

	var result MediaLibraryEntry
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateMediaLibraryEntry creates a new media library entry
func (s *Service) CreateMediaLibraryEntry(ctx context.Context, req *MediaLibraryEntryCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/media_library_entry/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMediaLibraryEntry updates an existing media library entry
func (s *Service) UpdateMediaLibraryEntry(ctx context.Context, id int, req *MediaLibraryEntryUpdateRequest) (*MediaLibraryEntry, error) {
	endpoint := fmt.Sprintf("configuration/v1/media_library_entry/%d/", id)

	var result MediaLibraryEntry
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMediaLibraryEntry deletes a media library entry
func (s *Service) DeleteMediaLibraryEntry(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/media_library_entry/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
