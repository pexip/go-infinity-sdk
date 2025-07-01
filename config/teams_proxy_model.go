package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// TeamsProxy represents a Teams proxy configuration
type TeamsProxy struct {
	ID                   int                `json:"id,omitempty"`
	Name                 string             `json:"name"`
	Description          string             `json:"description,omitempty"`
	Address              string             `json:"address"`
	Port                 int                `json:"port"`
	AzureTenant          string             `json:"azure_tenant"`
	EventhubID           *string            `json:"eventhub_id,omitempty"`
	MinNumberOfInstances int                `json:"min_number_of_instances"`
	NotificationsEnabled bool               `json:"notifications_enabled"`
	NotificationsQueue   *string            `json:"notifications_queue,omitempty"`
	Updated              *util.InfinityTime `json:"updated,omitempty"`
	ResourceURI          string             `json:"resource_uri,omitempty"`
}

// TeamsProxyCreateRequest represents a request to create a Teams proxy
type TeamsProxyCreateRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description,omitempty"`
	Address              string  `json:"address"`
	Port                 int     `json:"port"`
	AzureTenant          string  `json:"azure_tenant"`
	EventhubID           *string `json:"eventhub_id,omitempty"`
	MinNumberOfInstances int     `json:"min_number_of_instances"`
	NotificationsEnabled bool    `json:"notifications_enabled"`
	NotificationsQueue   *string `json:"notifications_queue,omitempty"`
}

// TeamsProxyUpdateRequest represents a request to update a Teams proxy
type TeamsProxyUpdateRequest struct {
	Name                 string  `json:"name,omitempty"`
	Description          string  `json:"description,omitempty"`
	Address              string  `json:"address,omitempty"`
	Port                 *int    `json:"port,omitempty"`
	AzureTenant          string  `json:"azure_tenant,omitempty"`
	EventhubID           *string `json:"eventhub_id,omitempty"`
	MinNumberOfInstances *int    `json:"min_number_of_instances,omitempty"`
	NotificationsEnabled *bool   `json:"notifications_enabled,omitempty"`
	NotificationsQueue   *string `json:"notifications_queue,omitempty"`
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
