package config

// DNSServer represents a DNS server configuration
type DNSServer struct {
	ID          int    `json:"id,omitempty"`
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// DNSServerCreateRequest represents a request to create a DNS server
type DNSServerCreateRequest struct {
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
}

// DNSServerUpdateRequest represents a request to update a DNS server
type DNSServerUpdateRequest struct {
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
}

// DNSServerListResponse represents the response from listing DNS servers
type DNSServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []DNSServer `json:"objects"`
}
