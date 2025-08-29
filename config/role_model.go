/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// Role represents a role configuration
type Role struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name"`
	Permissions []Permission `json:"permissions,omitempty"`
	ResourceURI string   `json:"resource_uri,omitempty"`
}

// RoleCreateRequest represents a request to create a role
type RoleCreateRequest struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions,omitempty"`
}

// RoleUpdateRequest represents a request to update a role
type RoleUpdateRequest struct {
	Name        string   `json:"name,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// RoleListResponse represents the response from listing roles
type RoleListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []Role `json:"objects"`
}
