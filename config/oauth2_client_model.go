/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// OAuth2Client represents an OAuth2 client configuration
type OAuth2Client struct {
	ClientID      string  `json:"client_id,omitempty"`
	ClientName    string  `json:"client_name"`
	Role          string  `json:"role"`
	PrivateKeyJWT *string `json:"private_key_jwt,omitempty"`
	ResourceURI   string  `json:"resource_uri,omitempty"`
}

// OAuth2ClientCreateRequest represents a request to create an OAuth2 client
type OAuth2ClientCreateRequest struct {
	ClientName string `json:"client_name"`
	Role       string `json:"role"`
}

// OAuth2ClientUpdateRequest represents a request to update an OAuth2 client
type OAuth2ClientUpdateRequest struct {
	ClientName string `json:"client_name,omitempty"`
	Role       string `json:"role,omitempty"`
}

// OAuth2ClientListResponse represents the response from listing OAuth2 clients
type OAuth2ClientListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []OAuth2Client `json:"objects"`
}
