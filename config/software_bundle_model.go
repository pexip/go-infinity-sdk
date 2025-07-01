package config

// SoftwareBundle represents a software bundle configuration
type SoftwareBundle struct {
	ID               int     `json:"id,omitempty"`
	BundleType       string  `json:"bundle_type"`
	SelectedRevision *string `json:"selected_revision,omitempty"`
	ResourceURI      string  `json:"resource_uri,omitempty"`
}

// SoftwareBundleUpdateRequest represents a request to update a software bundle
type SoftwareBundleUpdateRequest struct {
	SelectedRevision *string `json:"selected_revision,omitempty"`
}

// SoftwareBundleListResponse represents the response from listing software bundles
type SoftwareBundleListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SoftwareBundle `json:"objects"`
}
