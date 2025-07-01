package config

// MediaLibraryPlaylist represents a media library playlist configuration
type MediaLibraryPlaylist struct {
	ID              int      `json:"id,omitempty"`
	Name            string   `json:"name"`
	Description     string   `json:"description,omitempty"`
	Loop            bool     `json:"loop"`
	Shuffle         bool     `json:"shuffle"`
	PlaylistEntries []string `json:"playlist_entries,omitempty"`
	ResourceURI     string   `json:"resource_uri,omitempty"`
}

// MediaLibraryPlaylistCreateRequest represents a request to create a media library playlist
type MediaLibraryPlaylistCreateRequest struct {
	Name            string   `json:"name"`
	Description     string   `json:"description,omitempty"`
	Loop            bool     `json:"loop"`
	Shuffle         bool     `json:"shuffle"`
	PlaylistEntries []string `json:"playlist_entries,omitempty"`
}

// MediaLibraryPlaylistUpdateRequest represents a request to update a media library playlist
type MediaLibraryPlaylistUpdateRequest struct {
	Name            string   `json:"name,omitempty"`
	Description     string   `json:"description,omitempty"`
	Loop            *bool    `json:"loop,omitempty"`
	Shuffle         *bool    `json:"shuffle,omitempty"`
	PlaylistEntries []string `json:"playlist_entries,omitempty"`
}

// MediaLibraryPlaylistListResponse represents the response from listing media library playlists
type MediaLibraryPlaylistListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MediaLibraryPlaylist `json:"objects"`
}
