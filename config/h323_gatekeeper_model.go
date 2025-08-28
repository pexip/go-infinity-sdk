/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// H323Gatekeeper represents an H.323 gatekeeper configuration
type H323Gatekeeper struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address"`
	Port        *int   `json:"port,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// H323GatekeeperCreateRequest represents a request to create an H.323 gatekeeper
type H323GatekeeperCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address"`
	Port        *int   `json:"port,omitempty"`
}

// H323GatekeeperUpdateRequest represents a request to update an H.323 gatekeeper
type H323GatekeeperUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address,omitempty"`
	Port        *int   `json:"port,omitempty"`
}

// H323GatekeeperListResponse represents the response from listing H.323 gatekeepers
type H323GatekeeperListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []H323Gatekeeper `json:"objects"`
}
