package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// ScheduledScaling represents a scheduled scaling policy configuration
type ScheduledScaling struct {
	ID                 int                `json:"id,omitempty"`
	PolicyName         string             `json:"policy_name"`
	PolicyType         string             `json:"policy_type"`
	ResourceIdentifier string             `json:"resource_identifier"`
	Enabled            bool               `json:"enabled"`
	InstancesToAdd     int                `json:"instances_to_add"`
	MinutesInAdvance   int                `json:"minutes_in_advance"`
	LocalTimezone      string             `json:"local_timezone"`
	StartDate          string             `json:"start_date"`
	TimeFrom           string             `json:"time_from"`
	TimeTo             string             `json:"time_to"`
	Mon                bool               `json:"mon"`
	Tue                bool               `json:"tue"`
	Wed                bool               `json:"wed"`
	Thu                bool               `json:"thu"`
	Fri                bool               `json:"fri"`
	Sat                bool               `json:"sat"`
	Sun                bool               `json:"sun"`
	Updated            *util.InfinityTime `json:"updated,omitempty"`
	ResourceURI        string             `json:"resource_uri,omitempty"`
}

// ScheduledScalingCreateRequest represents a request to create a scheduled scaling policy
type ScheduledScalingCreateRequest struct {
	PolicyName         string `json:"policy_name"`
	PolicyType         string `json:"policy_type"`
	ResourceIdentifier string `json:"resource_identifier"`
	Enabled            bool   `json:"enabled"`
	InstancesToAdd     int    `json:"instances_to_add"`
	MinutesInAdvance   int    `json:"minutes_in_advance"`
	LocalTimezone      string `json:"local_timezone"`
	StartDate          string `json:"start_date"`
	TimeFrom           string `json:"time_from"`
	TimeTo             string `json:"time_to"`
	Mon                bool   `json:"mon"`
	Tue                bool   `json:"tue"`
	Wed                bool   `json:"wed"`
	Thu                bool   `json:"thu"`
	Fri                bool   `json:"fri"`
	Sat                bool   `json:"sat"`
	Sun                bool   `json:"sun"`
}

// ScheduledScalingUpdateRequest represents a request to update a scheduled scaling policy
type ScheduledScalingUpdateRequest struct {
	PolicyName         string `json:"policy_name,omitempty"`
	PolicyType         string `json:"policy_type,omitempty"`
	ResourceIdentifier string `json:"resource_identifier,omitempty"`
	Enabled            *bool  `json:"enabled,omitempty"`
	InstancesToAdd     *int   `json:"instances_to_add,omitempty"`
	MinutesInAdvance   *int   `json:"minutes_in_advance,omitempty"`
	LocalTimezone      string `json:"local_timezone,omitempty"`
	StartDate          string `json:"start_date,omitempty"`
	TimeFrom           string `json:"time_from,omitempty"`
	TimeTo             string `json:"time_to,omitempty"`
	Mon                *bool  `json:"mon,omitempty"`
	Tue                *bool  `json:"tue,omitempty"`
	Wed                *bool  `json:"wed,omitempty"`
	Thu                *bool  `json:"thu,omitempty"`
	Fri                *bool  `json:"fri,omitempty"`
	Sat                *bool  `json:"sat,omitempty"`
	Sun                *bool  `json:"sun,omitempty"`
}

// ScheduledScalingListResponse represents the response from listing scheduled scaling policies
type ScheduledScalingListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ScheduledScaling `json:"objects"`
}
