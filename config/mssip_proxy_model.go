package config

// MSSIPProxy represents an MS-SIP proxy configuration
type MSSIPProxy struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address"`
	Port        *int   `json:"port,omitempty"`
	Transport   string `json:"transport"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// MSSIPProxyCreateRequest represents a request to create an MS-SIP proxy
type MSSIPProxyCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address"`
	Port        *int   `json:"port,omitempty"`
	Transport   string `json:"transport"`
}

// MSSIPProxyUpdateRequest represents a request to update an MS-SIP proxy
type MSSIPProxyUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address,omitempty"`
	Port        *int   `json:"port,omitempty"`
	Transport   string `json:"transport,omitempty"`
}

// MSSIPProxyListResponse represents the response from listing MS-SIP proxies
type MSSIPProxyListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MSSIPProxy `json:"objects"`
}
