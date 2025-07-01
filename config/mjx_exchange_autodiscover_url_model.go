package config

// MjxExchangeAutodiscoverURL represents a MJX Exchange autodiscover URL configuration
type MjxExchangeAutodiscoverURL struct {
	ID                 int     `json:"id,omitempty"`
	Name               string  `json:"name"`
	Description        string  `json:"description,omitempty"`
	URL                string  `json:"url"`
	ExchangeDeployment *string `json:"exchange_deployment,omitempty"`
	ResourceURI        string  `json:"resource_uri,omitempty"`
}

// MjxExchangeAutodiscoverURLCreateRequest represents a request to create a MJX Exchange autodiscover URL
type MjxExchangeAutodiscoverURLCreateRequest struct {
	Name               string  `json:"name"`
	Description        string  `json:"description,omitempty"`
	URL                string  `json:"url"`
	ExchangeDeployment *string `json:"exchange_deployment,omitempty"`
}

// MjxExchangeAutodiscoverURLUpdateRequest represents a request to update a MJX Exchange autodiscover URL
type MjxExchangeAutodiscoverURLUpdateRequest struct {
	Name               string  `json:"name,omitempty"`
	Description        string  `json:"description,omitempty"`
	URL                string  `json:"url,omitempty"`
	ExchangeDeployment *string `json:"exchange_deployment,omitempty"`
}

// MjxExchangeAutodiscoverURLListResponse represents the response from listing MJX Exchange autodiscover URLs
type MjxExchangeAutodiscoverURLListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MjxExchangeAutodiscoverURL `json:"objects"`
}
