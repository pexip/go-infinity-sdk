package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// ScheduledAlias represents a scheduled alias configuration
type ScheduledAlias struct {
	ID                     int                `json:"id,omitempty"`
	Alias                  string             `json:"alias"`
	AliasNumber            int                `json:"alias_number"`
	NumericAlias           string             `json:"numeric_alias"`
	UUID                   string             `json:"uuid"`
	ExchangeConnector      string             `json:"exchange_connector"`
	IsUsed                 bool               `json:"is_used"`
	EWSItemUID             *string            `json:"ews_item_uid,omitempty"`
	CreationTime           util.InfinityTime  `json:"creation_time,omitempty"`
	ConferenceDeletionTime *util.InfinityTime `json:"conference_deletion_time,omitempty"`
	ResourceURI            string             `json:"resource_uri,omitempty"`
}

// ScheduledAliasCreateRequest represents a request to create a scheduled alias
type ScheduledAliasCreateRequest struct {
	Alias             string  `json:"alias"`
	AliasNumber       int     `json:"alias_number"`
	NumericAlias      string  `json:"numeric_alias"`
	UUID              string  `json:"uuid"`
	ExchangeConnector string  `json:"exchange_connector"`
	IsUsed            bool    `json:"is_used"`
	EWSItemUID        *string `json:"ews_item_uid,omitempty"`
}

// ScheduledAliasUpdateRequest represents a request to update a scheduled alias
type ScheduledAliasUpdateRequest struct {
	Alias             string  `json:"alias,omitempty"`
	AliasNumber       *int    `json:"alias_number,omitempty"`
	NumericAlias      string  `json:"numeric_alias,omitempty"`
	UUID              string  `json:"uuid,omitempty"`
	ExchangeConnector string  `json:"exchange_connector,omitempty"`
	IsUsed            *bool   `json:"is_used,omitempty"`
	EWSItemUID        *string `json:"ews_item_uid,omitempty"`
}

// ScheduledAliasListResponse represents the response from listing scheduled aliases
type ScheduledAliasListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ScheduledAlias `json:"objects"`
}
