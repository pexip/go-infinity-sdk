/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// SIPCredential represents a SIP credential configuration
type SIPCredential struct {
	ID          int    `json:"id,omitempty"`
	Realm       string `json:"realm"`
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// SIPCredentialCreateRequest represents a request to create a SIP credential
type SIPCredentialCreateRequest struct {
	Realm    string `json:"realm"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

// SIPCredentialUpdateRequest represents a request to update a SIP credential
type SIPCredentialUpdateRequest struct {
	Realm    string `json:"realm,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// SIPCredentialListResponse represents the response from listing SIP credentials
type SIPCredentialListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SIPCredential `json:"objects"`
}
