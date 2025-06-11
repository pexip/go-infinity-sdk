// Package history provides access to the Pexip Infinity History API.
// It allows retrieval of historical data for conferences, participants, and media streams
// with support for time-based filtering and search capabilities.
package history

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/pexip/go-infinity-sdk/options"
)

// ClientInterface defines the interface for the HTTP client
type ClientInterface interface {
	GetJSON(ctx context.Context, endpoint string, result interface{}) error
}

// Service handles History API endpoints
type Service struct {
	client ClientInterface
}

// New creates a new History API service
func New(client ClientInterface) *Service {
	return &Service{
		client: client,
	}
}

// Conference represents a conference history record
type Conference struct {
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

// ConferenceListResponse represents the response from listing conference history
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

// ListOptions contains options for listing historical records
type ListOptions = options.TimeFilteredListOptions

// ListConferences retrieves a list of conference history records
func (s *Service) ListConferences(ctx context.Context, opts *ListOptions) (*ConferenceListResponse, error) {
	endpoint := "history/v1/conference/"

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

// GetConference retrieves a specific conference history record by ID
func (s *Service) GetConference(ctx context.Context, id int) (*Conference, error) {
	endpoint := fmt.Sprintf("history/v1/conference/%d/", id)

	var result Conference
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// ListParticipants retrieves a list of participant history records
func (s *Service) ListParticipants(ctx context.Context, opts *ListOptions) (*ParticipantListResponse, error) {
	endpoint := "history/v1/participant/"

	if opts != nil {
		params := opts.ToURLValuesWithSearchField("display_name__icontains")
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ParticipantListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetParticipant retrieves a specific participant history record by ID
func (s *Service) GetParticipant(ctx context.Context, id int) (*Participant, error) {
	endpoint := fmt.Sprintf("history/v1/participant/%d/", id)

	var result Participant
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// ListParticipantsByConference retrieves participant history for a specific conference
func (s *Service) ListParticipantsByConference(ctx context.Context, conferenceID int, opts *ListOptions) (*ParticipantListResponse, error) {
	endpoint := "history/v1/participant/"

	params := url.Values{}
	params.Set("conference_id", strconv.Itoa(conferenceID))

	if opts != nil {
		optParams := opts.ToURLValuesWithSearchField("display_name__icontains")
		for key, values := range optParams {
			for _, value := range values {
				params.Set(key, value)
			}
		}
	}

	endpoint += "?" + params.Encode()

	var result ParticipantListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// ListMediaStreams retrieves a list of media stream history records
func (s *Service) ListMediaStreams(ctx context.Context, opts *ListOptions) (*MediaStreamListResponse, error) {
	endpoint := "history/v1/media_stream/"

	if opts != nil {
		params := opts.BaseListOptions.ToURLValues()
		if opts.StartTime != nil {
			params.Set("start_time__gte", opts.StartTime.Format(time.RFC3339))
		}
		if opts.EndTime != nil {
			params.Set("end_time__lt", opts.EndTime.Format(time.RFC3339))
		}
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result MediaStreamListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetMediaStream retrieves a specific media stream history record by ID
func (s *Service) GetMediaStream(ctx context.Context, id int) (*MediaStream, error) {
	endpoint := fmt.Sprintf("history/v1/media_stream/%d/", id)

	var result MediaStream
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// ListMediaStreamsByParticipant retrieves media stream history for a specific participant
func (s *Service) ListMediaStreamsByParticipant(ctx context.Context, participantID int, opts *ListOptions) (*MediaStreamListResponse, error) {
	endpoint := "history/v1/media_stream/"

	params := url.Values{}
	params.Set("participant_id", strconv.Itoa(participantID))

	if opts != nil {
		optParams := opts.BaseListOptions.ToURLValues()
		for key, values := range optParams {
			for _, value := range values {
				params.Set(key, value)
			}
		}
		if opts.StartTime != nil {
			params.Set("start_time__gte", opts.StartTime.Format(time.RFC3339))
		}
		if opts.EndTime != nil {
			params.Set("end_time__lt", opts.EndTime.Format(time.RFC3339))
		}
	}

	endpoint += "?" + params.Encode()

	var result MediaStreamListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
