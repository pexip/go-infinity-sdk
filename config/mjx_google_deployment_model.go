package config

// MjxGoogleDeployment represents a MJX Google deployment configuration
type MjxGoogleDeployment struct {
	ID                         int      `json:"id,omitempty"`
	Name                       string   `json:"name"`
	Description                string   `json:"description,omitempty"`
	ClientEmail                string   `json:"client_email"`
	ClientID                   string   `json:"client_id,omitempty"`
	ClientSecret               string   `json:"client_secret,omitempty"`
	PrivateKey                 string   `json:"private_key,omitempty"`
	UseUserConsent             bool     `json:"use_user_consent"`
	AuthEndpoint               string   `json:"auth_endpoint,omitempty"`
	TokenEndpoint              string   `json:"token_endpoint,omitempty"`
	RedirectURI                string   `json:"redirect_uri,omitempty"`
	RefreshToken               string   `json:"refresh_token,omitempty"`
	OAuthState                 *string  `json:"oauth_state,omitempty"`
	MaximumNumberOfAPIRequests int      `json:"maximum_number_of_api_requests"`
	MjxIntegrations            []string `json:"mjx_integrations,omitempty"`
	ResourceURI                string   `json:"resource_uri,omitempty"`
}

// MjxGoogleDeploymentCreateRequest represents a request to create a MJX Google deployment
type MjxGoogleDeploymentCreateRequest struct {
	Name                       string `json:"name"`
	Description                string `json:"description,omitempty"`
	ClientEmail                string `json:"client_email"`
	ClientID                   string `json:"client_id,omitempty"`
	ClientSecret               string `json:"client_secret,omitempty"`
	PrivateKey                 string `json:"private_key,omitempty"`
	UseUserConsent             bool   `json:"use_user_consent"`
	AuthEndpoint               string `json:"auth_endpoint,omitempty"`
	TokenEndpoint              string `json:"token_endpoint,omitempty"`
	RedirectURI                string `json:"redirect_uri,omitempty"`
	RefreshToken               string `json:"refresh_token,omitempty"`
	MaximumNumberOfAPIRequests int    `json:"maximum_number_of_api_requests"`
}

// MjxGoogleDeploymentUpdateRequest represents a request to update a MJX Google deployment
type MjxGoogleDeploymentUpdateRequest struct {
	Name                       string `json:"name,omitempty"`
	Description                string `json:"description,omitempty"`
	ClientEmail                string `json:"client_email,omitempty"`
	ClientID                   string `json:"client_id,omitempty"`
	ClientSecret               string `json:"client_secret,omitempty"`
	PrivateKey                 string `json:"private_key,omitempty"`
	UseUserConsent             *bool  `json:"use_user_consent,omitempty"`
	AuthEndpoint               string `json:"auth_endpoint,omitempty"`
	TokenEndpoint              string `json:"token_endpoint,omitempty"`
	RedirectURI                string `json:"redirect_uri,omitempty"`
	RefreshToken               string `json:"refresh_token,omitempty"`
	MaximumNumberOfAPIRequests *int   `json:"maximum_number_of_api_requests,omitempty"`
}

// MjxGoogleDeploymentListResponse represents the response from listing MJX Google deployments
type MjxGoogleDeploymentListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MjxGoogleDeployment `json:"objects"`
}
