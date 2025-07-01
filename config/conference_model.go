package config

// Conference represents a conference configuration
type Conference struct {
	ID                 int    `json:"id,omitempty"`
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	ServiceType        string `json:"service_type"`
	PIN                string `json:"pin,omitempty"`
	GuestPIN           string `json:"guest_pin,omitempty"`
	AllowGuests        bool   `json:"allow_guests"`
	GuestsMuted        bool   `json:"guests_muted"`
	HostsCanUnmute     bool   `json:"hosts_can_unmute"`
	MaxPixelsPerSecond int    `json:"max_pixels_per_second,omitempty"`
	Tag                string `json:"tag,omitempty"`
	ResourceURI        string `json:"resource_uri,omitempty"`
}

// ConferenceCreateRequest represents a request to create a conference
type ConferenceCreateRequest struct {
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	ServiceType        string `json:"service_type"`
	PIN                string `json:"pin,omitempty"`
	GuestPIN           string `json:"guest_pin,omitempty"`
	AllowGuests        bool   `json:"allow_guests"`
	GuestsMuted        bool   `json:"guests_muted"`
	HostsCanUnmute     bool   `json:"hosts_can_unmute"`
	MaxPixelsPerSecond int    `json:"max_pixels_per_second,omitempty"`
	Tag                string `json:"tag,omitempty"`
}

// ConferenceUpdateRequest represents a request to update a conference
type ConferenceUpdateRequest struct {
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	PIN                string `json:"pin,omitempty"`
	GuestPIN           string `json:"guest_pin,omitempty"`
	AllowGuests        *bool  `json:"allow_guests,omitempty"`
	GuestsMuted        *bool  `json:"guests_muted,omitempty"`
	HostsCanUnmute     *bool  `json:"hosts_can_unmute,omitempty"`
	MaxPixelsPerSecond *int   `json:"max_pixels_per_second,omitempty"`
	Tag                string `json:"tag,omitempty"`
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
