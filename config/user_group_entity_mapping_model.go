/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// UserGroupEntityMapping represents a user group entity mapping configuration
type UserGroupEntityMapping struct {
	ID                int    `json:"id,omitempty"`
	Description       string `json:"description,omitempty"`
	EntityResourceURI string `json:"entity_resource_uri"`
	UserGroup         string `json:"user_group"`
	ResourceURI       string `json:"resource_uri,omitempty"`
}

// UserGroupEntityMappingCreateRequest represents a request to create a user group entity mapping
type UserGroupEntityMappingCreateRequest struct {
	Description       string `json:"description,omitempty"`
	EntityResourceURI string `json:"entity_resource_uri"`
	UserGroup         string `json:"user_group"`
}

// UserGroupEntityMappingUpdateRequest represents a request to update a user group entity mapping
type UserGroupEntityMappingUpdateRequest struct {
	Description       string `json:"description,omitempty"`
	EntityResourceURI string `json:"entity_resource_uri,omitempty"`
	UserGroup         string `json:"user_group,omitempty"`
}

// UserGroupEntityMappingListResponse represents the response from listing user group entity mappings
type UserGroupEntityMappingListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []UserGroupEntityMapping `json:"objects"`
}
