package config

// StaticRoute represents a static route configuration
type StaticRoute struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Prefix      int    `json:"prefix"`
	Gateway     string `json:"gateway"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// StaticRouteCreateRequest represents a request to create a static route
type StaticRouteCreateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Prefix  int    `json:"prefix"`
	Gateway string `json:"gateway"`
}

// StaticRouteUpdateRequest represents a request to update a static route
type StaticRouteUpdateRequest struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Prefix  int    `json:"prefix,omitempty"`
	Gateway string `json:"gateway,omitempty"`
}

// StaticRouteListResponse represents the response from listing static routes
type StaticRouteListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []StaticRoute `json:"objects"`
}
