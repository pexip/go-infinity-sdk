package config

// SIPProxy represents a SIP proxy configuration
type SIPProxy struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address"`
	Port        *int   `json:"port,omitempty"`
	Transport   string `json:"transport"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// SIPProxyCreateRequest represents a request to create a SIP proxy
type SIPProxyCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address"`
	Port        *int   `json:"port,omitempty"`
	Transport   string `json:"transport"`
}

// SIPProxyUpdateRequest represents a request to update a SIP proxy
type SIPProxyUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address,omitempty"`
	Port        *int   `json:"port,omitempty"`
	Transport   string `json:"transport,omitempty"`
}

// SIPProxyListResponse represents the response from listing SIP proxies
type SIPProxyListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SIPProxy `json:"objects"`
}
