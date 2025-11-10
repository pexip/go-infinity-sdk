/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"github.com/pexip/go-infinity-sdk/v38/util"
)

// IVRTheme represents an IVR theme configuration
type IVRTheme struct {
	ID             int               `json:"id,omitempty"`
	Name           string            `json:"name"`
	UUID           string            `json:"uuid,omitempty"`
	Conference     []string          `json:"conference,omitempty"`
	CustomLayouts  string            `json:"custom_layouts,omitempty"`
	PinningConfigs string            `json:"pinning_configs,omitempty"`
	LastUpdated    util.InfinityTime `json:"last_updated,omitempty"`
	ResourceURI    string            `json:"resource_uri,omitempty"`
}

// IVRThemeCreateRequest represents a request to create an IVR theme
type IVRThemeCreateRequest struct {
	Name           string   `json:"name"`
	UUID           string   `json:"uuid,omitempty"`
	Conference     []string `json:"conference,omitempty"`
	CustomLayouts  string   `json:"custom_layouts,omitempty"`
	PinningConfigs string   `json:"pinning_configs,omitempty"`
}

// IVRThemeUpdateRequest represents a request to update an IVR theme
type IVRThemeUpdateRequest struct {
	Name           string   `json:"name,omitempty"`
	UUID           string   `json:"uuid,omitempty"`
	Conference     []string `json:"conference,omitempty"`
	CustomLayouts  string   `json:"custom_layouts,omitempty"`
	PinningConfigs string   `json:"pinning_configs,omitempty"`
}

// IVRThemeListResponse represents the response from listing IVR themes
type IVRThemeListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []IVRTheme `json:"objects"`
}
