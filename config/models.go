package config

import (
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/util"
)

// ListOptions contains options for listing resources
type ListOptions = options.SearchableListOptions

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

// ConferenceAlias represents a conference alias configuration
type ConferenceAlias struct {
	ID           int               `json:"id,omitempty"`
	Alias        string            `json:"alias"`
	Conference   string            `json:"conference"`
	Description  string            `json:"description,omitempty"`
	CreationTime util.InfinityTime `json:"creation_time,omitempty"`
	ResourceURI  string            `json:"resource_uri,omitempty"`
}

// ConferenceAliasCreateRequest represents a request to create a conference alias
type ConferenceAliasCreateRequest struct {
	Alias       string `json:"alias"`
	Conference  string `json:"conference"`
	Description string `json:"description,omitempty"`
}

// ConferenceAliasUpdateRequest represents a request to update a conference alias
type ConferenceAliasUpdateRequest struct {
	Alias       string `json:"alias,omitempty"`
	Conference  string `json:"conference,omitempty"`
	Description string `json:"description,omitempty"`
}

// ConferenceAliasListResponse represents the response from listing conference aliases
type ConferenceAliasListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ConferenceAlias `json:"objects"`
}

// EndUser represents an end user configuration
type EndUser struct {
	ID                  int      `json:"id,omitempty"`
	PrimaryEmailAddress string   `json:"primary_email_address"`
	FirstName           string   `json:"first_name,omitempty"`
	LastName            string   `json:"last_name,omitempty"`
	DisplayName         string   `json:"display_name,omitempty"`
	TelephoneNumber     string   `json:"telephone_number,omitempty"`
	MobileNumber        string   `json:"mobile_number,omitempty"`
	Title               string   `json:"title,omitempty"`
	Department          string   `json:"department,omitempty"`
	AvatarURL           string   `json:"avatar_url,omitempty"`
	UUID                string   `json:"uuid,omitempty"`
	UserGroups          []string `json:"user_groups,omitempty"`
	UserOID             *string  `json:"user_oid,omitempty"`
	ExchangeUserID      *string  `json:"exchange_user_id,omitempty"`
	MSExchangeGUID      *string  `json:"ms_exchange_guid,omitempty"`
	SyncTag             string   `json:"sync_tag,omitempty"`
	ResourceURI         string   `json:"resource_uri,omitempty"`
}

// EndUserCreateRequest represents a request to create an end user
type EndUserCreateRequest struct {
	PrimaryEmailAddress string   `json:"primary_email_address"`
	FirstName           string   `json:"first_name,omitempty"`
	LastName            string   `json:"last_name,omitempty"`
	DisplayName         string   `json:"display_name,omitempty"`
	TelephoneNumber     string   `json:"telephone_number,omitempty"`
	MobileNumber        string   `json:"mobile_number,omitempty"`
	Title               string   `json:"title,omitempty"`
	Department          string   `json:"department,omitempty"`
	AvatarURL           string   `json:"avatar_url,omitempty"`
	UserGroups          []string `json:"user_groups,omitempty"`
	UserOID             *string  `json:"user_oid,omitempty"`
	ExchangeUserID      *string  `json:"exchange_user_id,omitempty"`
	MSExchangeGUID      *string  `json:"ms_exchange_guid,omitempty"`
	SyncTag             string   `json:"sync_tag,omitempty"`
}

// EndUserUpdateRequest represents a request to update an end user
type EndUserUpdateRequest struct {
	PrimaryEmailAddress string   `json:"primary_email_address,omitempty"`
	FirstName           string   `json:"first_name,omitempty"`
	LastName            string   `json:"last_name,omitempty"`
	DisplayName         string   `json:"display_name,omitempty"`
	TelephoneNumber     string   `json:"telephone_number,omitempty"`
	MobileNumber        string   `json:"mobile_number,omitempty"`
	Title               string   `json:"title,omitempty"`
	Department          string   `json:"department,omitempty"`
	AvatarURL           string   `json:"avatar_url,omitempty"`
	UserGroups          []string `json:"user_groups,omitempty"`
	UserOID             *string  `json:"user_oid,omitempty"`
	ExchangeUserID      *string  `json:"exchange_user_id,omitempty"`
	MSExchangeGUID      *string  `json:"ms_exchange_guid,omitempty"`
	SyncTag             string   `json:"sync_tag,omitempty"`
}

// EndUserListResponse represents the response from listing end users
type EndUserListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []EndUser `json:"objects"`
}

// SystemLocation represents a system location configuration
type SystemLocation struct {
	ID                        int         `json:"id,omitempty"`
	Name                      string      `json:"name"`
	Description               string      `json:"description,omitempty"`
	MTU                       int         `json:"mtu,omitempty"`
	MediaQoS                  *int        `json:"media_qos,omitempty"`
	SignallingQoS             *int        `json:"signalling_qos,omitempty"`
	LocalMSSIPDomain          string      `json:"local_mssip_domain,omitempty"`
	DNSServers                []DNSServer `json:"dns_servers,omitempty"`
	NTPServers                []NTPServer `json:"ntp_servers,omitempty"`
	SyslogServers             []string    `json:"syslog_servers,omitempty"`
	H323Gatekeeper            *string     `json:"h323_gatekeeper,omitempty"`
	SIPProxy                  *string     `json:"sip_proxy,omitempty"`
	MSSIPProxy                *string     `json:"mssip_proxy,omitempty"`
	TeamsProxy                *string     `json:"teams_proxy,omitempty"`
	OverflowLocation1         *string     `json:"overflow_location1,omitempty"`
	OverflowLocation2         *string     `json:"overflow_location2,omitempty"`
	TranscodingLocation       *string     `json:"transcoding_location,omitempty"`
	BDPMPinChecksEnabled      string      `json:"bdpm_pin_checks_enabled,omitempty"`
	BDPMScanQuarantineEnabled string      `json:"bdpm_scan_quarantine_enabled,omitempty"`
	UseRelayCandidatesOnly    bool        `json:"use_relay_candidates_only,omitempty"`
	ResourceURI               string      `json:"resource_uri,omitempty"`
}

// SystemLocationCreateRequest represents a request to create a system location
type SystemLocationCreateRequest struct {
	Name                      string   `json:"name"`
	Description               string   `json:"description,omitempty"`
	MTU                       int      `json:"mtu,omitempty"`
	MediaQoS                  *int     `json:"media_qos,omitempty"`
	SignallingQoS             *int     `json:"signalling_qos,omitempty"`
	LocalMSSIPDomain          string   `json:"local_mssip_domain,omitempty"`
	DNSServers                []string `json:"dns_servers,omitempty"`
	NTPServers                []string `json:"ntp_servers,omitempty"`
	SyslogServers             []string `json:"syslog_servers,omitempty"`
	H323Gatekeeper            *string  `json:"h323_gatekeeper,omitempty"`
	SIPProxy                  *string  `json:"sip_proxy,omitempty"`
	MSSIPProxy                *string  `json:"mssip_proxy,omitempty"`
	TeamsProxy                *string  `json:"teams_proxy,omitempty"`
	OverflowLocation1         *string  `json:"overflow_location1,omitempty"`
	OverflowLocation2         *string  `json:"overflow_location2,omitempty"`
	TranscodingLocation       *string  `json:"transcoding_location,omitempty"`
	BDPMPinChecksEnabled      string   `json:"bdpm_pin_checks_enabled,omitempty"`
	BDPMScanQuarantineEnabled string   `json:"bdpm_scan_quarantine_enabled,omitempty"`
	UseRelayCandidatesOnly    bool     `json:"use_relay_candidates_only,omitempty"`
}

// SystemLocationUpdateRequest represents a request to update a system location
type SystemLocationUpdateRequest struct {
	Name                      string   `json:"name,omitempty"`
	Description               string   `json:"description,omitempty"`
	MTU                       int      `json:"mtu,omitempty"`
	MediaQoS                  *int     `json:"media_qos,omitempty"`
	SignallingQoS             *int     `json:"signalling_qos,omitempty"`
	LocalMSSIPDomain          string   `json:"local_mssip_domain,omitempty"`
	DNSServers                []string `json:"dns_servers,omitempty"`
	NTPServers                []string `json:"ntp_servers,omitempty"`
	SyslogServers             []string `json:"syslog_servers,omitempty"`
	H323Gatekeeper            *string  `json:"h323_gatekeeper,omitempty"`
	SIPProxy                  *string  `json:"sip_proxy,omitempty"`
	MSSIPProxy                *string  `json:"mssip_proxy,omitempty"`
	TeamsProxy                *string  `json:"teams_proxy,omitempty"`
	OverflowLocation1         *string  `json:"overflow_location1,omitempty"`
	OverflowLocation2         *string  `json:"overflow_location2,omitempty"`
	TranscodingLocation       *string  `json:"transcoding_location,omitempty"`
	BDPMPinChecksEnabled      string   `json:"bdpm_pin_checks_enabled,omitempty"`
	BDPMScanQuarantineEnabled string   `json:"bdpm_scan_quarantine_enabled,omitempty"`
	UseRelayCandidatesOnly    bool     `json:"use_relay_candidates_only,omitempty"`
}

// SystemLocationListResponse represents the response from listing system locations
type SystemLocationListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SystemLocation `json:"objects"`
}

// WorkerVM represents a worker VM configuration
type WorkerVM struct {
	ID                    int     `json:"id,omitempty"`
	Name                  string  `json:"name"`
	Hostname              string  `json:"hostname"`
	Domain                string  `json:"domain"`
	Address               string  `json:"address"`
	Netmask               string  `json:"netmask"`
	Gateway               string  `json:"gateway"`
	IPv6Address           *string `json:"ipv6_address,omitempty"`
	IPv6Gateway           *string `json:"ipv6_gateway,omitempty"`
	VMCPUCount            int     `json:"vm_cpu_count,omitempty"`
	VMSystemMemory        int     `json:"vm_system_memory,omitempty"`
	NodeType              string  `json:"node_type,omitempty"`
	Transcoding           bool    `json:"transcoding,omitempty"`
	Password              string  `json:"password,omitempty"`
	MaintenanceMode       bool    `json:"maintenance_mode,omitempty"`
	MaintenanceModeReason string  `json:"maintenance_mode_reason,omitempty"`
	SystemLocation        string  `json:"system_location"`
	ResourceURI           string  `json:"resource_uri,omitempty"`
}

// WorkerVMCreateRequest represents a request to create a worker VM
type WorkerVMCreateRequest struct {
	Name                  string  `json:"name"`
	Hostname              string  `json:"hostname"`
	Domain                string  `json:"domain"`
	Address               string  `json:"address"`
	Netmask               string  `json:"netmask"`
	Gateway               string  `json:"gateway"`
	IPv6Address           *string `json:"ipv6_address,omitempty"`
	IPv6Gateway           *string `json:"ipv6_gateway,omitempty"`
	VMCPUCount            int     `json:"vm_cpu_count,omitempty"`
	VMSystemMemory        int     `json:"vm_system_memory,omitempty"`
	NodeType              string  `json:"node_type,omitempty"`
	Transcoding           bool    `json:"transcoding,omitempty"`
	Password              string  `json:"password,omitempty"`
	MaintenanceMode       bool    `json:"maintenance_mode,omitempty"`
	MaintenanceModeReason string  `json:"maintenance_mode_reason,omitempty"`
	SystemLocation        string  `json:"system_location"`
}

// WorkerVMUpdateRequest represents a request to update a worker VM
type WorkerVMUpdateRequest struct {
	Name                  string  `json:"name,omitempty"`
	Hostname              string  `json:"hostname,omitempty"`
	Domain                string  `json:"domain,omitempty"`
	Address               string  `json:"address,omitempty"`
	Netmask               string  `json:"netmask,omitempty"`
	Gateway               string  `json:"gateway,omitempty"`
	IPv6Address           *string `json:"ipv6_address,omitempty"`
	IPv6Gateway           *string `json:"ipv6_gateway,omitempty"`
	VMCPUCount            int     `json:"vm_cpu_count,omitempty"`
	VMSystemMemory        int     `json:"vm_system_memory,omitempty"`
	NodeType              string  `json:"node_type,omitempty"`
	Transcoding           bool    `json:"transcoding,omitempty"`
	Password              string  `json:"password,omitempty"`
	MaintenanceMode       bool    `json:"maintenance_mode,omitempty"`
	MaintenanceModeReason string  `json:"maintenance_mode_reason,omitempty"`
	SystemLocation        string  `json:"system_location,omitempty"`
}

// WorkerVMListResponse represents the response from listing worker VMs
type WorkerVMListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []WorkerVM `json:"objects"`
}

// DNSServer represents a DNS server configuration
type DNSServer struct {
	ID          int    `json:"id,omitempty"`
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// DNSServerCreateRequest represents a request to create a DNS server
type DNSServerCreateRequest struct {
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
}

// DNSServerUpdateRequest represents a request to update a DNS server
type DNSServerUpdateRequest struct {
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
}

// DNSServerListResponse represents the response from listing DNS servers
type DNSServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []DNSServer `json:"objects"`
}

// NTPServer represents an NTP server configuration
type NTPServer struct {
	ID          int    `json:"id,omitempty"`
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
	Key         string `json:"key,omitempty"`
	KeyID       *int   `json:"key_id,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// NTPServerCreateRequest represents a request to create an NTP server
type NTPServerCreateRequest struct {
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
	Key         string `json:"key,omitempty"`
	KeyID       *int   `json:"key_id,omitempty"`
}

// NTPServerUpdateRequest represents a request to update an NTP server
type NTPServerUpdateRequest struct {
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
	Key         string `json:"key,omitempty"`
	KeyID       *int   `json:"key_id,omitempty"`
}

// NTPServerListResponse represents the response from listing NTP servers
type NTPServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []NTPServer `json:"objects"`
}

// StaticRoute represents a static route configuration
type StaticRoute struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Prefix      int    `json:"prefix"`
	Gateway     string `json:"gateway"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// StaticRouteCreateRequest represents a request to create a static route
type StaticRouteCreateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Prefix  int    `json:"prefix"`
	Gateway string `json:"gateway"`
}

// StaticRouteUpdateRequest represents a request to update a static route
type StaticRouteUpdateRequest struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Prefix  int    `json:"prefix,omitempty"`
	Gateway string `json:"gateway,omitempty"`
}

// StaticRouteListResponse represents the response from listing static routes
type StaticRouteListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []StaticRoute `json:"objects"`
}

// GlobalConfiguration represents the global system configuration (singleton resource)
type GlobalConfiguration struct {
	ID                           int      `json:"id,omitempty"`
	EnableWebRTC                 bool     `json:"enable_webrtc,omitempty"`
	EnableSIP                    bool     `json:"enable_sip,omitempty"`
	EnableH323                   bool     `json:"enable_h323,omitempty"`
	EnableRTMP                   bool     `json:"enable_rtmp,omitempty"`
	CryptoMode                   string   `json:"crypto_mode,omitempty"`
	MaxPixelsPerSecond           string   `json:"max_pixels_per_second,omitempty"`
	MediaPortsStart              int      `json:"media_ports_start,omitempty"`
	MediaPortsEnd                int      `json:"media_ports_end,omitempty"`
	SignallingPortsStart         int      `json:"signalling_ports_start,omitempty"`
	SignallingPortsEnd           int      `json:"signalling_ports_end,omitempty"`
	BurstingEnabled              bool     `json:"bursting_enabled,omitempty"`
	CloudProvider                string   `json:"cloud_provider,omitempty"`
	AWSAccessKey                 *string  `json:"aws_access_key,omitempty"`
	AWSSecretKey                 *string  `json:"aws_secret_key,omitempty"`
	AzureClientID                *string  `json:"azure_client_id,omitempty"`
	AzureSecret                  *string  `json:"azure_secret,omitempty"`
	GuestsOnlyTimeout            int      `json:"guests_only_timeout,omitempty"`
	WaitingForChairTimeout       int      `json:"waiting_for_chair_timeout,omitempty"`
	ConferenceCreatePermissions  string   `json:"conference_create_permissions,omitempty"`
	ConferenceCreationMode       string   `json:"conference_creation_mode,omitempty"`
	EnableAnalytics              bool     `json:"enable_analytics,omitempty"`
	EnableErrorReporting         bool     `json:"enable_error_reporting,omitempty"`
	BandwidthRestrictions        string   `json:"bandwidth_restrictions,omitempty"`
	AdministratorEmail           string   `json:"administrator_email,omitempty"`
	GlobalConferenceCreateGroups []string `json:"global_conference_create_groups,omitempty"`
	ResourceURI                  string   `json:"resource_uri,omitempty"`
}

// GlobalConfigurationUpdateRequest represents a request to update global configuration
type GlobalConfigurationUpdateRequest struct {
	EnableWebRTC                 *bool    `json:"enable_webrtc,omitempty"`
	EnableSIP                    *bool    `json:"enable_sip,omitempty"`
	EnableH323                   *bool    `json:"enable_h323,omitempty"`
	EnableRTMP                   *bool    `json:"enable_rtmp,omitempty"`
	CryptoMode                   string   `json:"crypto_mode,omitempty"`
	MaxPixelsPerSecond           string   `json:"max_pixels_per_second,omitempty"`
	MediaPortsStart              *int     `json:"media_ports_start,omitempty"`
	MediaPortsEnd                *int     `json:"media_ports_end,omitempty"`
	SignallingPortsStart         *int     `json:"signalling_ports_start,omitempty"`
	SignallingPortsEnd           *int     `json:"signalling_ports_end,omitempty"`
	BurstingEnabled              *bool    `json:"bursting_enabled,omitempty"`
	CloudProvider                string   `json:"cloud_provider,omitempty"`
	AWSAccessKey                 *string  `json:"aws_access_key,omitempty"`
	AWSSecretKey                 *string  `json:"aws_secret_key,omitempty"`
	AzureClientID                *string  `json:"azure_client_id,omitempty"`
	AzureSecret                  *string  `json:"azure_secret,omitempty"`
	GuestsOnlyTimeout            *int     `json:"guests_only_timeout,omitempty"`
	WaitingForChairTimeout       *int     `json:"waiting_for_chair_timeout,omitempty"`
	ConferenceCreatePermissions  string   `json:"conference_create_permissions,omitempty"`
	ConferenceCreationMode       string   `json:"conference_creation_mode,omitempty"`
	EnableAnalytics              *bool    `json:"enable_analytics,omitempty"`
	EnableErrorReporting         *bool    `json:"enable_error_reporting,omitempty"`
	BandwidthRestrictions        string   `json:"bandwidth_restrictions,omitempty"`
	AdministratorEmail           string   `json:"administrator_email,omitempty"`
	GlobalConferenceCreateGroups []string `json:"global_conference_create_groups,omitempty"`
}
