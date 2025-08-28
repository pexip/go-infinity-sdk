/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// SystemSyncpoint represents a system syncpoint configuration
type SystemSyncpoint struct {
	ID           int               `json:"id,omitempty"`
	CreationTime util.InfinityTime `json:"creation_time"`
	ResourceURI  string            `json:"resource_uri,omitempty"`
}

// SystemSyncpointCreateRequest represents a request to create a system syncpoint
type SystemSyncpointCreateRequest struct {
	// No fields needed for creation
}

// SystemSyncpointListResponse represents the response from listing system syncpoints
type SystemSyncpointListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SystemSyncpoint `json:"objects"`
}
