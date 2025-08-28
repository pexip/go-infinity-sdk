/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// ConferenceAlias represents a conference alias configuration
type ConferenceAlias struct {
	ID           int               `json:"id,omitempty"`
	Alias        string            `json:"alias"`
	Conference   string            `json:"conference"`
	Description  string            `json:"description,omitempty"`
	CreationTime util.InfinityTime `json:"creation_time,omitempty"`
	ResourceURI  string            `json:"resource_uri,omitempty"`
}

// ConferenceAliasCreateRequest represents a request to create a conference alias
type ConferenceAliasCreateRequest struct {
	Alias       string `json:"alias"`
	Conference  string `json:"conference"`
	Description string `json:"description,omitempty"`
}

// ConferenceAliasUpdateRequest represents a request to update a conference alias
type ConferenceAliasUpdateRequest struct {
	Alias       string `json:"alias,omitempty"`
	Conference  string `json:"conference,omitempty"`
	Description string `json:"description,omitempty"`
}

// ConferenceAliasListResponse represents the response from listing conference aliases
type ConferenceAliasListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ConferenceAlias `json:"objects"`
}
