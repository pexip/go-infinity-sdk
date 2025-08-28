/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// MediaLibraryEntry represents a media library entry configuration
type MediaLibraryEntry struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	FileName    string `json:"file_name,omitempty"`
	MediaType   string `json:"media_type,omitempty"`
	MediaFormat string `json:"media_format,omitempty"`
	MediaSize   int    `json:"media_size,omitempty"`
	MediaFile   string `json:"media_file,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// MediaLibraryEntryCreateRequest represents a request to create a media library entry
type MediaLibraryEntryCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	MediaFile   string `json:"media_file"`
}

// MediaLibraryEntryUpdateRequest represents a request to update a media library entry
type MediaLibraryEntryUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	MediaFile   string `json:"media_file,omitempty"`
}

// MediaLibraryEntryListResponse represents the response from listing media library entries
type MediaLibraryEntryListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MediaLibraryEntry `json:"objects"`
}
