/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// PexipStreamingCredential represents a Pexip Streaming credential configuration
type PexipStreamingCredential struct {
	ID          int    `json:"id,omitempty"`
	Kid         string `json:"kid"`
	PublicKey   string `json:"public_key"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// PexipStreamingCredentialCreateRequest represents a request to create a Pexip Streaming credential
type PexipStreamingCredentialCreateRequest struct {
	Kid       string `json:"kid"`
	PublicKey string `json:"public_key"`
}

// PexipStreamingCredentialUpdateRequest represents a request to update a Pexip Streaming credential
type PexipStreamingCredentialUpdateRequest struct {
	Kid       string `json:"kid,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

// PexipStreamingCredentialListResponse represents the response from listing Pexip Streaming credentials
type PexipStreamingCredentialListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []PexipStreamingCredential `json:"objects"`
}
