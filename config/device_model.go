/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// Device represents a device configuration
type Device struct {
	ID                          int               `json:"id,omitempty"`
	Alias                       string            `json:"alias"`
	Description                 string            `json:"description,omitempty"`
	Username                    string            `json:"username,omitempty"`
	Password                    string            `json:"password,omitempty"`
	PrimaryOwnerEmailAddress    string            `json:"primary_owner_email_address,omitempty"`
	EnableSIP                   bool              `json:"enable_sip"`
	EnableH323                  bool              `json:"enable_h323"`
	EnableInfinityConnectNonSSO bool              `json:"enable_infinity_connect_non_sso"`
	EnableInfinityConnectSSO    bool              `json:"enable_infinity_connect_sso"`
	EnableStandardSSO           bool              `json:"enable_standard_sso"`
	SSOIdentityProviderGroup    *string           `json:"sso_identity_provider_group,omitempty"`
	Tag                         string            `json:"tag,omitempty"`
	SyncTag                     string            `json:"sync_tag,omitempty"`
	CreationTime                util.InfinityTime `json:"creation_time,omitempty"`
	ResourceURI                 string            `json:"resource_uri,omitempty"`
}

// DeviceCreateRequest represents a request to create a device
type DeviceCreateRequest struct {
	Alias                       string  `json:"alias"`
	Description                 string  `json:"description,omitempty"`
	Username                    string  `json:"username,omitempty"`
	Password                    string  `json:"password,omitempty"`
	PrimaryOwnerEmailAddress    string  `json:"primary_owner_email_address,omitempty"`
	EnableSIP                   bool    `json:"enable_sip"`
	EnableH323                  bool    `json:"enable_h323"`
	EnableInfinityConnectNonSSO bool    `json:"enable_infinity_connect_non_sso"`
	EnableInfinityConnectSSO    bool    `json:"enable_infinity_connect_sso"`
	EnableStandardSSO           bool    `json:"enable_standard_sso"`
	SSOIdentityProviderGroup    *string `json:"sso_identity_provider_group,omitempty"`
	Tag                         string  `json:"tag,omitempty"`
	SyncTag                     string  `json:"sync_tag,omitempty"`
}

// DeviceUpdateRequest represents a request to update a device
type DeviceUpdateRequest struct {
	Alias                       string  `json:"alias,omitempty"`
	Description                 string  `json:"description,omitempty"`
	Username                    string  `json:"username,omitempty"`
	Password                    string  `json:"password,omitempty"`
	PrimaryOwnerEmailAddress    string  `json:"primary_owner_email_address,omitempty"`
	EnableSIP                   *bool   `json:"enable_sip,omitempty"`
	EnableH323                  *bool   `json:"enable_h323,omitempty"`
	EnableInfinityConnectNonSSO *bool   `json:"enable_infinity_connect_non_sso,omitempty"`
	EnableInfinityConnectSSO    *bool   `json:"enable_infinity_connect_sso,omitempty"`
	EnableStandardSSO           *bool   `json:"enable_standard_sso,omitempty"`
	SSOIdentityProviderGroup    *string `json:"sso_identity_provider_group,omitempty"`
	Tag                         string  `json:"tag,omitempty"`
	SyncTag                     string  `json:"sync_tag,omitempty"`
}

// DeviceListResponse represents the response from listing devices
type DeviceListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []Device `json:"objects"`
}
