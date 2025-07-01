package config

// Location represents a location configuration
type Location struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// LocationCreateRequest represents a request to create a location
type LocationCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// LocationUpdateRequest represents a request to update a location
type LocationUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// LocationListResponse represents the response from listing locations
type LocationListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []Location `json:"objects"`
}
