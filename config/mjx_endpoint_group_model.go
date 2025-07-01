package config

// MjxEndpointGroup represents a MJX endpoint group configuration
type MjxEndpointGroup struct {
	ID             int      `json:"id,omitempty"`
	Name           string   `json:"name"`
	Description    string   `json:"description,omitempty"`
	MjxIntegration *string  `json:"mjx_integration,omitempty"`
	SystemLocation *string  `json:"system_location,omitempty"`
	DisableProxy   bool     `json:"disable_proxy"`
	Endpoints      []string `json:"endpoints,omitempty"`
	ResourceURI    string   `json:"resource_uri,omitempty"`
}

// MjxEndpointGroupCreateRequest represents a request to create a MJX endpoint group
type MjxEndpointGroupCreateRequest struct {
	Name           string  `json:"name"`
	Description    string  `json:"description,omitempty"`
	MjxIntegration *string `json:"mjx_integration,omitempty"`
	SystemLocation *string `json:"system_location,omitempty"`
	DisableProxy   bool    `json:"disable_proxy"`
}

// MjxEndpointGroupUpdateRequest represents a request to update a MJX endpoint group
type MjxEndpointGroupUpdateRequest struct {
	Name           string  `json:"name,omitempty"`
	Description    string  `json:"description,omitempty"`
	MjxIntegration *string `json:"mjx_integration,omitempty"`
	SystemLocation *string `json:"system_location,omitempty"`
	DisableProxy   *bool   `json:"disable_proxy,omitempty"`
}

// MjxEndpointGroupListResponse represents the response from listing MJX endpoint groups
type MjxEndpointGroupListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MjxEndpointGroup `json:"objects"`
}
