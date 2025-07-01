package config

// MediaLibraryPlaylistEntry represents a media library playlist entry configuration
type MediaLibraryPlaylistEntry struct {
	ID          int     `json:"id,omitempty"`
	EntryType   string  `json:"entry_type"`
	Media       *string `json:"media,omitempty"`
	Playlist    *string `json:"playlist,omitempty"`
	Position    int     `json:"position"`
	Playcount   int     `json:"playcount"`
	ResourceURI string  `json:"resource_uri,omitempty"`
}

// MediaLibraryPlaylistEntryCreateRequest represents a request to create a media library playlist entry
type MediaLibraryPlaylistEntryCreateRequest struct {
	EntryType string  `json:"entry_type"`
	Media     *string `json:"media,omitempty"`
	Playlist  *string `json:"playlist,omitempty"`
	Position  int     `json:"position"`
	Playcount int     `json:"playcount"`
}

// MediaLibraryPlaylistEntryUpdateRequest represents a request to update a media library playlist entry
type MediaLibraryPlaylistEntryUpdateRequest struct {
	EntryType string  `json:"entry_type,omitempty"`
	Media     *string `json:"media,omitempty"`
	Playlist  *string `json:"playlist,omitempty"`
	Position  *int    `json:"position,omitempty"`
	Playcount *int    `json:"playcount,omitempty"`
}

// MediaLibraryPlaylistEntryListResponse represents the response from listing media library playlist entries
type MediaLibraryPlaylistEntryListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MediaLibraryPlaylistEntry `json:"objects"`
}
