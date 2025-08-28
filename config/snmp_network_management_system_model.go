/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// SnmpNetworkManagementSystem represents an SNMP network management system configuration
type SnmpNetworkManagementSystem struct {
	ID                int    `json:"id,omitempty"`
	Name              string `json:"name"`
	Description       string `json:"description,omitempty"`
	Address           string `json:"address"`
	Port              int    `json:"port"`
	SnmpTrapCommunity string `json:"snmp_trap_community"`
	ResourceURI       string `json:"resource_uri,omitempty"`
}

// SnmpNetworkManagementSystemCreateRequest represents a request to create an SNMP network management system
type SnmpNetworkManagementSystemCreateRequest struct {
	Name              string `json:"name"`
	Description       string `json:"description,omitempty"`
	Address           string `json:"address"`
	Port              int    `json:"port"`
	SnmpTrapCommunity string `json:"snmp_trap_community"`
}

// SnmpNetworkManagementSystemUpdateRequest represents a request to update an SNMP network management system
type SnmpNetworkManagementSystemUpdateRequest struct {
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	Address           string `json:"address,omitempty"`
	Port              *int   `json:"port,omitempty"`
	SnmpTrapCommunity string `json:"snmp_trap_community,omitempty"`
}

// SnmpNetworkManagementSystemListResponse represents the response from listing SNMP network management systems
type SnmpNetworkManagementSystemListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SnmpNetworkManagementSystem `json:"objects"`
}
