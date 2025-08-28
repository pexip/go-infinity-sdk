/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// ADFSAuthServer represents an AD FS OAuth 2.0 Client configuration
type ADFSAuthServer struct {
	ID                             int    `json:"id,omitempty"`
	Name                           string `json:"name"`
	Description                    string `json:"description,omitempty"`
	ClientID                       string `json:"client_id"`
	FederationServiceName          string `json:"federation_service_name"`
	FederationServiceIdentifier    string `json:"federation_service_identifier"`
	RelyingPartyTrustIdentifierURL string `json:"relying_party_trust_identifier_url"`
	ResourceURI                    string `json:"resource_uri,omitempty"`
}

// ADFSAuthServerCreateRequest represents a request to create an AD FS OAuth 2.0 Client
type ADFSAuthServerCreateRequest struct {
	Name                           string `json:"name"`
	Description                    string `json:"description,omitempty"`
	ClientID                       string `json:"client_id"`
	FederationServiceName          string `json:"federation_service_name"`
	FederationServiceIdentifier    string `json:"federation_service_identifier"`
	RelyingPartyTrustIdentifierURL string `json:"relying_party_trust_identifier_url"`
}

// ADFSAuthServerUpdateRequest represents a request to update an AD FS OAuth 2.0 Client
type ADFSAuthServerUpdateRequest struct {
	Name                           string `json:"name,omitempty"`
	Description                    string `json:"description,omitempty"`
	ClientID                       string `json:"client_id,omitempty"`
	FederationServiceName          string `json:"federation_service_name,omitempty"`
	FederationServiceIdentifier    string `json:"federation_service_identifier,omitempty"`
	RelyingPartyTrustIdentifierURL string `json:"relying_party_trust_identifier_url,omitempty"`
}

// ADFSAuthServerListResponse represents the response from listing AD FS OAuth 2.0 Clients
type ADFSAuthServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ADFSAuthServer `json:"objects"`
}
