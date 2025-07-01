package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// ScheduledConference represents a scheduled conference configuration
type ScheduledConference struct {
	ID                  int               `json:"id,omitempty"`
	Conference          string            `json:"conference"`
	StartTime           util.InfinityTime `json:"start_time"`
	EndTime             util.InfinityTime `json:"end_time"`
	Subject             string            `json:"subject,omitempty"`
	EWSItemID           string            `json:"ews_item_id"`
	EWSItemUID          string            `json:"ews_item_uid,omitempty"`
	RecurringConference *string           `json:"recurring_conference,omitempty"`
	ScheduledAlias      *string           `json:"scheduled_alias,omitempty"`
	ResourceURI         string            `json:"resource_uri,omitempty"`
}

// ScheduledConferenceCreateRequest represents a request to create a scheduled conference
type ScheduledConferenceCreateRequest struct {
	Conference          string            `json:"conference"`
	StartTime           util.InfinityTime `json:"start_time"`
	EndTime             util.InfinityTime `json:"end_time"`
	Subject             string            `json:"subject,omitempty"`
	EWSItemID           string            `json:"ews_item_id"`
	EWSItemUID          string            `json:"ews_item_uid,omitempty"`
	RecurringConference *string           `json:"recurring_conference,omitempty"`
	ScheduledAlias      *string           `json:"scheduled_alias,omitempty"`
}

// ScheduledConferenceUpdateRequest represents a request to update a scheduled conference
type ScheduledConferenceUpdateRequest struct {
	Conference          string             `json:"conference,omitempty"`
	StartTime           *util.InfinityTime `json:"start_time,omitempty"`
	EndTime             *util.InfinityTime `json:"end_time,omitempty"`
	Subject             string             `json:"subject,omitempty"`
	EWSItemID           string             `json:"ews_item_id,omitempty"`
	EWSItemUID          string             `json:"ews_item_uid,omitempty"`
	RecurringConference *string            `json:"recurring_conference,omitempty"`
	ScheduledAlias      *string            `json:"scheduled_alias,omitempty"`
}

// ScheduledConferenceListResponse represents the response from listing scheduled conferences
type ScheduledConferenceListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ScheduledConference `json:"objects"`
}
