package config

// SystemTuneable represents a system tuneable configuration
type SystemTuneable struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Setting     string `json:"setting"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// SystemTuneableCreateRequest represents a request to create a system tuneable
type SystemTuneableCreateRequest struct {
	Name    string `json:"name"`
	Setting string `json:"setting"`
}

// SystemTuneableUpdateRequest represents a request to update a system tuneable
type SystemTuneableUpdateRequest struct {
	Name    string `json:"name,omitempty"`
	Setting string `json:"setting,omitempty"`
}

// SystemTuneableListResponse represents the response from listing system tuneables
type SystemTuneableListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SystemTuneable `json:"objects"`
}
