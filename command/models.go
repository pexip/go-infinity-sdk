package command

// CommandResponse represents a generic command response
type CommandResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Result  string `json:"result,omitempty"`
}

// Participant commands request types

// ParticipantDialRequest represents a request to dial out to a participant
type ParticipantDialRequest struct {
	ConferenceAlias     string `json:"conference_alias"`
	Destination         string `json:"destination"`
	CallType            string `json:"call_type,omitempty"` // "audio", "video", "video-only"
	CustomSIPHeaders    string `json:"custom_sip_headers,omitempty"`
	DTMFSequence        string `json:"dtmf_sequence,omitempty"`
	KeepConferenceAlive string `json:"keep_conference_alive,omitempty"` // "keep_conference_alive", "keep_conference_alive_if_multiple", "keep_conference_alive_never"
	LocalDisplayName    string `json:"local_display_name,omitempty"`
	Node                string `json:"node,omitempty"`
	PresentationURL     string `json:"presentation_url,omitempty"`
	Protocol            string `json:"protocol,omitempty"` // "gms", "h323", "sip", "mssip", "rtmp", "teams"
	RemoteDisplayName   string `json:"remote_display_name,omitempty"`
	Role                string `json:"role,omitempty"`    // "guest", "chair"
	Routing             string `json:"routing,omitempty"` // "manual", "routing_rule"
	Streaming           bool   `json:"streaming"`
	SystemLocation      string `json:"system_location,omitempty"`
}

// ParticipantDisconnectRequest represents a request to disconnect a participant
type ParticipantDisconnectRequest struct {
	ParticipantID string `json:"participant_id"`
}

// ParticipantMuteRequest represents a request to mute/unmute a participant
type ParticipantMuteRequest struct {
	ParticipantID string `json:"participant_id"`
}

// ParticipantUnmuteRequest represents a request to unmute a participant
type ParticipantUnmuteRequest struct {
	ParticipantID string `json:"participant_id"`
}

// ParticipantUnlockRequest represents a request to unlock a participant
type ParticipantUnlockRequest struct {
	ParticipantID string `json:"participant_id"`
}

// ParticipantRoleRequest represents a request to change a participant's role
type ParticipantRoleRequest struct {
	ParticipantID string `json:"participant_id"`
	Role          string `json:"role,omitempty"` // "chair" or "guest"
}

// ParticipantTransferRequest represents a request to transfer a participant
type ParticipantTransferRequest struct {
	ParticipantID   string `json:"participant_id"`
	ConferenceAlias string `json:"conference_alias"`
	Role            string `json:"role,omitempty"`
}

// Conference commands request types

// ConferenceLockRequest represents a request to lock a conference
type ConferenceLockRequest struct {
	ConferenceID string `json:"conference_id"`
}

// ConferenceUnlockRequest represents a request to unlock a conference
type ConferenceUnlockRequest struct {
	ConferenceID string `json:"conference_id"`
}

// ConferenceMuteGuestsRequest represents a request to mute all guests in a conference
type ConferenceMuteGuestsRequest struct {
	ConferenceID string `json:"conference_id"`
}

// ConferenceUnmuteGuestsRequest represents a request to unmute all guests in a conference
type ConferenceUnmuteGuestsRequest struct {
	ConferenceID string `json:"conference_id"`
}

// ConferenceTransformLayoutRequest represents a request to transform conference layout
type ConferenceTransformLayoutRequest struct {
	ConferenceID          string `json:"conference_id"`
	AIEnabledIndicator    *bool  `json:"ai_enabled_indicator,omitempty"`
	EnableOverlayText     *bool  `json:"enable_overlay_text,omitempty"`
	FreeFormOverlayText   string `json:"free_form_overlay_text,omitempty"`
	GuestLayout           string `json:"guest_layout,omitempty"`
	HostLayout            string `json:"host_layout,omitempty"`
	Layout                string `json:"layout,omitempty"`
	LiveCaptionsIndicator *bool  `json:"live_captions_indicator,omitempty"`
	PlusNPipEnabled       *bool  `json:"plus_n_pip_enabled,omitempty"`
	RecordingIndicator    *bool  `json:"recording_indicator,omitempty"`
	StreamingIndicator    *bool  `json:"streaming_indicator,omitempty"`
	TranscribingIndicator *bool  `json:"transcribing_indicator,omitempty"`
}

// ConferenceSendEmailRequest represents a request to send conference email
type ConferenceSendEmailRequest struct {
	ConferenceID             int  `json:"conference_id"`
	ConferenceSyncTemplateID *int `json:"conference_sync_template_id,omitempty"`
}

// System management commands request types

// BackupCreateRequest represents a request to create a backup
type BackupCreateRequest struct {
	Passphrase string `json:"passphrase"`
	Request    bool   `json:"request"`
}

// BackupRestoreRequest represents a request to restore from backup
type BackupRestoreRequest struct {
	Package    string `json:"package"`
	Passphrase string `json:"passphrase"`
}

// CertificatesImportRequest represents a request to import certificates
type CertificatesImportRequest struct {
	Bundle               string `json:"bundle"`
	PrivateKeyPassphrase string `json:"private_key_passphrase"`
}

// SnapshotRequest represents a request to create a snapshot
type SnapshotRequest struct {
	EndLimit                 *int  `json:"end_limit,omitempty"`
	IncludeDiagnosticMetrics *bool `json:"include_diagnostic_metrics,omitempty"`
	Limit                    *int  `json:"limit,omitempty"`
	Request                  *bool `json:"request,omitempty"`
}

// SoftwareBundleRequest represents a request to manage software bundle
type SoftwareBundleRequest struct {
	Package string `json:"package,omitempty"`
}

// UpgradeRequest represents a request to upgrade system
type UpgradeRequest struct {
	Package string `json:"package,omitempty"`
}

// StartCloudNodeRequest represents a request to start a cloud node
type StartCloudNodeRequest struct {
	InstanceID string `json:"instance_id"`
}

// SyncRequest represents a request to perform sync
type SyncRequest struct {
	ConferenceSyncTemplateID string `json:"conference_sync_template_id,omitempty"`
}

// DeviceSendEmailRequest represents a request to send device email
type DeviceSendEmailRequest struct {
	DeviceID                 int  `json:"device_id"`
	ConferenceSyncTemplateID *int `json:"conference_sync_template_id,omitempty"`
}

// Legacy request types for backward compatibility

// ParticipantDisconnectRequestLegacy represents the legacy disconnect request
type ParticipantDisconnectRequestLegacy struct {
	ParticipantUUID string `json:"participant_uuid"`
}

// ParticipantMuteRequestLegacy represents the legacy mute request
type ParticipantMuteRequestLegacy struct {
	ParticipantUUID string `json:"participant_uuid"`
	Setting         string `json:"setting"` // "mute", "unmute", or "toggle"
}

// ParticipantSpotlightRequest represents a request to spotlight a participant
type ParticipantSpotlightRequest struct {
	ParticipantUUID string `json:"participant_uuid"`
	Setting         string `json:"setting"` // "on", "off", or "toggle"
}

// ConferenceLockRequestLegacy represents the legacy conference lock request
type ConferenceLockRequestLegacy struct {
	ConferenceID int    `json:"conference_id"`
	Setting      string `json:"setting"` // "lock", "unlock", or "toggle"
}

// ParticipantMessageRequest represents a request to send a message to a participant
type ParticipantMessageRequest struct {
	ParticipantUUID string `json:"participant_uuid"`
	Message         string `json:"message"`
}

// ConferenceMessageRequest represents a request to send a message to all participants in a conference
type ConferenceMessageRequest struct {
	ConferenceID int    `json:"conference_id"`
	Message      string `json:"message"`
}

// ParticipantTransferRequestLegacy represents the legacy transfer request
type ParticipantTransferRequestLegacy struct {
	ParticipantUUID string `json:"participant_uuid"`
	ConferenceAlias string `json:"conference_alias"`
	Role            string `json:"role,omitempty"` // "chair" or "guest"
	PIN             string `json:"pin,omitempty"`
}

// ConferenceStartRequest represents a request to start a conference
type ConferenceStartRequest struct {
	ConferenceAlias string `json:"conference_alias"`
}

// ConferenceStopRequest represents a request to stop a conference
type ConferenceStopRequest struct {
	ConferenceID int `json:"conference_id"`
}

// ParticipantRoleRequestLegacy represents the legacy role change request
type ParticipantRoleRequestLegacy struct {
	ParticipantUUID string `json:"participant_uuid"`
	Role            string `json:"role"` // "chair" or "guest"
}

// TransferOptions contains options for transferring participants
type TransferOptions struct {
	Role string // "chair" or "guest"
	PIN  string
}
