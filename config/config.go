package config

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ClientInterface defines the interface for the HTTP client
type ClientInterface interface {
	GetJSON(ctx context.Context, endpoint string, result interface{}) error
	PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error
	PutJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error
	DeleteJSON(ctx context.Context, endpoint string, result interface{}) error
}

// Service handles Configuration API endpoints
type Service struct {
	client ClientInterface
}

// New creates a new Configuration API service
func New(client ClientInterface) *Service {
	return &Service{
		client: client,
	}
}

// Conference represents a conference configuration
type Conference struct {
	ID                    int    `json:"id,omitempty"`
	Name                  string `json:"name"`
	Description           string `json:"description,omitempty"`
	ServiceType           string `json:"service_type"`
	PIN                   string `json:"pin,omitempty"`
	GuestPIN              string `json:"guest_pin,omitempty"`
	AllowGuests           bool   `json:"allow_guests"`
	GuestsMuted           bool   `json:"guests_muted"`
	HostsCanUnmute        bool   `json:"hosts_can_unmute"`
	MaxPixelsPerSecond    int    `json:"max_pixels_per_second,omitempty"`
	Tag                   string `json:"tag,omitempty"`
	ResourceURI           string `json:"resource_uri,omitempty"`
}

// ConferenceCreateRequest represents a request to create a conference
type ConferenceCreateRequest struct {
	Name                  string `json:"name"`
	Description           string `json:"description,omitempty"`
	ServiceType           string `json:"service_type"`
	PIN                   string `json:"pin,omitempty"`
	GuestPIN              string `json:"guest_pin,omitempty"`
	AllowGuests           bool   `json:"allow_guests"`
	GuestsMuted           bool   `json:"guests_muted"`
	HostsCanUnmute        bool   `json:"hosts_can_unmute"`
	MaxPixelsPerSecond    int    `json:"max_pixels_per_second,omitempty"`
	Tag                   string `json:"tag,omitempty"`
}

// ConferenceUpdateRequest represents a request to update a conference
type ConferenceUpdateRequest struct {
	Name                  string `json:"name,omitempty"`
	Description           string `json:"description,omitempty"`
	PIN                   string `json:"pin,omitempty"`
	GuestPIN              string `json:"guest_pin,omitempty"`
	AllowGuests           *bool  `json:"allow_guests,omitempty"`
	GuestsMuted           *bool  `json:"guests_muted,omitempty"`
	HostsCanUnmute        *bool  `json:"hosts_can_unmute,omitempty"`
	MaxPixelsPerSecond    *int   `json:"max_pixels_per_second,omitempty"`
	Tag                   string `json:"tag,omitempty"`
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

// ListOptions contains options for listing resources
type ListOptions struct {
	Limit  int
	Offset int
	Search string
}

// ListConferences retrieves a list of conferences
func (s *Service) ListConferences(ctx context.Context, opts *ListOptions) (*ConferenceListResponse, error) {
	endpoint := "configuration/v1/conference/"
	
	if opts != nil {
		params := url.Values{}
		if opts.Limit > 0 {
			params.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.Offset > 0 {
			params.Set("offset", strconv.Itoa(opts.Offset))
		}
		if opts.Search != "" {
			params.Set("name__icontains", opts.Search)
		}
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ConferenceListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetConference retrieves a specific conference by ID
func (s *Service) GetConference(ctx context.Context, id int) (*Conference, error) {
	endpoint := fmt.Sprintf("configuration/v1/conference/%d/", id)
	
	var result Conference
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateConference creates a new conference
func (s *Service) CreateConference(ctx context.Context, req *ConferenceCreateRequest) (*Conference, error) {
	endpoint := "configuration/v1/conference/"
	
	var result Conference
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// UpdateConference updates an existing conference
func (s *Service) UpdateConference(ctx context.Context, id int, req *ConferenceUpdateRequest) (*Conference, error) {
	endpoint := fmt.Sprintf("configuration/v1/conference/%d/", id)
	
	var result Conference
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteConference deletes a conference
func (s *Service) DeleteConference(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/conference/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}

// Location represents a location configuration
type Location struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// LocationCreateRequest represents a request to create a location
type LocationCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// LocationUpdateRequest represents a request to update a location
type LocationUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// LocationListResponse represents the response from listing locations
type LocationListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []Location `json:"objects"`
}

// ListLocations retrieves a list of locations
func (s *Service) ListLocations(ctx context.Context, opts *ListOptions) (*LocationListResponse, error) {
	endpoint := "configuration/v1/location/"
	
	if opts != nil {
		params := url.Values{}
		if opts.Limit > 0 {
			params.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.Offset > 0 {
			params.Set("offset", strconv.Itoa(opts.Offset))
		}
		if opts.Search != "" {
			params.Set("name__icontains", opts.Search)
		}
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result LocationListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetLocation retrieves a specific location by ID
func (s *Service) GetLocation(ctx context.Context, id int) (*Location, error) {
	endpoint := fmt.Sprintf("configuration/v1/location/%d/", id)
	
	var result Location
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateLocation creates a new location
func (s *Service) CreateLocation(ctx context.Context, req *LocationCreateRequest) (*Location, error) {
	endpoint := "configuration/v1/location/"
	
	var result Location
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// UpdateLocation updates an existing location
func (s *Service) UpdateLocation(ctx context.Context, id int, req *LocationUpdateRequest) (*Location, error) {
	endpoint := fmt.Sprintf("configuration/v1/location/%d/", id)
	
	var result Location
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteLocation deletes a location
func (s *Service) DeleteLocation(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/location/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}