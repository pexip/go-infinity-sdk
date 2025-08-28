/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// AutomaticParticipant represents an automatic participant configuration
type AutomaticParticipant struct {
	ID                  int               `json:"id,omitempty"`
	Alias               string            `json:"alias"`
	Description         string            `json:"description,omitempty"`
	Conference          string            `json:"conference"`
	Protocol            string            `json:"protocol"`
	CallType            string            `json:"call_type"`
	Role                string            `json:"role"`
	DTMFSequence        string            `json:"dtmf_sequence,omitempty"`
	KeepConferenceAlive string            `json:"keep_conference_alive"`
	Routing             string            `json:"routing"`
	SystemLocation      *string           `json:"system_location,omitempty"`
	Streaming           bool              `json:"streaming"`
	RemoteDisplayName   string            `json:"remote_display_name,omitempty"`
	PresentationURL     string            `json:"presentation_url,omitempty"`
	CreationTime        util.InfinityTime `json:"creation_time,omitempty"`
	ResourceURI         string            `json:"resource_uri,omitempty"`
}

// AutomaticParticipantCreateRequest represents a request to create an automatic participant
type AutomaticParticipantCreateRequest struct {
	Alias               string  `json:"alias"`
	Description         string  `json:"description,omitempty"`
	Conference          string  `json:"conference"`
	Protocol            string  `json:"protocol"`
	CallType            string  `json:"call_type"`
	Role                string  `json:"role"`
	DTMFSequence        string  `json:"dtmf_sequence,omitempty"`
	KeepConferenceAlive string  `json:"keep_conference_alive"`
	Routing             string  `json:"routing"`
	SystemLocation      *string `json:"system_location,omitempty"`
	Streaming           bool    `json:"streaming"`
	RemoteDisplayName   string  `json:"remote_display_name,omitempty"`
	PresentationURL     string  `json:"presentation_url,omitempty"`
}

// AutomaticParticipantUpdateRequest represents a request to update an automatic participant
type AutomaticParticipantUpdateRequest struct {
	Alias               string  `json:"alias,omitempty"`
	Description         string  `json:"description,omitempty"`
	Conference          string  `json:"conference,omitempty"`
	Protocol            string  `json:"protocol,omitempty"`
	CallType            string  `json:"call_type,omitempty"`
	Role                string  `json:"role,omitempty"`
	DTMFSequence        string  `json:"dtmf_sequence,omitempty"`
	KeepConferenceAlive string  `json:"keep_conference_alive,omitempty"`
	Routing             string  `json:"routing,omitempty"`
	SystemLocation      *string `json:"system_location,omitempty"`
	Streaming           *bool   `json:"streaming,omitempty"`
	RemoteDisplayName   string  `json:"remote_display_name,omitempty"`
	PresentationURL     string  `json:"presentation_url,omitempty"`
}

// AutomaticParticipantListResponse represents the response from listing automatic participants
type AutomaticParticipantListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []AutomaticParticipant `json:"objects"`
}
