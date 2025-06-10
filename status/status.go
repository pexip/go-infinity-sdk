package status

import (
	"context"
	"fmt"
	"time"

	"github.com/pexip/go-infinity-sdk/options"
)

// ClientInterface defines the interface for the HTTP client
type ClientInterface interface {
	GetJSON(ctx context.Context, endpoint string, result interface{}) error
}

// Service handles Status API endpoints
type Service struct {
	client ClientInterface
}

// New creates a new Status API service
func New(client ClientInterface) *Service {
	return &Service{
		client: client,
	}
}

// SystemStatus represents the overall system status
type SystemStatus struct {
	Status      string    `json:"status"`
	Version     string    `json:"version"`
	Uptime      int       `json:"uptime"`
	Timestamp   time.Time `json:"timestamp"`
	HostName    string    `json:"hostname"`
	TotalMemory int64     `json:"total_memory"`
	UsedMemory  int64     `json:"used_memory"`
	CPULoad     float64   `json:"cpu_load"`
}

// Conference represents a conference status
type ConferenceStatus struct {
	ID                    int       `json:"id"`
	Name                  string    `json:"name"`
	ServiceType           string    `json:"service_type"`
	Started               bool      `json:"started"`
	Locked                bool      `json:"locked"`
	GuestsCanPresent      bool      `json:"guests_can_present"`
	DirectMedia           bool      `json:"direct_media"`
	ParticipantCount      int       `json:"participant_count"`
	MaxPixelsPerSecond    int       `json:"max_pixels_per_second"`
	ResourceURI           string    `json:"resource_uri"`
	StartTime             time.Time `json:"start_time,omitempty"`
}

// Participant represents a participant status
type Participant struct {
	ID                     string    `json:"id"`
	UUID                   string    `json:"uuid"`
	DisplayName            string    `json:"display_name"`
	ConferenceID           int       `json:"conference_id"`
	ConferenceName         string    `json:"conference_name"`
	ServiceType            string    `json:"service_type"`
	Role                   string    `json:"role"`
	ConnectTime            time.Time `json:"connect_time"`
	IsOnHold               bool      `json:"is_on_hold"`
	IsMuted                bool      `json:"is_muted"`
	IsGuest                bool      `json:"is_guest"`
	IsWaiting              bool      `json:"is_waiting"`
	HasMedia               bool      `json:"has_media"`
	IsPresenting           bool      `json:"is_presenting"`
	Bandwidth              int       `json:"bandwidth"`
	CallDirection          string    `json:"call_direction"`
	CallUUID               string    `json:"call_uuid"`
	DestinationAlias       string    `json:"destination_alias"`
	SourceAlias            string    `json:"source_alias"`
	RemoteAddress          string    `json:"remote_address"`
	RemotePort             int       `json:"remote_port"`
	Protocol               string    `json:"protocol"`
	VendorID               string    `json:"vendor_id"`
	ParentID               string    `json:"parent_id"`
	SystemLocation         string    `json:"system_location"`
	NodeID                 string    `json:"node_id"`
	ResourceURI            string    `json:"resource_uri"`
}

// Worker represents a worker node status
type Worker struct {
	NodeID         string    `json:"node_id"`
	HostName       string    `json:"hostname"`
	Status         string    `json:"status"`
	CPU            float64   `json:"cpu"`
	Memory         float64   `json:"memory"`
	Alarms         []string  `json:"alarms"`
	Conferences    int       `json:"conferences"`
	Participants   int       `json:"participants"`
	LastSeen       time.Time `json:"last_seen"`
	ResourceURI    string    `json:"resource_uri"`
}

// Alarm represents a system alarm
type Alarm struct {
	ID          int       `json:"id"`
	Level       string    `json:"level"`
	Name        string    `json:"name"`
	Details     string    `json:"details"`
	Instance    string    `json:"instance"`
	NodeID      string    `json:"node_id"`
	Cleared     bool      `json:"cleared"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	ResourceURI string    `json:"resource_uri"`
}

// ConferenceListResponse represents the response from listing conference statuses
type ConferenceListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ConferenceStatus `json:"objects"`
}

// ParticipantListResponse represents the response from listing participants
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

// WorkerListResponse represents the response from listing workers
type WorkerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []Worker `json:"objects"`
}

// AlarmListResponse represents the response from listing alarms
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

// ListOptions contains options for listing resources
type ListOptions = options.BaseListOptions

// GetSystemStatus retrieves the overall system status
func (s *Service) GetSystemStatus(ctx context.Context) (*SystemStatus, error) {
	endpoint := "status/v1/system_status/"
	
	var result SystemStatus
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// ListConferences retrieves a list of conference statuses
func (s *Service) ListConferences(ctx context.Context, opts *ListOptions) (*ConferenceListResponse, error) {
	endpoint := "status/v1/conference/"
	
	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ConferenceListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetConference retrieves a specific conference status by ID
func (s *Service) GetConference(ctx context.Context, id int) (*ConferenceStatus, error) {
	endpoint := fmt.Sprintf("status/v1/conference/%d/", id)
	
	var result ConferenceStatus
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// ListParticipants retrieves a list of participants
func (s *Service) ListParticipants(ctx context.Context, opts *ListOptions) (*ParticipantListResponse, error) {
	endpoint := "status/v1/participant/"
	
	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ParticipantListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetParticipant retrieves a specific participant by UUID
func (s *Service) GetParticipant(ctx context.Context, uuid string) (*Participant, error) {
	endpoint := fmt.Sprintf("status/v1/participant/%s/", uuid)
	
	var result Participant
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// ListWorkers retrieves a list of worker nodes
func (s *Service) ListWorkers(ctx context.Context, opts *ListOptions) (*WorkerListResponse, error) {
	endpoint := "status/v1/worker/"
	
	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result WorkerListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetWorker retrieves a specific worker by node ID
func (s *Service) GetWorker(ctx context.Context, nodeID string) (*Worker, error) {
	endpoint := fmt.Sprintf("status/v1/worker/%s/", nodeID)
	
	var result Worker
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// ListAlarms retrieves a list of system alarms
func (s *Service) ListAlarms(ctx context.Context, opts *ListOptions) (*AlarmListResponse, error) {
	endpoint := "status/v1/alarm/"
	
	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result AlarmListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetAlarm retrieves a specific alarm by ID
func (s *Service) GetAlarm(ctx context.Context, id int) (*Alarm, error) {
	endpoint := fmt.Sprintf("status/v1/alarm/%d/", id)
	
	var result Alarm
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}