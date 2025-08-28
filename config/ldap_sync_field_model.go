/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// LdapSyncField represents an LDAP sync field configuration
type LdapSyncField struct {
	ID                   int    `json:"id,omitempty"`
	Name                 string `json:"name"`
	Description          string `json:"description,omitempty"`
	TemplateVariableName string `json:"template_variable_name"`
	IsBinary             bool   `json:"is_binary"`
	ResourceURI          string `json:"resource_uri,omitempty"`
}

// LdapSyncFieldCreateRequest represents a request to create an LDAP sync field
type LdapSyncFieldCreateRequest struct {
	Name                 string `json:"name"`
	Description          string `json:"description,omitempty"`
	TemplateVariableName string `json:"template_variable_name"`
	IsBinary             bool   `json:"is_binary"`
}

// LdapSyncFieldUpdateRequest represents a request to update an LDAP sync field
type LdapSyncFieldUpdateRequest struct {
	Name                 string `json:"name,omitempty"`
	Description          string `json:"description,omitempty"`
	TemplateVariableName string `json:"template_variable_name,omitempty"`
	IsBinary             *bool  `json:"is_binary,omitempty"`
}

// LdapSyncFieldListResponse represents the response from listing LDAP sync fields
type LdapSyncFieldListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []LdapSyncField `json:"objects"`
}
