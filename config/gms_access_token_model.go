package config

// GMSAccessToken represents a Google Meet access token configuration
type GMSAccessToken struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Token       string `json:"token,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// GMSAccessTokenCreateRequest represents a request to create a Google Meet access token
type GMSAccessTokenCreateRequest struct {
	Name  string `json:"name"`
	Token string `json:"token,omitempty"`
}

// GMSAccessTokenUpdateRequest represents a request to update a Google Meet access token
type GMSAccessTokenUpdateRequest struct {
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
}

// GMSAccessTokenListResponse represents the response from listing Google Meet access tokens
type GMSAccessTokenListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []GMSAccessToken `json:"objects"`
}
