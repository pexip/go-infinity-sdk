/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// NTPServer represents an NTP server configuration
type NTPServer struct {
	ID          int    `json:"id,omitempty"`
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
	Key         string `json:"key,omitempty"`
	KeyID       *int   `json:"key_id,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// NTPServerCreateRequest represents a request to create an NTP server
type NTPServerCreateRequest struct {
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
	Key         string `json:"key,omitempty"`
	KeyID       *int   `json:"key_id,omitempty"`
}

// NTPServerUpdateRequest represents a request to update an NTP server
type NTPServerUpdateRequest struct {
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
	Key         string `json:"key,omitempty"`
	KeyID       *int   `json:"key_id,omitempty"`
}

// NTPServerListResponse represents the response from listing NTP servers
type NTPServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []NTPServer `json:"objects"`
}
