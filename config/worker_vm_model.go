/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// WorkerVM represents a worker VM configuration
type WorkerVM struct {
	ID                         int      `json:"id,omitempty"`
	Name                       string   `json:"name"`
	Description                string   `json:"description,omitempty"`
	Hostname                   string   `json:"hostname"`
	Domain                     string   `json:"domain"`
	Address                    string   `json:"address"`
	Netmask                    string   `json:"netmask"`
	Gateway                    string   `json:"gateway"`
	VMCPUCount                 int      `json:"vm_cpu_count,omitempty"`
	VMSystemMemory             int      `json:"vm_system_memory,omitempty"`
	DeploymentType             string   `json:"deployment_type,omitempty"`
	NodeType                   string   `json:"node_type,omitempty"`
	Transcoding                bool     `json:"transcoding,omitempty"`
	CloudBursting              bool     `json:"cloud_bursting,omitempty"`
	Password                   string   `json:"password,omitempty"`
	MaintenanceMode            bool     `json:"maintenance_mode,omitempty"`
	MaintenanceModeReason      string   `json:"maintenance_mode_reason,omitempty"`
	SystemLocation             string   `json:"system_location"`
	AlternativeFQDN            string   `json:"alternative_fqdn,omitempty"`
	EnableDistributedDatabase  bool     `json:"enable_distributed_database,omitempty"`
	EnableSSH                  string   `json:"enable_ssh,omitempty"`
	IPv6Address                *string  `json:"ipv6_address,omitempty"`
	IPv6Gateway                *string  `json:"ipv6_gateway,omitempty"`
	TLSCertificate             *string  `json:"tls_certificate,omitempty"`
	SecondaryAddress           *string  `json:"secondary_address,omitempty"`
	SecondaryNetmask           *string  `json:"secondary_netmask,omitempty"`
	MediaPriorityWeight        *int     `json:"media_priority_weight,omitempty"`
	SSHAuthorizedKeys          []string `json:"ssh_authorized_keys,omitempty"`
	SSHAuthorizedKeysUseCloud  bool     `json:"ssh_authorized_keys_use_cloud,omitempty"`
	StaticNATAddress           *string  `json:"static_nat_address,omitempty"`
	StaticRoutes               []string `json:"static_routes,omitempty"`
	SNMPAuthenticationPassword string   `json:"snmp_authentication_password,omitempty"`
	SNMPCommunity              string   `json:"snmp_community,omitempty"`
	SNMPMode                   string   `json:"snmp_mode,omitempty"`
	SNMPPrivacyPassword        string   `json:"snmp_privacy_password,omitempty"`
	SNMPSystemContact          string   `json:"snmp_system_contact,omitempty"`
	SNMPSystemLocation         string   `json:"snmp_system_location,omitempty"`
	SNMPUsername               string   `json:"snmp_username,omitempty"`
	ServiceManager             bool     `json:"service_manager,omitempty"`
	ServicePolicy              bool     `json:"service_policy,omitempty"`
	Signalling                 bool     `json:"signalling,omitempty"`
	Managed                    bool     `json:"managed,omitempty"`
	ResourceURI                string   `json:"resource_uri,omitempty"`
}

// WorkerVMCreateRequest represents a request to create a worker VM
type WorkerVMCreateRequest struct {
	ID                         int      `json:"id,omitempty"`
	Name                       string   `json:"name"`
	Description                string   `json:"description,omitempty"`
	Hostname                   string   `json:"hostname"`
	Domain                     string   `json:"domain"`
	Address                    string   `json:"address"`
	Netmask                    string   `json:"netmask"`
	Gateway                    string   `json:"gateway"`
	VMCPUCount                 int      `json:"vm_cpu_count,omitempty"`
	VMSystemMemory             int      `json:"vm_system_memory,omitempty"`
	DeploymentType             string   `json:"deployment_type,omitempty"`
	NodeType                   string   `json:"node_type,omitempty"`
	Transcoding                bool     `json:"transcoding,omitempty"`
	CloudBursting              bool     `json:"cloud_bursting,omitempty"`
	Password                   string   `json:"password,omitempty"`
	MaintenanceMode            bool     `json:"maintenance_mode,omitempty"`
	MaintenanceModeReason      string   `json:"maintenance_mode_reason,omitempty"`
	SystemLocation             string   `json:"system_location"`
	AlternativeFQDN            string   `json:"alternative_fqdn,omitempty"`
	EnableDistributedDatabase  bool     `json:"enable_distributed_database"`
	EnableSSH                  string   `json:"enable_ssh,omitempty"`
	IPv6Address                *string  `json:"ipv6_address,omitempty"`
	IPv6Gateway                *string  `json:"ipv6_gateway,omitempty"`
	TLSCertificate             *string  `json:"tls_certificate,omitempty"`
	SecondaryAddress           *string  `json:"secondary_address,omitempty"`
	SecondaryNetmask           *string  `json:"secondary_netmask,omitempty"`
	MediaPriorityWeight        *int     `json:"media_priority_weight,omitempty"`
	SSHAuthorizedKeys          []string `json:"ssh_authorized_keys,omitempty"`
	SSHAuthorizedKeysUseCloud  bool     `json:"ssh_authorized_keys_use_cloud,omitempty"`
	StaticNATAddress           *string  `json:"static_nat_address,omitempty"`
	StaticRoutes               []string `json:"static_routes,omitempty"`
	SNMPAuthenticationPassword string   `json:"snmp_authentication_password,omitempty"`
	SNMPCommunity              string   `json:"snmp_community,omitempty"`
	SNMPMode                   string   `json:"snmp_mode,omitempty"`
	SNMPPrivacyPassword        string   `json:"snmp_privacy_password,omitempty"`
	SNMPSystemContact          string   `json:"snmp_system_contact,omitempty"`
	SNMPSystemLocation         string   `json:"snmp_system_location,omitempty"`
	SNMPUsername               string   `json:"snmp_username,omitempty"`
	ServiceManager             bool     `json:"service_manager,omitempty"`
	ServicePolicy              bool     `json:"service_policy,omitempty"`
	Signalling                 bool     `json:"signalling,omitempty"`
	Managed                    bool     `json:"managed,omitempty"`
}

// WorkerVMUpdateRequest represents a request to update a worker VM
type WorkerVMUpdateRequest struct {
	ID                         int      `json:"id,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Description                string   `json:"description,omitempty"`
	Hostname                   string   `json:"hostname,omitempty"`
	Domain                     string   `json:"domain,omitempty"`
	Address                    string   `json:"address,omitempty"`
	Netmask                    string   `json:"netmask,omitempty"`
	Gateway                    string   `json:"gateway,omitempty"`
	VMCPUCount                 int      `json:"vm_cpu_count"`
	VMSystemMemory             int      `json:"vm_system_memory"`
	DeploymentType             string   `json:"deployment_type"`
	NodeType                   string   `json:"node_type"`
	Transcoding                bool     `json:"transcoding"`
	CloudBursting              bool     `json:"cloud_bursting"`
	Password                   string   `json:"password"`
	MaintenanceMode            bool     `json:"maintenance_mode"`
	MaintenanceModeReason      string   `json:"maintenance_mode_reason"`
	SystemLocation             string   `json:"system_location"`
	AlternativeFQDN            string   `json:"alternative_fqdn"`
	EnableDistributedDatabase  bool     `json:"enable_distributed_database"`
	EnableSSH                  string   `json:"enable_ssh"`
	IPv6Address                *string  `json:"ipv6_address"`
	IPv6Gateway                *string  `json:"ipv6_gateway"`
	TLSCertificate             *string  `json:"tls_certificate"`
	SecondaryAddress           *string  `json:"secondary_address"`
	SecondaryNetmask           *string  `json:"secondary_netmask"`
	MediaPriorityWeight        *int     `json:"media_priority_weight"`
	SSHAuthorizedKeys          []string `json:"ssh_authorized_keys"`
	SSHAuthorizedKeysUseCloud  bool     `json:"ssh_authorized_keys_use_cloud"`
	StaticNATAddress           *string  `json:"static_nat_address"`
	StaticRoutes               []string `json:"static_routes"`
	SNMPAuthenticationPassword string   `json:"snmp_authentication_password"`
	SNMPCommunity              string   `json:"snmp_community"`
	SNMPMode                   string   `json:"snmp_mode"`
	SNMPPrivacyPassword        string   `json:"snmp_privacy_password"`
	SNMPSystemContact          string   `json:"snmp_system_contact"`
	SNMPSystemLocation         string   `json:"snmp_system_location"`
	SNMPUsername               string   `json:"snmp_username"`
	ServiceManager             bool     `json:"service_manager,omitempty"`
	ServicePolicy              bool     `json:"service_policy,omitempty"`
	Signalling                 bool     `json:"signalling,omitempty"`
	Managed                    bool     `json:"managed,omitempty"`
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
