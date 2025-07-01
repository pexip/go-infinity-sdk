package config

// GoogleAuthServer represents a Google OAuth 2.0 Credential configuration
type GoogleAuthServer struct {
	ID              int     `json:"id,omitempty"`
	Name            string  `json:"name"`
	Description     string  `json:"description,omitempty"`
	ApplicationType string  `json:"application_type"`
	ClientID        *string `json:"client_id,omitempty"`
	ClientSecret    string  `json:"client_secret,omitempty"`
	ResourceURI     string  `json:"resource_uri,omitempty"`
}

// GoogleAuthServerCreateRequest represents a request to create a Google OAuth 2.0 Credential
type GoogleAuthServerCreateRequest struct {
	Name            string  `json:"name"`
	Description     string  `json:"description,omitempty"`
	ApplicationType string  `json:"application_type"`
	ClientID        *string `json:"client_id,omitempty"`
	ClientSecret    string  `json:"client_secret,omitempty"`
}

// GoogleAuthServerUpdateRequest represents a request to update a Google OAuth 2.0 Credential
type GoogleAuthServerUpdateRequest struct {
	Name            string  `json:"name,omitempty"`
	Description     string  `json:"description,omitempty"`
	ApplicationType string  `json:"application_type,omitempty"`
	ClientID        *string `json:"client_id,omitempty"`
	ClientSecret    string  `json:"client_secret,omitempty"`
}

// GoogleAuthServerListResponse represents the response from listing Google OAuth 2.0 Credentials
type GoogleAuthServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []GoogleAuthServer `json:"objects"`
}
