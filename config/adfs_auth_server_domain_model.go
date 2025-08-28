/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// ADFSAuthServerDomain represents a domain associated with an AD FS OAuth 2.0 Client
type ADFSAuthServerDomain struct {
	ID             int    `json:"id,omitempty"`
	Domain         string `json:"domain"`
	Description    string `json:"description,omitempty"`
	ADFSAuthServer string `json:"adfs_auth_server"`
	ResourceURI    string `json:"resource_uri,omitempty"`
}

// ADFSAuthServerDomainCreateRequest represents a request to create an AD FS OAuth 2.0 Client domain
type ADFSAuthServerDomainCreateRequest struct {
	Domain         string `json:"domain"`
	Description    string `json:"description,omitempty"`
	ADFSAuthServer string `json:"adfs_auth_server"`
}

// ADFSAuthServerDomainUpdateRequest represents a request to update an AD FS OAuth 2.0 Client domain
type ADFSAuthServerDomainUpdateRequest struct {
	Domain         string `json:"domain,omitempty"`
	Description    string `json:"description,omitempty"`
	ADFSAuthServer string `json:"adfs_auth_server,omitempty"`
}

// ADFSAuthServerDomainListResponse represents the response from listing AD FS OAuth 2.0 Client domains
type ADFSAuthServerDomainListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ADFSAuthServerDomain `json:"objects"`
}
