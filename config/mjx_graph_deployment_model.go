/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// MjxGraphDeployment represents a MJX Graph deployment configuration
type MjxGraphDeployment struct {
	ID              int      `json:"id,omitempty"`
	Name            string   `json:"name"`
	Description     string   `json:"description,omitempty"`
	ClientID        string   `json:"client_id"`
	ClientSecret    string   `json:"client_secret,omitempty"`
	OAuthTokenURL   string   `json:"oauth_token_url"`
	GraphAPIDomain  string   `json:"graph_api_domain"`
	RequestQuota    int      `json:"request_quota"`
	DisableProxy    bool     `json:"disable_proxy"`
	MjxIntegrations []string `json:"mjx_integrations,omitempty"`
	ResourceURI     string   `json:"resource_uri,omitempty"`
}

// MjxGraphDeploymentCreateRequest represents a request to create a MJX Graph deployment
type MjxGraphDeploymentCreateRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret,omitempty"`
	OAuthTokenURL  string `json:"oauth_token_url"`
	GraphAPIDomain string `json:"graph_api_domain"`
	RequestQuota   int    `json:"request_quota"`
	DisableProxy   bool   `json:"disable_proxy"`
}

// MjxGraphDeploymentUpdateRequest represents a request to update a MJX Graph deployment
type MjxGraphDeploymentUpdateRequest struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	ClientID       string `json:"client_id,omitempty"`
	ClientSecret   string `json:"client_secret,omitempty"`
	OAuthTokenURL  string `json:"oauth_token_url,omitempty"`
	GraphAPIDomain string `json:"graph_api_domain,omitempty"`
	RequestQuota   *int   `json:"request_quota,omitempty"`
	DisableProxy   *bool  `json:"disable_proxy,omitempty"`
}

// MjxGraphDeploymentListResponse represents the response from listing MJX Graph deployments
type MjxGraphDeploymentListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MjxGraphDeployment `json:"objects"`
}
