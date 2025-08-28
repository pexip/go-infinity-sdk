/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// BreakInAllowListAddress represents a break-in attempt IP allow list entry
type BreakInAllowListAddress struct {
	ID                     int    `json:"id,omitempty"`
	Name                   string `json:"name"`
	Description            string `json:"description,omitempty"`
	Address                string `json:"address"`
	Prefix                 int    `json:"prefix"`
	AllowlistEntryType     string `json:"allowlist_entry_type"`
	IgnoreIncorrectAliases bool   `json:"ignore_incorrect_aliases"`
	IgnoreIncorrectPins    bool   `json:"ignore_incorrect_pins"`
	ResourceURI            string `json:"resource_uri,omitempty"`
}

// BreakInAllowListAddressCreateRequest represents a request to create a break-in attempt IP allow list entry
type BreakInAllowListAddressCreateRequest struct {
	Name                   string `json:"name"`
	Description            string `json:"description,omitempty"`
	Address                string `json:"address"`
	Prefix                 int    `json:"prefix"`
	AllowlistEntryType     string `json:"allowlist_entry_type"`
	IgnoreIncorrectAliases bool   `json:"ignore_incorrect_aliases"`
	IgnoreIncorrectPins    bool   `json:"ignore_incorrect_pins"`
}

// BreakInAllowListAddressUpdateRequest represents a request to update a break-in attempt IP allow list entry
type BreakInAllowListAddressUpdateRequest struct {
	Name                   string `json:"name,omitempty"`
	Description            string `json:"description,omitempty"`
	Address                string `json:"address,omitempty"`
	Prefix                 *int   `json:"prefix,omitempty"`
	AllowlistEntryType     string `json:"allowlist_entry_type,omitempty"`
	IgnoreIncorrectAliases *bool  `json:"ignore_incorrect_aliases,omitempty"`
	IgnoreIncorrectPins    *bool  `json:"ignore_incorrect_pins,omitempty"`
}

// BreakInAllowListAddressListResponse represents the response from listing break-in attempt IP allow list entries
type BreakInAllowListAddressListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []BreakInAllowListAddress `json:"objects"`
}
