/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// GoogleAuthServerDomain represents a domain associated with a Google OAuth 2.0 Credential
type GoogleAuthServerDomain struct {
	ID               int    `json:"id,omitempty"`
	Domain           string `json:"domain"`
	Description      string `json:"description,omitempty"`
	GoogleAuthServer string `json:"google_auth_server"`
	ResourceURI      string `json:"resource_uri,omitempty"`
}

// GoogleAuthServerDomainCreateRequest represents a request to create a Google OAuth 2.0 Credential domain
type GoogleAuthServerDomainCreateRequest struct {
	Domain           string `json:"domain"`
	Description      string `json:"description,omitempty"`
	GoogleAuthServer string `json:"google_auth_server"`
}

// GoogleAuthServerDomainUpdateRequest represents a request to update a Google OAuth 2.0 Credential domain
type GoogleAuthServerDomainUpdateRequest struct {
	Domain           string `json:"domain,omitempty"`
	Description      string `json:"description,omitempty"`
	GoogleAuthServer string `json:"google_auth_server,omitempty"`
}

// GoogleAuthServerDomainListResponse represents the response from listing Google OAuth 2.0 Credential domains
type GoogleAuthServerDomainListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []GoogleAuthServerDomain `json:"objects"`
}
