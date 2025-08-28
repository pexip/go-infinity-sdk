/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// IdentityProviderAttribute represents an identity provider attribute configuration
type IdentityProviderAttribute struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// IdentityProviderAttributeCreateRequest represents a request to create an identity provider attribute
type IdentityProviderAttributeCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// IdentityProviderAttributeUpdateRequest represents a request to update an identity provider attribute
type IdentityProviderAttributeUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// IdentityProviderAttributeListResponse represents the response from listing identity provider attributes
type IdentityProviderAttributeListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []IdentityProviderAttribute `json:"objects"`
}
