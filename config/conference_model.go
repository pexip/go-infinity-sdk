/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// Conference represents a conference configuration
type Conference struct {
	ID                              int                `json:"id,omitempty"`
	ResourceURI                     string             `json:"resource_uri,omitempty"`
	Name                            string             `json:"name"`
	Aliases                         *[]ConferenceAlias `json:"aliases,omitempty"`
	AllowGuests                     bool               `json:"allow_guests,omitempty"`
	AutomaticParticipants           []string           `json:"automatic_participants,omitempty"`
	BreakoutRooms                   bool               `json:"breakout_rooms,omitempty"`
	CallType                        string             `json:"call_type,omitempty"`
	CreationTime                    util.InfinityTime  `json:"creation_time,omitempty"`
	CryptoMode                      *string            `json:"crypto_mode,omitempty"`
	DenoiseEnabled                  bool               `json:"denoise_enabled,omitempty"`
	Description                     string             `json:"description,omitempty"`
	DirectMedia                     string             `json:"direct_media,omitempty"`
	DirectMediaNotificationDuration int                `json:"direct_media_notification_duration,omitempty"`
	EnableActiveSpeakerIndication   bool               `json:"enable_active_speaker_indication,omitempty"`
	EnableChat                      string             `json:"enable_chat,omitempty"`
	EnableOverlayText               bool               `json:"enable_overlay_text,omitempty"`
	ForcePresenterIntoMain          bool               `json:"force_presenter_into_main,omitempty"`
	GMSAccessToken                  *string            `json:"gms_access_token,omitempty"`
	GuestIdentityProviderGroup      *string            `json:"guest_identity_provider_group,omitempty"`
	GuestPIN                        string             `json:"guest_pin,omitempty"`
	GuestView                       *string            `json:"guest_view,omitempty"`
	GuestsCanPresent                bool               `json:"guests_can_present,omitempty"`
	GuestsCanSeeGuests              string             `json:"guests_can_see_guests,omitempty"`
	HostIdentityProviderGroup       *string            `json:"host_identity_provider_group,omitempty"`
	HostView                        *string            `json:"host_view,omitempty"`
	IVRTheme                        *string            `json:"ivr_theme,omitempty"`
	LiveCaptionsEnabled             string             `json:"live_captions_enabled,omitempty"`
	MatchString                     string             `json:"match_string,omitempty"`
	MaxCallRateIn                   *int               `json:"max_callrate_in,omitempty"`
	MaxCallRateOut                  *int               `json:"max_callrate_out,omitempty"`
	MaxPixelsPerSecond              *string            `json:"max_pixels_per_second,omitempty"`
	MediaPlaylist                   *string            `json:"media_playlist,omitempty"`
	MSSIPProxy                      *string            `json:"mssip_proxy,omitempty"`
	MuteAllGuests                   bool               `json:"mute_all_guests,omitempty"`
	NonIdpParticipants              string             `json:"non_idp_participants,omitempty"`
	OnCompletion                    *string            `json:"on_completion,omitempty"`
	ParticipantLimit                *int               `json:"participant_limit,omitempty"`
	PIN                             string             `json:"pin,omitempty"`
	PinningConfig                   *string            `json:"pinning_config,omitempty"`
	PostMatchString                 string             `json:"post_match_string,omitempty"`
	PostReplaceString               string             `json:"post_replace_string,omitempty"`
	PrimaryOwnerEmailAddress        string             `json:"primary_owner_email_address,omitempty"`
	ReplaceString                   string             `json:"replace_string,omitempty"`
	ScheduledConferences            *[]string          `json:"scheduled_conferences,omitempty"`
	ScheduledConferencesCount       int                `json:"scheduled_conferences_count,omitempty"`
	ServiceType                     string             `json:"service_type,omitempty"`
	SoftmuteEnabled                 bool               `json:"softmute_enabled,omitempty"`
	SyncTag                         string             `json:"sync_tag,omitempty"`
	SystemLocation                  *string            `json:"system_location,omitempty"`
	Tag                             string             `json:"tag,omitempty"`
	TeamsProxy                      *string            `json:"teams_proxy,omitempty"`
	TwoStageDialType                string             `json:"two_stage_dial_type,omitempty"`
}

// ConferenceCreateRequest represents a request to create a conference
type ConferenceCreateRequest struct {
	Name                            string    `json:"name"`
	Aliases                         *[]string `json:"aliases,omitempty"`
	AllowGuests                     bool      `json:"allow_guests,omitempty"`
	AutomaticParticipants           []string  `json:"automatic_participants,omitempty"`
	BreakoutRooms                   bool      `json:"breakout_rooms,omitempty"`
	CallType                        string    `json:"call_type,omitempty"`
	CryptoMode                      *string   `json:"crypto_mode,omitempty"`
	DenoiseEnabled                  bool      `json:"denoise_enabled,omitempty"`
	Description                     string    `json:"description,omitempty"`
	DirectMedia                     string    `json:"direct_media,omitempty"`
	DirectMediaNotificationDuration int       `json:"direct_media_notification_duration,omitempty"`
	EnableActiveSpeakerIndication   bool      `json:"enable_active_speaker_indication,omitempty"`
	EnableChat                      string    `json:"enable_chat,omitempty"`
	EnableOverlayText               bool      `json:"enable_overlay_text,omitempty"`
	ForcePresenterIntoMain          bool      `json:"force_presenter_into_main,omitempty"`
	GMSAccessToken                  *string   `json:"gms_access_token,omitempty"`
	GuestIdentityProviderGroup      *string   `json:"guest_identity_provider_group,omitempty"`
	GuestPIN                        string    `json:"guest_pin,omitempty"`
	GuestView                       *string   `json:"guest_view,omitempty"`
	GuestsCanPresent                bool      `json:"guests_can_present,omitempty"`
	GuestsCanSeeGuests              string    `json:"guests_can_see_guests,omitempty"`
	HostIdentityProviderGroup       *string   `json:"host_identity_provider_group,omitempty"`
	HostView                        *string   `json:"host_view,omitempty"`
	IVRTheme                        *string   `json:"ivr_theme,omitempty"`
	LiveCaptionsEnabled             string    `json:"live_captions_enabled,omitempty"`
	MatchString                     string    `json:"match_string,omitempty"`
	MaxCallRateIn                   *int      `json:"max_callrate_in,omitempty"`
	MaxCallRateOut                  *int      `json:"max_callrate_out,omitempty"`
	MaxPixelsPerSecond              *string   `json:"max_pixels_per_second,omitempty"`
	MediaPlaylist                   *string   `json:"media_playlist,omitempty"`
	MSSIPProxy                      *string   `json:"mssip_proxy,omitempty"`
	MuteAllGuests                   bool      `json:"mute_all_guests,omitempty"`
	NonIdpParticipants              string    `json:"non_idp_participants,omitempty"`
	OnCompletion                    *string   `json:"on_completion,omitempty"`
	ParticipantLimit                *int      `json:"participant_limit,omitempty"`
	PIN                             string    `json:"pin,omitempty"`
	PinningConfig                   *string   `json:"pinning_config,omitempty"`
	PostMatchString                 string    `json:"post_match_string,omitempty"`
	PostReplaceString               string    `json:"post_replace_string,omitempty"`
	PrimaryOwnerEmailAddress        string    `json:"primary_owner_email_address,omitempty"`
	ReplaceString                   string    `json:"replace_string,omitempty"`
	ScheduledConferences            *[]string `json:"scheduled_conferences,omitempty"`
	ScheduledConferencesCount       int       `json:"scheduled_conferences_count,omitempty"`
	ServiceType                     string    `json:"service_type,omitempty"`
	SoftmuteEnabled                 bool      `json:"softmute_enabled,omitempty"`
	SyncTag                         string    `json:"sync_tag,omitempty"`
	SystemLocation                  *string   `json:"system_location,omitempty"`
	Tag                             string    `json:"tag,omitempty"`
	TeamsProxy                      *string   `json:"teams_proxy,omitempty"`
	TwoStageDialType                string    `json:"two_stage_dial_type,omitempty"`
}

// ConferenceUpdateRequest represents a request to update a conference
type ConferenceUpdateRequest struct {
	Name                            string            `json:"name,omitempty"`
	Aliases                         *[]string         `json:"aliases"`
	AllowGuests                     bool              `json:"allow_guests"`
	AutomaticParticipants           []string          `json:"automatic_participants"`
	BreakoutRooms                   bool              `json:"breakout_rooms"`
	CallType                        string            `json:"call_type"`
	CryptoMode                      *string           `json:"crypto_mode"`
	DenoiseEnabled                  bool              `json:"denoise_enabled"`
	Description                     string            `json:"description"`
	DirectMedia                     string            `json:"direct_media"`
	DirectMediaNotificationDuration int               `json:"direct_media_notification_duration"`
	EnableActiveSpeakerIndication   bool              `json:"enable_active_speaker_indication"`
	EnableChat                      string            `json:"enable_chat"`
	EnableOverlayText               bool              `json:"enable_overlay_text"`
	ForcePresenterIntoMain          bool              `json:"force_presenter_into_main"`
	GMSAccessToken                  *string           `json:"gms_access_token"`
	GuestIdentityProviderGroup      *string           `json:"guest_identity_provider_group"`
	GuestPIN                        string            `json:"guest_pin"`
	GuestView                       *string           `json:"guest_view"`
	GuestsCanPresent                bool              `json:"guests_can_present"`
	GuestsCanSeeGuests              string            `json:"guests_can_see_guests"`
	HostIdentityProviderGroup       *string           `json:"host_identity_provider_group"`
	HostView                        *string           `json:"host_view"`
	IVRTheme                        *string           `json:"ivr_theme"`
	LiveCaptionsEnabled             string            `json:"live_captions_enabled"`
	MatchString                     string            `json:"match_string"`
	MaxCallRateIn                   *int              `json:"max_callrate_in"`
	MaxCallRateOut                  *int              `json:"max_callrate_out"`
	MaxPixelsPerSecond              *string           `json:"max_pixels_per_second"`
	MediaPlaylist                   *string           `json:"media_playlist"`
	MSSIPProxy                      *string           `json:"mssip_proxy"`
	MuteAllGuests                   bool              `json:"mute_all_guests"`
	NonIdpParticipants              string            `json:"non_idp_participants"`
	OnCompletion                    *string           `json:"on_completion"`
	ParticipantLimit                *int              `json:"participant_limit"`
	PIN                             string            `json:"pin"`
	PinningConfig                   *string           `json:"pinning_config"`
	PostMatchString                 string            `json:"post_match_string"`
	PostReplaceString               string            `json:"post_replace_string"`
	PrimaryOwnerEmailAddress        string            `json:"primary_owner_email_address"`
	ReplaceString                   string            `json:"replace_string"`
	ScheduledConferences            *[]string         `json:"scheduled_conferences"`
	ScheduledConferencesCount       int               `json:"scheduled_conferences_count"`
	ServiceType                     string            `json:"service_type"`
	SoftmuteEnabled                 bool              `json:"softmute_enabled"`
	SyncTag                         string            `json:"sync_tag"`
	SystemLocation                  *string           `json:"system_location"`
	Tag                             string            `json:"tag"`
	TeamsProxy                      *string           `json:"teams_proxy"`
	TwoStageDialType                string            `json:"two_stage_dial_type"`
}

// ConferenceListResponse represents the response from listing conferences
type ConferenceListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []Conference `json:"objects"`
}
