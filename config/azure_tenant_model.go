/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// AzureTenant represents a Microsoft Teams tenant configuration
type AzureTenant struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	TenantID    string `json:"tenant_id"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// AzureTenantCreateRequest represents a request to create a Microsoft Teams tenant
type AzureTenantCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	TenantID    string `json:"tenant_id"`
}

// AzureTenantUpdateRequest represents a request to update a Microsoft Teams tenant
type AzureTenantUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
	TenantID    string `json:"tenant_id,omitempty"`
}

// AzureTenantListResponse represents the response from listing Microsoft Teams tenants
type AzureTenantListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []AzureTenant `json:"objects"`
}
