/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// IdentityProviderGroup represents an identity provider group configuration
type IdentityProviderGroup struct {
	ID               int      `json:"id,omitempty"`
	Name             string   `json:"name"`
	Description      string   `json:"description,omitempty"`
	IdentityProvider []string `json:"identity_provider,omitempty"`
	ResourceURI      string   `json:"resource_uri,omitempty"`
}

// IdentityProviderGroupCreateRequest represents a request to create an identity provider group
type IdentityProviderGroupCreateRequest struct {
	Name             string   `json:"name"`
	Description      string   `json:"description,omitempty"`
	IdentityProvider []string `json:"identity_provider,omitempty"`
}

// IdentityProviderGroupUpdateRequest represents a request to update an identity provider group
type IdentityProviderGroupUpdateRequest struct {
	Name             string   `json:"name,omitempty"`
	Description      string   `json:"description,omitempty"`
	IdentityProvider []string `json:"identity_provider,omitempty"`
}

// IdentityProviderGroupListResponse represents the response from listing identity provider groups
type IdentityProviderGroupListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []IdentityProviderGroup `json:"objects"`
}
