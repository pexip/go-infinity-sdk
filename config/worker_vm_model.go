package config

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
