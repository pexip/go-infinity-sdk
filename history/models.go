package history

import (
	"time"

	"github.com/pexip/go-infinity-sdk/v38/options"
)

// ListOptions contains options for listing historical records
type ListOptions = options.TimeFilteredListOptions

// Alarm represents an alarm history record
type Alarm struct {
	ID          int        `json:"id"`
	Details     string     `json:"details"`
	Identifier  int        `json:"identifier"`
	Instance    string     `json:"instance"`
	Level       string     `json:"level"`
	Name        string     `json:"name"`
	Node        string     `json:"node"`
	TimeRaised  *time.Time `json:"time_raised"`
	TimeLowered *time.Time `json:"time_lowered"`
	ResourceURI string     `json:"resource_uri"`
}

// AlarmListResponse represents the response from listing alarm history
type AlarmListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []Alarm `json:"objects"`
}

// Backplane represents a backplane history record
type Backplane struct {
	ID                   string     `json:"id"`
	ConferenceName       string     `json:"conference_name"`
	DisconnectReason     string     `json:"disconnect_reason"`
	Duration             *int       `json:"duration"`
	StartTime            *time.Time `json:"start_time"`
	EndTime              *time.Time `json:"end_time"`
	MediaNode            string     `json:"media_node"`
	Protocol             string     `json:"protocol"`
	ProxyNode            string     `json:"proxy_node"`
	RemoteConferenceName string     `json:"remote_conference_name"`
	RemoteMediaNode      string     `json:"remote_media_node"`
	RemoteNodeName       string     `json:"remote_node_name"`
	ServiceTag           string     `json:"service_tag"`
	SystemLocation       string     `json:"system_location"`
	Type                 string     `json:"type"`
	ResourceURI          string     `json:"resource_uri"`
}

// BackplaneListResponse represents the response from listing backplane history
type BackplaneListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []Backplane `json:"objects"`
}

// BackplaneMediaStream represents a backplane media stream history record
type BackplaneMediaStream struct {
	ID                int        `json:"id"`
	StreamID          string     `json:"stream_id"`
	StreamType        string     `json:"stream_type"`
	StartTime         *time.Time `json:"start_time"`
	EndTime           *time.Time `json:"end_time"`
	Node              string     `json:"node"`
	RxBitrate         int        `json:"rx_bitrate"`
	RxCodec           string     `json:"rx_codec"`
	RxFPS             float64    `json:"rx_fps"`
	RxPacketLoss      float64    `json:"rx_packet_loss"`
	RxPacketsLost     int        `json:"rx_packets_lost"`
	RxPacketsReceived int        `json:"rx_packets_received"`
	RxResolution      string     `json:"rx_resolution"`
	TxBitrate         int        `json:"tx_bitrate"`
	TxCodec           string     `json:"tx_codec"`
	TxFPS             float64    `json:"tx_fps"`
	TxPacketLoss      float64    `json:"tx_packet_loss"`
	TxPacketsLost     int        `json:"tx_packets_lost"`
	TxPacketsSent     int        `json:"tx_packets_sent"`
	TxResolution      string     `json:"tx_resolution"`
	ResourceURI       string     `json:"resource_uri"`
}

// BackplaneMediaStreamListResponse represents the response from listing backplane media stream history
type BackplaneMediaStreamListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []BackplaneMediaStream `json:"objects"`
}

// ConferenceRecord represents a conference history record
type ConferenceRecord struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	ServiceType         string    `json:"service_type"`
	Tag                 string    `json:"tag"`
	StartTime           time.Time `json:"start_time"`
	EndTime             time.Time `json:"end_time"`
	CreateTime          time.Time `json:"create_time"`
	DurationSeconds     int       `json:"duration_seconds"`
	TotalParticipants   int       `json:"total_participants"`
	MaxConcurrentGuests int       `json:"max_concurrent_guests"`
	MaxConcurrentHosts  int       `json:"max_concurrent_hosts"`
	TotalBandwidthKbps  int       `json:"total_bandwidth_kbps"`
	InstanceType        string    `json:"instance_type"`
	ResourceURI         string    `json:"resource_uri"`
}

// ConferenceRecordListResponse represents the response from listing conference history
type ConferenceRecordListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ConferenceRecord `json:"objects"`
}

// MediaStream represents a media stream history record
type MediaStream struct {
	ID              int       `json:"id"`
	ParticipantID   int       `json:"participant_id"`
	Node            string    `json:"node"`
	StreamType      string    `json:"stream_type"`
	Direction       string    `json:"direction"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	DurationSeconds int       `json:"duration_seconds"`
	RxPackets       int       `json:"rx_packets"`
	TxPackets       int       `json:"tx_packets"`
	RxBytes         int64     `json:"rx_bytes"`
	TxBytes         int64     `json:"tx_bytes"`
	RxPacketsLost   int       `json:"rx_packets_lost"`
	TxPacketsLost   int       `json:"tx_packets_lost"`
	RxJitter        float64   `json:"rx_jitter"`
	TxJitter        float64   `json:"tx_jitter"`
	Codec           string    `json:"codec"`
	Resolution      string    `json:"resolution"`
	FrameRate       float64   `json:"frame_rate"`
	Bitrate         int       `json:"bitrate"`
	PacketRate      float64   `json:"packet_rate"`
	ResourceURI     string    `json:"resource_uri"`
}

// MediaStreamListResponse represents the response from listing media stream history
type MediaStreamListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MediaStream `json:"objects"`
}

// Participant represents a participant history record
type Participant struct {
	ID               int       `json:"id"`
	ConferenceID     int       `json:"conference_id"`
	ConferenceName   string    `json:"conference_name"`
	LocalAlias       string    `json:"local_alias"`
	RemoteAlias      string    `json:"remote_alias"`
	DisplayName      string    `json:"display_name"`
	Role             string    `json:"role"`
	ServiceType      string    `json:"service_type"`
	CallDirection    string    `json:"call_direction"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	DurationSeconds  int       `json:"duration_seconds"`
	DisconnectReason string    `json:"disconnect_reason"`
	RemoteAddress    string    `json:"remote_address"`
	RemotePort       int       `json:"remote_port"`
	SIPCallID        string    `json:"sip_call_id"`
	Vendor           string    `json:"vendor"`
	UserAgent        string    `json:"user_agent"`
	TotalRxPackets   int       `json:"total_rx_packets"`
	TotalTxPackets   int       `json:"total_tx_packets"`
	TotalRxBytes     int64     `json:"total_rx_bytes"`
	TotalTxBytes     int64     `json:"total_tx_bytes"`
	MediaNode        string    `json:"media_node"`
	SignalingNode    string    `json:"signaling_node"`
	Encryption       string    `json:"encryption"`
	ConversationID   string    `json:"conversation_id"`
	CallUUID         string    `json:"call_uuid"`
	ParentID         string    `json:"parent_id"`
	ResourceURI      string    `json:"resource_uri"`
}

// ParticipantListResponse represents the response from listing participant history
type ParticipantListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []Participant `json:"objects"`
}

// RegistrationAlias represents a registration alias history record
type RegistrationAlias struct {
	ID             int        `json:"id"`
	Alias          string     `json:"alias"`
	RegistrationID string     `json:"registration_id"`
	Username       string     `json:"username"`
	Node           string     `json:"node"`
	Protocol       string     `json:"protocol"`
	RemoteAddress  string     `json:"remote_address"`
	IsNatted       bool       `json:"is_natted"`
	StartTime      *time.Time `json:"start_time"`
	EndTime        *time.Time `json:"end_time"`
	ResourceURI    string     `json:"resource_uri"`
}

// RegistrationAliasListResponse represents the response from listing registration alias history
type RegistrationAliasListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []RegistrationAlias `json:"objects"`
}

// WorkerVMStatusEvent represents a worker VM status event history record
type WorkerVMStatusEvent struct {
	ID                                int        `json:"id"`
	EventType                         string     `json:"event_type"`
	State                             *string    `json:"state"`
	Context                           *string    `json:"context"`
	Details                           *string    `json:"details"`
	TimeChanged                       *time.Time `json:"time_changed"`
	WorkerVMAddress                   string     `json:"workervm_address"`
	WorkerVMConfigurationID           *int       `json:"workervm_configuration_id"`
	WorkerVMConfigurationName         *string    `json:"workervm_configuration_name"`
	WorkerVMConfigurationLocationName *string    `json:"workervm_configuration_location_name"`
	ResourceURI                       string     `json:"resource_uri"`
}

// WorkerVMStatusEventListResponse represents the response from listing worker VM status event history
type WorkerVMStatusEventListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []WorkerVMStatusEvent `json:"objects"`
}
