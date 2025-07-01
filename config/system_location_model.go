package config

// SystemLocation represents a system location configuration
type SystemLocation struct {
	ID                        int            `json:"id,omitempty"`
	Name                      string         `json:"name"`
	Description               string         `json:"description,omitempty"`
	MTU                       int            `json:"mtu,omitempty"`
	MediaQoS                  *int           `json:"media_qos,omitempty"`
	SignallingQoS             *int           `json:"signalling_qos,omitempty"`
	LocalMSSIPDomain          string         `json:"local_mssip_domain,omitempty"`
	DNSServers                []DNSServer    `json:"dns_servers,omitempty"`
	NTPServers                []NTPServer    `json:"ntp_servers,omitempty"`
	SyslogServers             []SyslogServer `json:"syslog_servers,omitempty"`
	H323Gatekeeper            *string        `json:"h323_gatekeeper,omitempty"`
	SIPProxy                  *string        `json:"sip_proxy,omitempty"`
	MSSIPProxy                *string        `json:"mssip_proxy,omitempty"`
	TeamsProxy                *string        `json:"teams_proxy,omitempty"`
	OverflowLocation1         *string        `json:"overflow_location1,omitempty"`
	OverflowLocation2         *string        `json:"overflow_location2,omitempty"`
	TranscodingLocation       *string        `json:"transcoding_location,omitempty"`
	BDPMPinChecksEnabled      string         `json:"bdpm_pin_checks_enabled,omitempty"`
	BDPMScanQuarantineEnabled string         `json:"bdpm_scan_quarantine_enabled,omitempty"`
	UseRelayCandidatesOnly    bool           `json:"use_relay_candidates_only,omitempty"`
	ResourceURI               string         `json:"resource_uri,omitempty"`
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
