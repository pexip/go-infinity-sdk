/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// WebappBranding represents a webapp branding configuration
type WebappBranding struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	UUID        string            `json:"uuid"`
	WebappType  string            `json:"webapp_type"`
	IsDefault   bool              `json:"is_default"`
	LastUpdated util.InfinityTime `json:"last_updated"`
	ResourceURI string            `json:"resource_uri,omitempty"`
}

// WebappBrandingCreateRequest represents a request to create a webapp branding
type WebappBrandingCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	WebappType  string `json:"webapp_type"`
}

// WebappBrandingUpdateRequest represents a request to update a webapp branding
type WebappBrandingUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	UUID        string `json:"uuid"`
	WebappType  string `json:"webapp_type"`
}

// WebappBrandingListResponse represents the response from listing webapp brandings
type WebappBrandingListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []WebappBranding `json:"objects"`
}
