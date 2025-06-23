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
	Bandwidth               *int               `json:"bandwidth,omitempty"`
	CallDirection           string             `json:"call_direction"`
	CallQuality             string             `json:"call_quality"`
	CallTag                 string             `json:"call_tag"`
	CallUUID                string             `json:"call_uuid"`
	Conference              string             `json:"conference"`
	ConnectTime             *util.InfinityTime `json:"connect_time,omitempty"`
	ConversationID          string             `json:"conversation_id"`
	DestinationAlias        string             `json:"destination_alias"`
	DisplayName             string             `json:"display_name"`
	Encryption              string             `json:"encryption"`
	HasMedia                bool               `json:"has_media"`
	ID                      string             `json:"id"`
	IdpUUID                 string             `json:"idp_uuid"`
	IsClientMuted           bool               `json:"is_client_muted"`
	IsDirect                bool               `json:"is_direct"`
	IsDisconnectSupported   bool               `json:"is_disconnect_supported"`
	IsIdpAuthenticated      bool               `json:"is_idp_authenticated"`
	IsMuteSupported         bool               `json:"is_mute_supported"`
	IsMuted                 bool               `json:"is_muted"`
	IsOnHold                bool               `json:"is_on_hold"`
	IsPresentationSupported bool               `json:"is_presentation_supported"`
	IsPresenting            bool               `json:"is_presenting"`
	IsRecording             bool               `json:"is_recording"`
	IsStreaming             bool               `json:"is_streaming"`
	IsTranscribing          bool               `json:"is_transcribing"`
	IsTransferSupported     bool               `json:"is_transfer_supported"`
	LicenseCount            int                `json:"license_count"`
	LicenseType             string             `json:"license_type"`
	MediaNode               string             `json:"media_node"`
	ParentID                string             `json:"parent_id"`
	ParticipantAlias        string             `json:"participant_alias"`
	Protocol                string             `json:"protocol"`
	ProxyNode               string             `json:"proxy_node"`
	RemoteAddress           string             `json:"remote_address"`
	RemotePort              int                `json:"remote_port"`
	ResourceURI             string             `json:"resource_uri"`
	Role                    string             `json:"role"`
	RxBandwidth             *int               `json:"rx_bandwidth,omitempty"`
	ServiceTag              string             `json:"service_tag"`
	ServiceType             string             `json:"service_type"`
	SignallingNode          string             `json:"signalling_node"`
	SourceAlias             string             `json:"source_alias"`
	SystemLocation          string             `json:"system_location"`
	TranscodingEnabled      bool               `json:"transcoding_enabled"`
	TxBandwidth             *int               `json:"tx_bandwidth,omitempty"`
	Vendor                  string             `json:"vendor"`
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
	CreatedAt   *util.InfinityTime `json:"created_at,omitempty"`
	UpdatedAt   *util.InfinityTime `json:"updated_at,omitempty"`
	DownloadURI string             `json:"download_uri"`
	Message     string             `json:"message"`
	ResourceURI string             `json:"resource_uri"`
	State       string             `json:"state"`
}

// CloudMonitoredLocation represents a cloud monitored location status
type CloudMonitoredLocation struct {
	FreeHDCalls      *int    `json:"free_hd_calls,omitempty"`
	ID               int     `json:"id"`
	MaxHDCalls       *int    `json:"max_hd_calls,omitempty"`
	MediaLoad        *int    `json:"media_load,omitempty"`
	Name             *string `json:"name,omitempty"`
	OverflowLocation string  `json:"overflow_location"`
	ResourceURI      string  `json:"resource_uri"`
}

// CloudNode represents a cloud node status
type CloudNode struct {
	AWSInstanceID                     string             `json:"aws_instance_id"`
	AWSInstanceIP                     string             `json:"aws_instance_ip"`
	AWSInstanceLaunchTime             *util.InfinityTime `json:"aws_instance_launch_time,omitempty"`
	AWSInstanceName                   string             `json:"aws_instance_name"`
	AWSInstanceState                  string             `json:"aws_instance_state"`
	CloudInstanceID                   string             `json:"cloud_instance_id"`
	CloudInstanceIP                   string             `json:"cloud_instance_ip"`
	CloudInstanceLaunchTime           *util.InfinityTime `json:"cloud_instance_launch_time,omitempty"`
	CloudInstanceName                 string             `json:"cloud_instance_name"`
	CloudInstanceState                string             `json:"cloud_instance_state"`
	MaxHDCalls                        *int               `json:"max_hd_calls,omitempty"`
	MediaLoad                         *int               `json:"media_load,omitempty"`
	ResourceURI                       string             `json:"resource_uri"`
	WorkerVMConfigurationID           *int               `json:"workervm_configuration_id,omitempty"`
	WorkerVMConfigurationLocationName *string            `json:"workervm_configuration_location_name,omitempty"`
	WorkerVMConfigurationName         *string            `json:"workervm_configuration_name,omitempty"`
}

// CloudOverflowLocation represents a cloud overflow location status
type CloudOverflowLocation struct {
	FreeHDCalls      int    `json:"free_hd_calls"`
	ID               int    `json:"id"`
	MaxHDCalls       int    `json:"max_hd_calls"`
	MediaLoad        int    `json:"media_load"`
	Name             string `json:"name"`
	ResourceURI      string `json:"resource_uri"`
	SystemLocationID int    `json:"systemlocation_id"`
}

// ConferenceShard represents a conference shard status
type ConferenceShard struct {
	Conference         string             `json:"conference"`
	GuestsMuted        bool               `json:"guests_muted"`
	ID                 string             `json:"id"`
	IsDirect           bool               `json:"is_direct"`
	IsLocked           bool               `json:"is_locked"`
	IsStarted          bool               `json:"is_started"`
	Node               string             `json:"node"`
	ResourceURI        string             `json:"resource_uri"`
	ServiceType        string             `json:"service_type"`
	StartTime          *util.InfinityTime `json:"start_time,omitempty"`
	SystemLocation     string             `json:"system_location"`
	Tag                string             `json:"tag"`
	TranscodingEnabled bool               `json:"transcoding_enabled"`
}

// ConferenceSync represents a conference sync status
type ConferenceSync struct {
	ConfigurationID          int                `json:"configuration_id"`
	DevicesCreated           int                `json:"devices_created"`
	DevicesDeleted           int                `json:"devices_deleted"`
	DevicesUnchanged         int                `json:"devices_unchanged"`
	DevicesUpdated           int                `json:"devices_updated"`
	EndUsersCreated          int                `json:"end_users_created"`
	EndUsersDeleted          int                `json:"end_users_deleted"`
	EndUsersUnchanged        int                `json:"end_users_unchanged"`
	EndUsersUpdated          int                `json:"end_users_updated"`
	ID                       int                `json:"id"`
	LastUpdated              *util.InfinityTime `json:"last_updated,omitempty"`
	ResourceURI              string             `json:"resource_uri"`
	SyncErrors               int                `json:"sync_errors"`
	SyncLastErrorDescription string             `json:"sync_last_error_description"`
	SyncProgress             int                `json:"sync_progress"`
	SyncStatus               string             `json:"sync_status"`
	VMRsCreated              int                `json:"vmrs_created"`
	VMRsDeleted              int                `json:"vmrs_deleted"`
	VMRsUnchanged            int                `json:"vmrs_unchanged"`
	VMRsUpdated              int                `json:"vmrs_updated"`
}

// ExchangeScheduler represents an Exchange scheduler status
type ExchangeScheduler struct {
	ExchangeConnectorID int                `json:"exchange_connector_id"`
	ID                  int                `json:"id"`
	LastModifiedTime    *util.InfinityTime `json:"last_modified_time,omitempty"`
	ResourceURI         string             `json:"resource_uri"`
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
	ConfigurationID      int                `json:"configuration_id"`
	ID                   int                `json:"id"`
	LastAttemptedContact *util.InfinityTime `json:"last_attempted_contact,omitempty"`
	LastUpdated          *util.InfinityTime `json:"last_updated,omitempty"`
	Name                 string             `json:"name"`
	Primary              bool               `json:"primary"`
	ResourceURI          string             `json:"resource_uri"`
	SyncStatus           string             `json:"sync_status"`
	UpgradeStatus        string             `json:"upgrade_status"`
	Version              string             `json:"version"`
}

// MJXEndpoint represents an MJX endpoint status
type MJXEndpoint struct {
	EndpointAddress    string             `json:"endpoint_address"`
	EndpointName       string             `json:"endpoint_name"`
	EndpointType       string             `json:"endpoint_type"`
	ID                 int                `json:"id"`
	LastContactTime    *util.InfinityTime `json:"last_contact_time,omitempty"`
	LastWorker         string             `json:"last_worker"`
	MJXIntegrationName string             `json:"mjx_integration_name"`
	NumberOfMeetings   int                `json:"number_of_meetings"`
	ResourceURI        string             `json:"resource_uri"`
	RoomEmail          string             `json:"room_email"`
}

// MJXMeeting represents an MJX meeting status
type MJXMeeting struct {
	Alias                        string             `json:"alias"`
	EndTime                      *util.InfinityTime `json:"end_time,omitempty"`
	EndpointName                 string             `json:"endpoint_name"`
	ID                           int                `json:"id"`
	LastModifiedTime             *util.InfinityTime `json:"last_modified_time,omitempty"`
	MatchedMeetingProcessingRule string             `json:"matched_meeting_processing_rule"`
	MeetingID                    string             `json:"meeting_id"`
	MJXIntegrationID             int                `json:"mjx_integration_id"`
	MJXIntegrationName           string             `json:"mjx_integration_name"`
	OrganizerEmail               string             `json:"organizer_email"`
	OrganizerName                string             `json:"organizer_name"`
	ResourceURI                  string             `json:"resource_uri"`
	RoomEmail                    string             `json:"room_email"`
	StartTime                    *util.InfinityTime `json:"start_time,omitempty"`
	Subject                      string             `json:"subject"`
	WorkerID                     int                `json:"worker_id"`
}

// RegistrationAlias represents a registration alias status
type RegistrationAlias struct {
	Alias          string             `json:"alias"`
	ID             int                `json:"id"`
	IsNatted       bool               `json:"is_natted"`
	Node           string             `json:"node"`
	Protocol       string             `json:"protocol"`
	PushToken      string             `json:"push_token"`
	RegistrationID string             `json:"registration_id"`
	RemoteAddress  string             `json:"remote_address"`
	ResourceURI    string             `json:"resource_uri"`
	StartTime      *util.InfinityTime `json:"start_time,omitempty"`
	Username       string             `json:"username"`
}

// SchedulingOperation represents a scheduling operation status
type SchedulingOperation struct {
	CreationTime     *util.InfinityTime `json:"creation_time,omitempty"`
	ErrorCode        string             `json:"error_code"`
	ErrorDescription string             `json:"error_description"`
	ID               int                `json:"id"`
	ResourceID       *int               `json:"resource_id,omitempty"`
	ResourceURI      string             `json:"resource_uri"`
	Success          bool               `json:"success"`
	TransactionUUID  string             `json:"transaction_uuid"`
}

// SnapshotRequest represents a snapshot request status
type SnapshotRequest struct {
	CreatedAt   *util.InfinityTime `json:"created_at,omitempty"`
	DownloadURI string             `json:"download_uri"`
	Message     string             `json:"message"`
	ResourceURI string             `json:"resource_uri"`
	State       string             `json:"state"`
	UpdatedAt   *util.InfinityTime `json:"updated_at,omitempty"`
}

// SystemLocation represents a system location status
type SystemLocation struct {
	ID              int     `json:"id"`
	MaxAudioCalls   int     `json:"max_audio_calls"`
	MaxConnections  string  `json:"max_connections"`
	MaxFullHDCalls  int     `json:"max_full_hd_calls"`
	MaxHDCalls      int     `json:"max_hd_calls"`
	MaxMediaTokens  int     `json:"max_media_tokens"`
	MaxSDCalls      int     `json:"max_sd_calls"`
	MediaLoad       float64 `json:"media_load"`
	MediaTokensUsed int     `json:"media_tokens_used"`
	Name            string  `json:"name"`
	ResourceURI     string  `json:"resource_uri"`
}

// TeamsNode represents a Teams node status
type TeamsNode struct {
	CallCount        int                `json:"call_count"`
	EventHubID       *string            `json:"eventhub_id,omitempty"`
	HeartbeatTime    *util.InfinityTime `json:"heartbeat_time,omitempty"`
	ID               string             `json:"id"`
	InstanceStatus   *string            `json:"instance_status,omitempty"`
	IPAddress        string             `json:"ip_address"`
	MaxCalls         *int               `json:"max_calls,omitempty"`
	MediaLoad        int                `json:"media_load"`
	Name             string             `json:"name"`
	PrivateIPAddress string             `json:"private_ip_address"`
	ResourceURI      string             `json:"resource_uri"`
	ScalesetID       string             `json:"scaleset_id"`
	StartTime        *util.InfinityTime `json:"start_time,omitempty"`
	State            *string            `json:"state,omitempty"`
}

// TeamsNodeCall represents a Teams node call status
type TeamsNodeCall struct {
	Destination   string             `json:"destination"`
	HeartbeatTime *util.InfinityTime `json:"heartbeat_time,omitempty"`
	ID            string             `json:"id"`
	ResourceURI   string             `json:"resource_uri"`
	Source        string             `json:"source"`
	StartTime     *util.InfinityTime `json:"start_time,omitempty"`
	State         *string            `json:"state,omitempty"`
	TeamsNodeID   string             `json:"teamsnode_id"`
}

// WorkerVM represents a conferencing node VM status
type WorkerVM struct {
	BootTime              *util.InfinityTime `json:"boot_time,omitempty"`
	ConfigurationID       int                `json:"configuration_id"`
	CPUCapabilities       string             `json:"cpu_capabilities"`
	CPUCount              int                `json:"cpu_count"`
	CPUModel              string             `json:"cpu_model"`
	DeployStatus          string             `json:"deploy_status"`
	Hypervisor            string             `json:"hypervisor"`
	ID                    int                `json:"id"`
	LastAttemptedContact  *util.InfinityTime `json:"last_attempted_contact,omitempty"`
	LastUpdated           *util.InfinityTime `json:"last_updated,omitempty"`
	MaintenanceMode       bool               `json:"maintenance_mode"`
	MaintenanceModeReason string             `json:"maintenance_mode_reason"`
	MaxAudioCalls         int                `json:"max_audio_calls"`
	MaxDirectParticipants int                `json:"max_direct_participants"`
	MaxFullHDCalls        int                `json:"max_full_hd_calls"`
	MaxHDCalls            int                `json:"max_hd_calls"`
	MaxMediaTokens        int                `json:"max_media_tokens"`
	MaxSDCalls            int                `json:"max_sd_calls"`
	MediaLoad             int                `json:"media_load"`
	MediaTokensUsed       int                `json:"media_tokens_used"`
	Name                  string             `json:"name"`
	NodeType              string             `json:"node_type"`
	ResourceURI           string             `json:"resource_uri"`
	SignalingCount        int                `json:"signaling_count"`
	SyncStatus            string             `json:"sync_status"`
	SystemLocation        string             `json:"system_location"`
	TotalRAM              int                `json:"total_ram"`
	UpgradeStatus         string             `json:"upgrade_status"`
	Version               string             `json:"version"`
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

// LicensingResponse wraps a singleton licensing payload in a list.
type LicensingResponse struct {
	Meta    Meta        `json:"meta"`
	Objects []Licensing `json:"objects"`
}
