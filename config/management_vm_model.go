/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// ManagementVM represents a management VM configuration
type ManagementVM struct {
	ID                          int            `json:"id,omitempty"`
	Name                        string         `json:"name"`
	Description                 string         `json:"description,omitempty"`
	Address                     string         `json:"address"`
	Netmask                     string         `json:"netmask"`
	Gateway                     string         `json:"gateway"`
	Hostname                    string         `json:"hostname"`
	Domain                      string         `json:"domain"`
	AlternativeFQDN             string         `json:"alternative_fqdn,omitempty"`
	IPV6Address                 *string        `json:"ipv6_address,omitempty"`
	IPV6Gateway                 *string        `json:"ipv6_gateway,omitempty"`
	MTU                         int            `json:"mtu"`
	StaticNATAddress            *string        `json:"static_nat_address,omitempty"`
	DNSServers                  []DNSServer    `json:"dns_servers,omitempty"`
	NTPServers                  []NTPServer    `json:"ntp_servers,omitempty"`
	SyslogServers               []SyslogServer `json:"syslog_servers,omitempty"`
	StaticRoutes                []string       `json:"static_routes,omitempty"`
	EventSinks                  []string       `json:"event_sinks,omitempty"`
	HTTPProxy                   *string        `json:"http_proxy,omitempty"`
	TLSCertificate              *string        `json:"tls_certificate,omitempty"`
	EnableSSH                   string         `json:"enable_ssh"`
	SSHAuthorizedKeys           []string       `json:"ssh_authorized_keys,omitempty"`
	SSHAuthorizedKeysUseCloud   bool           `json:"ssh_authorized_keys_use_cloud"`
	SecondaryConfigPassphrase   string         `json:"secondary_config_passphrase,omitempty"`
	SNMPMode                    string         `json:"snmp_mode"`
	SNMPCommunity               string         `json:"snmp_community,omitempty"`
	SNMPUsername                string         `json:"snmp_username,omitempty"`
	SNMPAuthenticationPassword  string         `json:"snmp_authentication_password,omitempty"`
	SNMPPrivacyPassword         string         `json:"snmp_privacy_password,omitempty"`
	SNMPSystemContact           string         `json:"snmp_system_contact,omitempty"`
	SNMPSystemLocation          string         `json:"snmp_system_location,omitempty"`
	SNMPNetworkManagementSystem *string        `json:"snmp_network_management_system,omitempty"`
	Initializing                bool           `json:"initializing"`
	Primary                     bool           `json:"primary,omitempty"`
	ResourceURI                 string         `json:"resource_uri,omitempty"`
}

// ManagementVMCreateRequest represents a request to create a management VM
type ManagementVMCreateRequest struct {
	Name                        string   `json:"name"`
	Description                 string   `json:"description,omitempty"`
	Address                     string   `json:"address"`
	Netmask                     string   `json:"netmask"`
	Gateway                     string   `json:"gateway"`
	Hostname                    string   `json:"hostname"`
	Domain                      string   `json:"domain"`
	AlternativeFQDN             string   `json:"alternative_fqdn,omitempty"`
	IPV6Address                 *string  `json:"ipv6_address,omitempty"`
	IPV6Gateway                 *string  `json:"ipv6_gateway,omitempty"`
	MTU                         int      `json:"mtu"`
	StaticNATAddress            *string  `json:"static_nat_address,omitempty"`
	DNSServers                  []string `json:"dns_servers,omitempty"`
	NTPServers                  []string `json:"ntp_servers,omitempty"`
	SyslogServers               []string `json:"syslog_servers,omitempty"`
	StaticRoutes                []string `json:"static_routes,omitempty"`
	EventSinks                  []string `json:"event_sinks,omitempty"`
	HTTPProxy                   *string  `json:"http_proxy,omitempty"`
	TLSCertificate              *string  `json:"tls_certificate,omitempty"`
	EnableSSH                   string   `json:"enable_ssh"`
	SSHAuthorizedKeys           []string `json:"ssh_authorized_keys,omitempty"`
	SSHAuthorizedKeysUseCloud   bool     `json:"ssh_authorized_keys_use_cloud"`
	SecondaryConfigPassphrase   string   `json:"secondary_config_passphrase,omitempty"`
	SNMPMode                    string   `json:"snmp_mode"`
	SNMPCommunity               string   `json:"snmp_community,omitempty"`
	SNMPUsername                string   `json:"snmp_username,omitempty"`
	SNMPAuthenticationPassword  string   `json:"snmp_authentication_password,omitempty"`
	SNMPPrivacyPassword         string   `json:"snmp_privacy_password,omitempty"`
	SNMPSystemContact           string   `json:"snmp_system_contact,omitempty"`
	SNMPSystemLocation          string   `json:"snmp_system_location,omitempty"`
	SNMPNetworkManagementSystem *string  `json:"snmp_network_management_system,omitempty"`
	Initializing                bool     `json:"initializing"`
}

// ManagementVMUpdateRequest represents a request to update a management VM
type ManagementVMUpdateRequest struct {
	Name                        string   `json:"name,omitempty"`
	Description                 string   `json:"description"`
	Address                     string   `json:"address,omitempty"`
	Netmask                     string   `json:"netmask,omitempty"`
	Gateway                     string   `json:"gateway,omitempty"`
	Hostname                    string   `json:"hostname,omitempty"`
	Domain                      string   `json:"domain,omitempty"`
	AlternativeFQDN             string   `json:"alternative_fqdn,omitempty"`
	IPV6Address                 *string  `json:"ipv6_address,omitempty"`
	IPV6Gateway                 *string  `json:"ipv6_gateway,omitempty"`
	MTU                         int      `json:"mtu,omitempty"`
	StaticNATAddress            *string  `json:"static_nat_address,omitempty"`
	DNSServers                  []string `json:"dns_servers"`
	NTPServers                  []string `json:"ntp_servers"`
	SyslogServers               []string `json:"syslog_servers"`
	StaticRoutes                []string `json:"static_routes"`
	EventSinks                  []string `json:"event_sinks,omitempty"`
	HTTPProxy                   *string  `json:"http_proxy,omitempty"`
	TLSCertificate              *string  `json:"tls_certificate"`
	EnableSSH                   string   `json:"enable_ssh,omitempty"`
	SSHAuthorizedKeys           []string `json:"ssh_authorized_keys"`
	SSHAuthorizedKeysUseCloud   bool     `json:"ssh_authorized_keys_use_cloud"`
	SecondaryConfigPassphrase   string   `json:"secondary_config_passphrase,omitempty"`
	SNMPMode                    string   `json:"snmp_mode,omitempty"`
	SNMPCommunity               string   `json:"snmp_community,omitempty"`
	SNMPUsername                string   `json:"snmp_username,omitempty"`
	SNMPAuthenticationPassword  string   `json:"snmp_authentication_password,omitempty"`
	SNMPPrivacyPassword         string   `json:"snmp_privacy_password,omitempty"`
	SNMPSystemContact           string   `json:"snmp_system_contact,omitempty"`
	SNMPSystemLocation          string   `json:"snmp_system_location,omitempty"`
	SNMPNetworkManagementSystem *string  `json:"snmp_network_management_system"`
	Initializing                bool     `json:"initializing"`
}

// ManagementVMListResponse represents the response from listing management VMs
type ManagementVMListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []ManagementVM `json:"objects"`
}
