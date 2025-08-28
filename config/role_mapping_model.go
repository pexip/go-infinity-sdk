/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// RoleMapping represents a role mapping configuration
type RoleMapping struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name"`
	Source      string   `json:"source"`
	Value       string   `json:"value"`
	Roles       []string `json:"roles,omitempty"`
	ResourceURI string   `json:"resource_uri,omitempty"`
}

// RoleMappingCreateRequest represents a request to create a role mapping
type RoleMappingCreateRequest struct {
	Name   string   `json:"name"`
	Source string   `json:"source"`
	Value  string   `json:"value"`
	Roles  []string `json:"roles,omitempty"`
}

// RoleMappingUpdateRequest represents a request to update a role mapping
type RoleMappingUpdateRequest struct {
	Name   string   `json:"name,omitempty"`
	Source string   `json:"source,omitempty"`
	Value  string   `json:"value,omitempty"`
	Roles  []string `json:"roles,omitempty"`
}

// RoleMappingListResponse represents the response from listing role mappings
type RoleMappingListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []RoleMapping `json:"objects"`
}
