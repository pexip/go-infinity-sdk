/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// TeamsProxy represents a Teams proxy configuration
type TeamsProxy struct {
	ID                   int                `json:"id"`
	Name                 string             `json:"name"`
	Description          string             `json:"description"`
	Address              string             `json:"address"`
	Port                 int                `json:"port"`
	AzureTenant          string             `json:"azure_tenant"`
	EventhubID           *string            `json:"eventhub_id"`
	MinNumberOfInstances int                `json:"min_number_of_instances"`
	NotificationsEnabled bool               `json:"notifications_enabled"`
	NotificationsQueue   *string            `json:"notifications_queue"`
	Updated              *util.InfinityTime `json:"updated"`
	ResourceURI          string             `json:"resource_uri"`
}

// TeamsProxyCreateRequest represents a request to create a Teams proxy
type TeamsProxyCreateRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Address              string  `json:"address"`
	Port                 int     `json:"port"`
	AzureTenant          string  `json:"azure_tenant"`
	EventhubID           *string `json:"eventhub_id"`
	MinNumberOfInstances int     `json:"min_number_of_instances"`
	NotificationsEnabled bool    `json:"notifications_enabled"`
	NotificationsQueue   *string `json:"notifications_queue"`
}

// TeamsProxyUpdateRequest represents a request to update a Teams proxy
type TeamsProxyUpdateRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Address              string  `json:"address"`
	Port                 int     `json:"port"`
	AzureTenant          string  `json:"azure_tenant"`
	EventhubID           *string `json:"eventhub_id"`
	MinNumberOfInstances int     `json:"min_number_of_instances"`
	NotificationsEnabled bool    `json:"notifications_enabled"`
	NotificationsQueue   *string `json:"notifications_queue"`
}

// TeamsProxyListResponse represents the response from listing Teams proxies
type TeamsProxyListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []TeamsProxy `json:"objects"`
}
