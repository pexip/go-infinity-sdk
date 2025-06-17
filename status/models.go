package status

import (
	"github.com/pexip/go-infinity-sdk/v38/util"
)

// Meta represents the pagination metadata for list responses
type Meta struct {
	Limit      int    `json:"limit"`
	Next       string `json:"next"`
	Offset     int    `json:"offset"`
	Previous   string `json:"previous"`
	TotalCount int    `json:"total_count"`
}

// SystemStatus represents the overall system status
type SystemStatus struct {
	Status      string            `json:"status"`
	Version     string            `json:"version"`
	Uptime      int               `json:"uptime"`
	Timestamp   util.InfinityTime `json:"timestamp"`
	HostName    string            `json:"hostname"`
	TotalMemory int64             `json:"total_memory"`
	UsedMemory  int64             `json:"used_memory"`
	CPULoad     float64           `json:"cpu_load"`
}

// ConferenceStatus represents a conference status
type ConferenceStatus struct {
	ID                   string             `json:"id"`
	Name                 string             `json:"name"`
	ServiceType          string             `json:"service_type"`
	IsStarted            bool               `json:"is_started"`
	IsLocked             bool               `json:"is_locked"`
	GuestsMuted          bool               `json:"guests_muted"`
	DirectMediaAvailable bool               `json:"direct_media_available"`
	StartTime            *util.InfinityTime `json:"start_time,omitempty"`
	Tag                  string             `json:"tag"`
	ResourceURI          string             `json:"resource_uri"`
}

// Participant represents a participant status
type Participant struct {
	Bandwidth               int               `json:"bandwidth"`
	CallDirection           string            `json:"call_direction"`
	CallQuality             string            `json:"call_quality"`
	CallTag                 string            `json:"call_tag"`
	CallUUID                string            `json:"call_uuid"`
	Conference              string            `json:"conference"`
	ConnectTime             util.InfinityTime `json:"connect_time"`
	ConversationID          string            `json:"conversation_id"`
	DestinationAlias        string            `json:"destination_alias"`
	DisplayName             string            `json:"display_name"`
	Encryption              string            `json:"encryption"`
	HasMedia                bool              `json:"has_media"`
	ID                      string            `json:"id"`
	IdpUUID                 string            `json:"idp_uuid"`
	IsClientMuted           bool              `json:"is_client_muted"`
	IsDirect                bool              `json:"is_direct"`
	IsDisconnectSupported   bool              `json:"is_disconnect_supported"`
	IsIdpAuthenticated      bool              `json:"is_idp_authenticated"`
	IsMuteSupported         bool              `json:"is_mute_supported"`
	IsMuted                 bool              `json:"is_muted"`
	IsOnHold                bool              `json:"is_on_hold"`
	IsPresentationSupported bool              `json:"is_presentation_supported"`
	IsPresenting            bool              `json:"is_presenting"`
	IsRecording             bool              `json:"is_recording"`
	IsStreaming             bool              `json:"is_streaming"`
	IsTranscribing          bool              `json:"is_transcribing"`
	IsTransferSupported     bool              `json:"is_transfer_supported"`
	LicenseCount            int               `json:"license_count"`
	LicenseType             string            `json:"license_type"`
	MediaNode               string            `json:"media_node"`
	ParentID                string            `json:"parent_id"`
	ParticipantAlias        string            `json:"participant_alias"`
	Protocol                string            `json:"protocol"`
	ProxyNode               string            `json:"proxy_node"`
	RemoteAddress           string            `json:"remote_address"`
	RemotePort              int               `json:"remote_port"`
	ResourceURI             string            `json:"resource_uri"`
	Role                    string            `json:"role"`
	RxBandwidth             int               `json:"rx_bandwidth"`
	ServiceTag              string            `json:"service_tag"`
	ServiceType             string            `json:"service_type"`
	SignallingNode          string            `json:"signalling_node"`
	SourceAlias             string            `json:"source_alias"`
	SystemLocation          string            `json:"system_location"`
	TranscodingEnabled      bool              `json:"transcoding_enabled"`
	TxBandwidth             int               `json:"tx_bandwidth"`
	Vendor                  string            `json:"vendor"`
}

// Alarm represents a system alarm
type Alarm struct {
	ID           int                `json:"id"`
	Identifier   int                `json:"identifier"`
	Level        string             `json:"level"`
	Name         string             `json:"name"`
	Details      string             `json:"details"`
	Instance     string             `json:"instance"`
	Node         string             `json:"node"`
	Acknowledged bool               `json:"acknowledged"`
	TimeRaised   *util.InfinityTime `json:"time_raised,omitempty"`
	ResourceURI  string             `json:"resource_uri"`
}

// Backplane represents a backplane connection status
type Backplane struct {
	ID                   string             `json:"id"`
	Conference           string             `json:"conference"`
	Type                 string             `json:"type"`
	Protocol             string             `json:"protocol"`
	ConnectTime          *util.InfinityTime `json:"connect_time,omitempty"`
	ServiceTag           string             `json:"service_tag"`
	SystemLocation       string             `json:"system_location"`
	MediaNode            string             `json:"media_node"`
	ProxyNode            string             `json:"proxy_node"`
	RemoteConferenceName string             `json:"remote_conference_name"`
	RemoteMediaNode      string             `json:"remote_media_node"`
	RemoteNodeName       string             `json:"remote_node_name"`
	ResourceURI          string             `json:"resource_uri"`
}

// BackupRequest represents a backup request status
type BackupRequest struct {
	ID          int                `json:"id"`
	Status      string             `json:"status"`
	Created     *util.InfinityTime `json:"created,omitempty"`
	Started     *util.InfinityTime `json:"started,omitempty"`
	Completed   *util.InfinityTime `json:"completed,omitempty"`
	Size        int64              `json:"size"`
	Description string             `json:"description"`
	ResourceURI string             `json:"resource_uri"`
}

// CloudMonitoredLocation represents a cloud monitored location status
type CloudMonitoredLocation struct {
	FreeHDCalls      int    `json:"free_hd_calls"`
	ID               int    `json:"id"`
	MaxHDCalls       int    `json:"max_hd_calls"`
	MediaLoad        int    `json:"media_load"`
	Name             string `json:"name"`
	OverflowLocation string `json:"overflow_location"`
	ResourceURI      string `json:"resource_uri"`
}

// CloudNode represents a cloud node status
type CloudNode struct {
	ID                 string             `json:"id"`
	Name               string             `json:"name"`
	Status             string             `json:"status"`
	InstanceType       string             `json:"instance_type"`
	Region             string             `json:"region"`
	LaunchTime         *util.InfinityTime `json:"launch_time,omitempty"`
	LastContactTime    *util.InfinityTime `json:"last_contact_time,omitempty"`
	CPU                float64            `json:"cpu"`
	Memory             float64            `json:"memory"`
	ActiveConferences  int                `json:"active_conferences"`
	ActiveParticipants int                `json:"active_participants"`
	ResourceURI        string             `json:"resource_uri"`
}

// CloudOverflowLocation represents a cloud overflow location status
type CloudOverflowLocation struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Region        string `json:"region"`
	Status        string `json:"status"`
	InstanceCount int    `json:"instance_count"`
	ResourceURI   string `json:"resource_uri"`
}

// ConferenceShard represents a conference shard status
type ConferenceShard struct {
	ID               string `json:"id"`
	ConferenceName   string `json:"conference_name"`
	ShardNumber      int    `json:"shard_number"`
	NodeID           string `json:"node_id"`
	ParticipantCount int    `json:"participant_count"`
	ResourceURI      string `json:"resource_uri"`
}

// ConferenceSync represents a conference sync status
type ConferenceSync struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Status       string             `json:"status"`
	LastSync     *util.InfinityTime `json:"last_sync,omitempty"`
	SyncInterval int                `json:"sync_interval"`
	ErrorMessage string             `json:"error_message"`
	ResourceURI  string             `json:"resource_uri"`
}

// ExchangeScheduler represents an Exchange scheduler status
type ExchangeScheduler struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	Status            string             `json:"status"`
	LastSync          *util.InfinityTime `json:"last_sync,omitempty"`
	NextSync          *util.InfinityTime `json:"next_sync,omitempty"`
	ProcessedMeetings int                `json:"processed_meetings"`
	ErrorCount        int                `json:"error_count"`
	ResourceURI       string             `json:"resource_uri"`
}

// Licensing represents licensing status
type Licensing struct {
	AudioCount          int  `json:"audio_count"`
	AudioTotal          int  `json:"audio_total"`
	PortCount           int  `json:"port_count"`
	PortTotal           int  `json:"port_total"`
	SystemCount         int  `json:"system_count"`
	SystemTotal         int  `json:"system_total"`
	VMRCount            int  `json:"vmr_count"`
	VMRTotal            int  `json:"vmr_total"`
	TeamsCount          int  `json:"teams_count"`
	TeamsTotal          int  `json:"teams_total"`
	GHMCount            int  `json:"ghm_count"`
	GHMTotal            int  `json:"ghm_total"`
	OTJCount            int  `json:"otj_count"`
	OTJTotal            int  `json:"otj_total"`
	SchedulingCount     int  `json:"scheduling_count"`
	SchedulingTotal     int  `json:"scheduling_total"`
	TelehealthCount     int  `json:"telehealth_count"`
	TelehealthTotal     int  `json:"telehealth_total"`
	CustomLayoutsActive bool `json:"customlayouts_active"`
}

// ManagementVM represents a management VM status
type ManagementVM struct {
	ID                   int                `json:"id"`
	ConfigurationID      int                `json:"configuration_id"`
	Name                 string             `json:"name"`
	Primary              bool               `json:"primary"`
	SyncStatus           string             `json:"sync_status"`
	UpgradeStatus        string             `json:"upgrade_status"`
	Version              string             `json:"version"`
	LastAttemptedContact *util.InfinityTime `json:"last_attempted_contact,omitempty"`
	LastUpdated          *util.InfinityTime `json:"last_updated,omitempty"`
	ResourceURI          string             `json:"resource_uri"`
}

// MJXEndpoint represents an MJX endpoint status
type MJXEndpoint struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	Status            string             `json:"status"`
	EndpointType      string             `json:"endpoint_type"`
	LastContact       *util.InfinityTime `json:"last_contact,omitempty"`
	Version           string             `json:"version"`
	ActiveConnections int                `json:"active_connections"`
	ResourceURI       string             `json:"resource_uri"`
}

// MJXMeeting represents an MJX meeting status
type MJXMeeting struct {
	ID               string             `json:"id"`
	Subject          string             `json:"subject"`
	Organizer        string             `json:"organizer"`
	StartTime        *util.InfinityTime `json:"start_time,omitempty"`
	EndTime          *util.InfinityTime `json:"end_time,omitempty"`
	Status           string             `json:"status"`
	ParticipantCount int                `json:"participant_count"`
	ConferenceAlias  string             `json:"conference_alias"`
	ResourceURI      string             `json:"resource_uri"`
}

// RegistrationAlias represents a registration alias status
type RegistrationAlias struct {
	ID          int    `json:"id"`
	Alias       string `json:"alias"`
	Status      string `json:"status"`
	NodeID      string `json:"node_id"`
	Protocol    string `json:"protocol"`
	ResourceURI string `json:"resource_uri"`
}

// SchedulingOperation represents a scheduling operation status
type SchedulingOperation struct {
	ID             int                `json:"id"`
	OperationType  string             `json:"operation_type"`
	Status         string             `json:"status"`
	CreatedTime    *util.InfinityTime `json:"created_time,omitempty"`
	CompletedTime  *util.InfinityTime `json:"completed_time,omitempty"`
	ErrorMessage   string             `json:"error_message"`
	ConferenceName string             `json:"conference_name"`
	ResourceURI    string             `json:"resource_uri"`
}

// SnapshotRequest represents a snapshot request status
type SnapshotRequest struct {
	ID          int                `json:"id"`
	Status      string             `json:"status"`
	Created     *util.InfinityTime `json:"created,omitempty"`
	Started     *util.InfinityTime `json:"started,omitempty"`
	Completed   *util.InfinityTime `json:"completed,omitempty"`
	Size        int64              `json:"size"`
	Description string             `json:"description"`
	ResourceURI string             `json:"resource_uri"`
}

// SystemLocation represents a system location status
type SystemLocation struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ActiveNodes int    `json:"active_nodes"`
	TotalNodes  int    `json:"total_nodes"`
	ActiveCalls int    `json:"active_calls"`
	ResourceURI string `json:"resource_uri"`
}

// TeamsNode represents a Teams node status
type TeamsNode struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Status      string             `json:"status"`
	Version     string             `json:"version"`
	LastContact *util.InfinityTime `json:"last_contact,omitempty"`
	ActiveCalls int                `json:"active_calls"`
	ResourceURI string             `json:"resource_uri"`
}

// TeamsNodeCall represents a Teams node call status
type TeamsNodeCall struct {
	ID              string             `json:"id"`
	TeamsNodeID     int                `json:"teams_node_id"`
	ConferenceName  string             `json:"conference_name"`
	ParticipantName string             `json:"participant_name"`
	CallDirection   string             `json:"call_direction"`
	StartTime       *util.InfinityTime `json:"start_time,omitempty"`
	Duration        int                `json:"duration"`
	Status          string             `json:"status"`
	ResourceURI     string             `json:"resource_uri"`
}

type WorkerVM struct {
	BootTime              string            `json:"boot_time"`
	ConfigurationID       int               `json:"configuration_id"`
	CPUCapabilities       string            `json:"cpu_capabilities"`
	CPUCount              int               `json:"cpu_count"`
	CPUModel              string            `json:"cpu_model"`
	DeployStatus          string            `json:"deploy_status"`
	Hypervisor            string            `json:"hypervisor"`
	ID                    int               `json:"id"`
	LastAttemptedContact  util.InfinityTime `json:"last_attempted_contact"`
	LastReported          util.InfinityTime `json:"last_reported"`
	LastUpdated           util.InfinityTime `json:"last_updated"`
	MaintenanceMode       bool              `json:"maintenance_mode"`
	MaintenanceModeReason string            `json:"maintenance_mode_reason"`
	MaxAudioCalls         int               `json:"max_audio_calls"`
	MaxDirectParticipants int               `json:"max_direct_participants"`
	MaxFullHDCalls        int               `json:"max_full_hd_calls"`
	MaxHDCalls            int               `json:"max_hd_calls"`
	MaxMediaTokens        int               `json:"max_media_tokens"`
	MaxSDCalls            int               `json:"max_sd_calls"`
	MediaLoad             int               `json:"media_load"`
	MediaTokensUsed       int               `json:"media_tokens_used"`
	Name                  string            `json:"name"`
	NodeType              string            `json:"node_type"`
	ResourceURI           string            `json:"resource_uri"`
	SignalingCount        int               `json:"signaling_count"`
	SyncStatus            string            `json:"sync_status"`
	SystemLocation        string            `json:"system_location"`
	TotalRAM              int               `json:"total_ram"`
	UpgradeStatus         string            `json:"upgrade_status"`
	Version               string            `json:"version"`
}

type ConferenceListResponse struct {
	Meta    Meta               `json:"meta"`
	Objects []ConferenceStatus `json:"objects"`
}

type ParticipantListResponse struct {
	Meta    Meta          `json:"meta"`
	Objects []Participant `json:"objects"`
}

type WorkerVMListResponse struct {
	Meta    Meta       `json:"meta"`
	Objects []WorkerVM `json:"objects"`
}

type AlarmListResponse struct {
	Meta    Meta    `json:"meta"`
	Objects []Alarm `json:"objects"`
}

type BackplaneListResponse struct {
	Meta    Meta        `json:"meta"`
	Objects []Backplane `json:"objects"`
}

type BackupRequestListResponse struct {
	Meta    Meta            `json:"meta"`
	Objects []BackupRequest `json:"objects"`
}

type CloudMonitoredLocationListResponse struct {
	Meta    Meta                     `json:"meta"`
	Objects []CloudMonitoredLocation `json:"objects"`
}

type CloudNodeListResponse struct {
	Meta    Meta        `json:"meta"`
	Objects []CloudNode `json:"objects"`
}

type CloudOverflowLocationListResponse struct {
	Meta    Meta                    `json:"meta"`
	Objects []CloudOverflowLocation `json:"objects"`
}

type ConferenceShardListResponse struct {
	Meta    Meta              `json:"meta"`
	Objects []ConferenceShard `json:"objects"`
}

type ConferenceSyncListResponse struct {
	Meta    Meta             `json:"meta"`
	Objects []ConferenceSync `json:"objects"`
}

type ExchangeSchedulerListResponse struct {
	Meta    Meta                `json:"meta"`
	Objects []ExchangeScheduler `json:"objects"`
}

type ManagementVMListResponse struct {
	Meta    Meta           `json:"meta"`
	Objects []ManagementVM `json:"objects"`
}

type MJXEndpointListResponse struct {
	Meta    Meta          `json:"meta"`
	Objects []MJXEndpoint `json:"objects"`
}

type MJXMeetingListResponse struct {
	Meta    Meta         `json:"meta"`
	Objects []MJXMeeting `json:"objects"`
}

type RegistrationAliasListResponse struct {
	Meta    Meta                `json:"meta"`
	Objects []RegistrationAlias `json:"objects"`
}

type SchedulingOperationListResponse struct {
	Meta    Meta                  `json:"meta"`
	Objects []SchedulingOperation `json:"objects"`
}

type SnapshotRequestListResponse struct {
	Meta    Meta              `json:"meta"`
	Objects []SnapshotRequest `json:"objects"`
}

type SystemLocationListResponse struct {
	Meta    Meta             `json:"meta"`
	Objects []SystemLocation `json:"objects"`
}

type TeamsNodeListResponse struct {
	Meta    Meta        `json:"meta"`
	Objects []TeamsNode `json:"objects"`
}

type TeamsNodeCallListResponse struct {
	Meta    Meta            `json:"meta"`
	Objects []TeamsNodeCall `json:"objects"`
}
