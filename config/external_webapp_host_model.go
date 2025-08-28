/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// ExternalWebappHost represents an external web app host configuration
type ExternalWebappHost struct {
	ID          int    `json:"id,omitempty"`
	Address     string `json:"address"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// ExternalWebappHostCreateRequest represents a request to create an external web app host
type ExternalWebappHostCreateRequest struct {
	Address string `json:"address"`
}

// ExternalWebappHostUpdateRequest represents a request to update an external web app host
type ExternalWebappHostUpdateRequest struct {
	Address string `json:"address,omitempty"`
}

// ExternalWebappHostListResponse represents the response from listing external web app hosts
type ExternalWebappHostListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ExternalWebappHost `json:"objects"`
}
