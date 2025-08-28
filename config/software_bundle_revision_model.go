/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// SoftwareBundleRevision represents a software bundle revision configuration
type SoftwareBundleRevision struct {
	ID          int    `json:"id,omitempty"`
	BundleType  string `json:"bundle_type"`
	Core        bool   `json:"core"`
	Revision    string `json:"revision"`
	Version     string `json:"version"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// SoftwareBundleRevisionListResponse represents the response from listing software bundle revisions
type SoftwareBundleRevisionListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SoftwareBundleRevision `json:"objects"`
}
