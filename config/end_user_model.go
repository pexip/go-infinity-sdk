/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// EndUser represents an end user configuration
type EndUser struct {
	ID                  int      `json:"id,omitempty"`
	PrimaryEmailAddress string   `json:"primary_email_address"`
	FirstName           string   `json:"first_name,omitempty"`
	LastName            string   `json:"last_name,omitempty"`
	DisplayName         string   `json:"display_name,omitempty"`
	TelephoneNumber     string   `json:"telephone_number,omitempty"`
	MobileNumber        string   `json:"mobile_number,omitempty"`
	Title               string   `json:"title,omitempty"`
	Department          string   `json:"department,omitempty"`
	AvatarURL           string   `json:"avatar_url,omitempty"`
	UUID                string   `json:"uuid,omitempty"`
	UserGroups          []string `json:"user_groups,omitempty"`
	UserOID             *string  `json:"user_oid,omitempty"`
	ExchangeUserID      *string  `json:"exchange_user_id,omitempty"`
	MSExchangeGUID      *string  `json:"ms_exchange_guid,omitempty"`
	SyncTag             string   `json:"sync_tag,omitempty"`
	ResourceURI         string   `json:"resource_uri,omitempty"`
}

// EndUserCreateRequest represents a request to create an end user
type EndUserCreateRequest struct {
	PrimaryEmailAddress string   `json:"primary_email_address"`
	FirstName           string   `json:"first_name,omitempty"`
	LastName            string   `json:"last_name,omitempty"`
	DisplayName         string   `json:"display_name,omitempty"`
	TelephoneNumber     string   `json:"telephone_number,omitempty"`
	MobileNumber        string   `json:"mobile_number,omitempty"`
	Title               string   `json:"title,omitempty"`
	Department          string   `json:"department,omitempty"`
	AvatarURL           string   `json:"avatar_url,omitempty"`
	UserGroups          []string `json:"user_groups,omitempty"`
	UserOID             *string  `json:"user_oid,omitempty"`
	ExchangeUserID      *string  `json:"exchange_user_id,omitempty"`
	MSExchangeGUID      *string  `json:"ms_exchange_guid,omitempty"`
	SyncTag             string   `json:"sync_tag,omitempty"`
}

// EndUserUpdateRequest represents a request to update an end user
type EndUserUpdateRequest struct {
	PrimaryEmailAddress string   `json:"primary_email_address,omitempty"`
	FirstName           string   `json:"first_name,omitempty"`
	LastName            string   `json:"last_name,omitempty"`
	DisplayName         string   `json:"display_name,omitempty"`
	TelephoneNumber     string   `json:"telephone_number,omitempty"`
	MobileNumber        string   `json:"mobile_number,omitempty"`
	Title               string   `json:"title,omitempty"`
	Department          string   `json:"department,omitempty"`
	AvatarURL           string   `json:"avatar_url,omitempty"`
	UserGroups          []string `json:"user_groups,omitempty"`
	UserOID             *string  `json:"user_oid,omitempty"`
	ExchangeUserID      *string  `json:"exchange_user_id,omitempty"`
	MSExchangeGUID      *string  `json:"ms_exchange_guid,omitempty"`
	SyncTag             string   `json:"sync_tag,omitempty"`
}

// EndUserListResponse represents the response from listing end users
type EndUserListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []EndUser `json:"objects"`
}
