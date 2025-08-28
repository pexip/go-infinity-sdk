/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// GlobalConfiguration represents the global system configuration (singleton resource)
type GlobalConfiguration struct {
	ID                           int      `json:"id,omitempty"`
	EnableWebRTC                 bool     `json:"enable_webrtc,omitempty"`
	EnableSIP                    bool     `json:"enable_sip,omitempty"`
	EnableH323                   bool     `json:"enable_h323,omitempty"`
	EnableRTMP                   bool     `json:"enable_rtmp,omitempty"`
	CryptoMode                   string   `json:"crypto_mode,omitempty"`
	MaxPixelsPerSecond           string   `json:"max_pixels_per_second,omitempty"`
	MediaPortsStart              int      `json:"media_ports_start,omitempty"`
	MediaPortsEnd                int      `json:"media_ports_end,omitempty"`
	SignallingPortsStart         int      `json:"signalling_ports_start,omitempty"`
	SignallingPortsEnd           int      `json:"signalling_ports_end,omitempty"`
	BurstingEnabled              bool     `json:"bursting_enabled,omitempty"`
	CloudProvider                string   `json:"cloud_provider,omitempty"`
	AWSAccessKey                 *string  `json:"aws_access_key,omitempty"`
	AWSSecretKey                 *string  `json:"aws_secret_key,omitempty"`
	AzureClientID                *string  `json:"azure_client_id,omitempty"`
	AzureSecret                  *string  `json:"azure_secret,omitempty"`
	GuestsOnlyTimeout            int      `json:"guests_only_timeout,omitempty"`
	WaitingForChairTimeout       int      `json:"waiting_for_chair_timeout,omitempty"`
	ConferenceCreatePermissions  string   `json:"conference_create_permissions,omitempty"`
	ConferenceCreationMode       string   `json:"conference_creation_mode,omitempty"`
	EnableAnalytics              bool     `json:"enable_analytics,omitempty"`
	EnableErrorReporting         bool     `json:"enable_error_reporting,omitempty"`
	BandwidthRestrictions        string   `json:"bandwidth_restrictions,omitempty"`
	AdministratorEmail           string   `json:"administrator_email,omitempty"`
	GlobalConferenceCreateGroups []string `json:"global_conference_create_groups,omitempty"`
	ResourceURI                  string   `json:"resource_uri,omitempty"`
}

// GlobalConfigurationUpdateRequest represents a request to update global configuration
type GlobalConfigurationUpdateRequest struct {
	EnableWebRTC                 *bool    `json:"enable_webrtc,omitempty"`
	EnableSIP                    *bool    `json:"enable_sip,omitempty"`
	EnableH323                   *bool    `json:"enable_h323,omitempty"`
	EnableRTMP                   *bool    `json:"enable_rtmp,omitempty"`
	CryptoMode                   string   `json:"crypto_mode,omitempty"`
	MaxPixelsPerSecond           string   `json:"max_pixels_per_second,omitempty"`
	MediaPortsStart              *int     `json:"media_ports_start,omitempty"`
	MediaPortsEnd                *int     `json:"media_ports_end,omitempty"`
	SignallingPortsStart         *int     `json:"signalling_ports_start,omitempty"`
	SignallingPortsEnd           *int     `json:"signalling_ports_end,omitempty"`
	BurstingEnabled              *bool    `json:"bursting_enabled,omitempty"`
	CloudProvider                string   `json:"cloud_provider,omitempty"`
	AWSAccessKey                 *string  `json:"aws_access_key,omitempty"`
	AWSSecretKey                 *string  `json:"aws_secret_key,omitempty"`
	AzureClientID                *string  `json:"azure_client_id,omitempty"`
	AzureSecret                  *string  `json:"azure_secret,omitempty"`
	GuestsOnlyTimeout            *int     `json:"guests_only_timeout,omitempty"`
	WaitingForChairTimeout       *int     `json:"waiting_for_chair_timeout,omitempty"`
	ConferenceCreatePermissions  string   `json:"conference_create_permissions,omitempty"`
	ConferenceCreationMode       string   `json:"conference_creation_mode,omitempty"`
	EnableAnalytics              *bool    `json:"enable_analytics,omitempty"`
	EnableErrorReporting         *bool    `json:"enable_error_reporting,omitempty"`
	BandwidthRestrictions        string   `json:"bandwidth_restrictions,omitempty"`
	AdministratorEmail           string   `json:"administrator_email,omitempty"`
	GlobalConferenceCreateGroups []string `json:"global_conference_create_groups,omitempty"`
}
