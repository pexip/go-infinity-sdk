package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// GatewayRoutingRule represents a gateway routing rule configuration
type GatewayRoutingRule struct {
	ID                              int               `json:"id,omitempty"`
	Name                            string            `json:"name"`
	Description                     string            `json:"description,omitempty"`
	Priority                        int               `json:"priority"`
	Enable                          bool              `json:"enable"`
	MatchString                     string            `json:"match_string"`
	MatchStringFull                 bool              `json:"match_string_full"`
	ReplaceString                   string            `json:"replace_string,omitempty"`
	CalledDeviceType                string            `json:"called_device_type"`
	OutgoingProtocol                string            `json:"outgoing_protocol"`
	CallType                        string            `json:"call_type"`
	MatchIncomingCalls              bool              `json:"match_incoming_calls"`
	MatchOutgoingCalls              bool              `json:"match_outgoing_calls"`
	MatchIncomingSIP                bool              `json:"match_incoming_sip"`
	MatchIncomingH323               bool              `json:"match_incoming_h323"`
	MatchIncomingMSSIP              bool              `json:"match_incoming_mssip"`
	MatchIncomingWebRTC             bool              `json:"match_incoming_webrtc"`
	MatchIncomingTeams              bool              `json:"match_incoming_teams"`
	MatchIncomingOnlyIfRegistered   bool              `json:"match_incoming_only_if_registered"`
	MatchSourceLocation             *string           `json:"match_source_location,omitempty"`
	OutgoingLocation                *string           `json:"outgoing_location,omitempty"`
	SIPProxy                        *string           `json:"sip_proxy,omitempty"`
	H323Gatekeeper                  *string           `json:"h323_gatekeeper,omitempty"`
	MSSIPProxy                      *string           `json:"mssip_proxy,omitempty"`
	TeamsProxy                      *string           `json:"teams_proxy,omitempty"`
	STUNServer                      *string           `json:"stun_server,omitempty"`
	TURNServer                      *string           `json:"turn_server,omitempty"`
	GMSAccessToken                  *string           `json:"gms_access_token,omitempty"`
	TelehealthProfile               *string           `json:"telehealth_profile,omitempty"`
	IVRTheme                        *string           `json:"ivr_theme,omitempty"`
	MaxPixelsPerSecond              *string           `json:"max_pixels_per_second,omitempty"`
	MaxCallrateIn                   *int              `json:"max_callrate_in,omitempty"`
	MaxCallrateOut                  *int              `json:"max_callrate_out,omitempty"`
	CryptoMode                      *string           `json:"crypto_mode,omitempty"`
	DenoiseAudio                    bool              `json:"denoise_audio"`
	LiveCaptionsEnabled             string            `json:"live_captions_enabled"`
	ExternalParticipantAvatarLookup *string           `json:"external_participant_avatar_lookup,omitempty"`
	TreatAsTrusted                  bool              `json:"treat_as_trusted"`
	Tag                             string            `json:"tag,omitempty"`
	DisabledCodecs                  []string          `json:"disabled_codecs,omitempty"`
	CreationTime                    util.InfinityTime `json:"creation_time,omitempty"`
	ResourceURI                     string            `json:"resource_uri,omitempty"`
}

// GatewayRoutingRuleCreateRequest represents a request to create a gateway routing rule
type GatewayRoutingRuleCreateRequest struct {
	Name                            string   `json:"name"`
	Description                     string   `json:"description,omitempty"`
	Priority                        int      `json:"priority"`
	Enable                          bool     `json:"enable"`
	MatchString                     string   `json:"match_string"`
	MatchStringFull                 bool     `json:"match_string_full"`
	ReplaceString                   string   `json:"replace_string,omitempty"`
	CalledDeviceType                string   `json:"called_device_type"`
	OutgoingProtocol                string   `json:"outgoing_protocol"`
	CallType                        string   `json:"call_type"`
	MatchIncomingCalls              bool     `json:"match_incoming_calls"`
	MatchOutgoingCalls              bool     `json:"match_outgoing_calls"`
	MatchIncomingSIP                bool     `json:"match_incoming_sip"`
	MatchIncomingH323               bool     `json:"match_incoming_h323"`
	MatchIncomingMSSIP              bool     `json:"match_incoming_mssip"`
	MatchIncomingWebRTC             bool     `json:"match_incoming_webrtc"`
	MatchIncomingTeams              bool     `json:"match_incoming_teams"`
	MatchIncomingOnlyIfRegistered   bool     `json:"match_incoming_only_if_registered"`
	MatchSourceLocation             *string  `json:"match_source_location,omitempty"`
	OutgoingLocation                *string  `json:"outgoing_location,omitempty"`
	SIPProxy                        *string  `json:"sip_proxy,omitempty"`
	H323Gatekeeper                  *string  `json:"h323_gatekeeper,omitempty"`
	MSSIPProxy                      *string  `json:"mssip_proxy,omitempty"`
	TeamsProxy                      *string  `json:"teams_proxy,omitempty"`
	STUNServer                      *string  `json:"stun_server,omitempty"`
	TURNServer                      *string  `json:"turn_server,omitempty"`
	GMSAccessToken                  *string  `json:"gms_access_token,omitempty"`
	TelehealthProfile               *string  `json:"telehealth_profile,omitempty"`
	IVRTheme                        *string  `json:"ivr_theme,omitempty"`
	MaxPixelsPerSecond              *string  `json:"max_pixels_per_second,omitempty"`
	MaxCallrateIn                   *int     `json:"max_callrate_in,omitempty"`
	MaxCallrateOut                  *int     `json:"max_callrate_out,omitempty"`
	CryptoMode                      *string  `json:"crypto_mode,omitempty"`
	DenoiseAudio                    bool     `json:"denoise_audio"`
	LiveCaptionsEnabled             string   `json:"live_captions_enabled"`
	ExternalParticipantAvatarLookup *string  `json:"external_participant_avatar_lookup,omitempty"`
	TreatAsTrusted                  bool     `json:"treat_as_trusted"`
	Tag                             string   `json:"tag,omitempty"`
	DisabledCodecs                  []string `json:"disabled_codecs,omitempty"`
}

// GatewayRoutingRuleUpdateRequest represents a request to update a gateway routing rule
type GatewayRoutingRuleUpdateRequest struct {
	Name                            string   `json:"name,omitempty"`
	Description                     string   `json:"description,omitempty"`
	Priority                        *int     `json:"priority,omitempty"`
	Enable                          *bool    `json:"enable,omitempty"`
	MatchString                     string   `json:"match_string,omitempty"`
	MatchStringFull                 *bool    `json:"match_string_full,omitempty"`
	ReplaceString                   string   `json:"replace_string,omitempty"`
	CalledDeviceType                string   `json:"called_device_type,omitempty"`
	OutgoingProtocol                string   `json:"outgoing_protocol,omitempty"`
	CallType                        string   `json:"call_type,omitempty"`
	MatchIncomingCalls              *bool    `json:"match_incoming_calls,omitempty"`
	MatchOutgoingCalls              *bool    `json:"match_outgoing_calls,omitempty"`
	MatchIncomingSIP                *bool    `json:"match_incoming_sip,omitempty"`
	MatchIncomingH323               *bool    `json:"match_incoming_h323,omitempty"`
	MatchIncomingMSSIP              *bool    `json:"match_incoming_mssip,omitempty"`
	MatchIncomingWebRTC             *bool    `json:"match_incoming_webrtc,omitempty"`
	MatchIncomingTeams              *bool    `json:"match_incoming_teams,omitempty"`
	MatchIncomingOnlyIfRegistered   *bool    `json:"match_incoming_only_if_registered,omitempty"`
	MatchSourceLocation             *string  `json:"match_source_location,omitempty"`
	OutgoingLocation                *string  `json:"outgoing_location,omitempty"`
	SIPProxy                        *string  `json:"sip_proxy,omitempty"`
	H323Gatekeeper                  *string  `json:"h323_gatekeeper,omitempty"`
	MSSIPProxy                      *string  `json:"mssip_proxy,omitempty"`
	TeamsProxy                      *string  `json:"teams_proxy,omitempty"`
	STUNServer                      *string  `json:"stun_server,omitempty"`
	TURNServer                      *string  `json:"turn_server,omitempty"`
	GMSAccessToken                  *string  `json:"gms_access_token,omitempty"`
	TelehealthProfile               *string  `json:"telehealth_profile,omitempty"`
	IVRTheme                        *string  `json:"ivr_theme,omitempty"`
	MaxPixelsPerSecond              *string  `json:"max_pixels_per_second,omitempty"`
	MaxCallrateIn                   *int     `json:"max_callrate_in,omitempty"`
	MaxCallrateOut                  *int     `json:"max_callrate_out,omitempty"`
	CryptoMode                      *string  `json:"crypto_mode,omitempty"`
	DenoiseAudio                    *bool    `json:"denoise_audio,omitempty"`
	LiveCaptionsEnabled             string   `json:"live_captions_enabled,omitempty"`
	ExternalParticipantAvatarLookup *string  `json:"external_participant_avatar_lookup,omitempty"`
	TreatAsTrusted                  *bool    `json:"treat_as_trusted,omitempty"`
	Tag                             string   `json:"tag,omitempty"`
	DisabledCodecs                  []string `json:"disabled_codecs,omitempty"`
}

// GatewayRoutingRuleListResponse represents the response from listing gateway routing rules
type GatewayRoutingRuleListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []GatewayRoutingRule `json:"objects"`
}
