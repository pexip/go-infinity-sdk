/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// STUNServer represents a STUN server configuration
type STUNServer struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address"`
	Port        int    `json:"port"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// STUNServerCreateRequest represents a request to create a STUN server
type STUNServerCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address"`
	Port        int    `json:"port"`
}

// STUNServerUpdateRequest represents a request to update a STUN server
type STUNServerUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address,omitempty"`
	Port        *int   `json:"port,omitempty"`
}

// STUNServerListResponse represents the response from listing STUN servers
type STUNServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []STUNServer `json:"objects"`
}
