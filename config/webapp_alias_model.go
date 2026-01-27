/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// WebappAlias represents a web app alias configuration
type WebappAlias struct {
	ID          int                     `json:"id,omitempty"`
	Slug        string                  `json:"slug"`
	Description string                  `json:"description,omitempty"`
	WebappType  string                  `json:"webapp_type"`
	IsEnabled   bool                    `json:"is_enabled"`
	Bundle      *SoftwareBundleRevision `json:"bundle,omitempty"`
	Branding    *WebappBranding         `json:"branding,omitempty"`
	ResourceURI string                  `json:"resource_uri,omitempty"`
}

// WebappAliasCreateRequest represents a request to create a web app alias
type WebappAliasCreateRequest struct {
	Slug        string  `json:"slug"`
	Description string  `json:"description,omitempty"`
	WebappType  string  `json:"webapp_type"`
	IsEnabled   bool    `json:"is_enabled"`
	Bundle      *string `json:"bundle,omitempty"`
	Branding    *string `json:"branding,omitempty"`
}

// WebappAliasUpdateRequest represents a request to update a web app alias
type WebappAliasUpdateRequest struct {
	Slug        string  `json:"slug,omitempty"`
	Description string  `json:"description,omitempty"`
	WebappType  string  `json:"webapp_type,omitempty"`
	IsEnabled   *bool   `json:"is_enabled,omitempty"`
	Bundle      *string `json:"bundle,omitempty"`
	Branding    *string `json:"branding,omitempty"`
}

// WebappAliasListResponse represents the response from listing web app aliases
type WebappAliasListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []WebappAlias `json:"objects"`
}
