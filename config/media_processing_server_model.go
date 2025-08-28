/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// MediaProcessingServer represents a media processing server configuration
type MediaProcessingServer struct {
	ID           int     `json:"id,omitempty"`
	FQDN         string  `json:"fqdn"`
	AppID        *string `json:"app_id,omitempty"`
	PublicJWTKey string  `json:"public_jwt_key"`
	ResourceURI  string  `json:"resource_uri,omitempty"`
}

// MediaProcessingServerCreateRequest represents a request to create a media processing server
type MediaProcessingServerCreateRequest struct {
	FQDN         string  `json:"fqdn"`
	AppID        *string `json:"app_id,omitempty"`
	PublicJWTKey string  `json:"public_jwt_key"`
}

// MediaProcessingServerUpdateRequest represents a request to update a media processing server
type MediaProcessingServerUpdateRequest struct {
	FQDN         string  `json:"fqdn,omitempty"`
	AppID        *string `json:"app_id,omitempty"`
	PublicJWTKey string  `json:"public_jwt_key,omitempty"`
}

// MediaProcessingServerListResponse represents the response from listing media processing servers
type MediaProcessingServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MediaProcessingServer `json:"objects"`
}
