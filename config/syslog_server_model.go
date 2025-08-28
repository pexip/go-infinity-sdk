/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// SyslogServer represents a syslog server configuration
type SyslogServer struct {
	ID          int    `json:"id,omitempty"`
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
	Port        int    `json:"port"`
	Transport   string `json:"transport"`
	AuditLog    bool   `json:"audit_log"`
	SupportLog  bool   `json:"support_log"`
	WebLog      bool   `json:"web_log"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// SyslogServerCreateRequest represents a request to create a syslog server
type SyslogServerCreateRequest struct {
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
	Port        int    `json:"port"`
	Transport   string `json:"transport"`
	AuditLog    bool   `json:"audit_log"`
	SupportLog  bool   `json:"support_log"`
	WebLog      bool   `json:"web_log"`
}

// SyslogServerUpdateRequest represents a request to update a syslog server
type SyslogServerUpdateRequest struct {
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
	Port        int    `json:"port,omitempty"`
	Transport   string `json:"transport,omitempty"`
	AuditLog    *bool  `json:"audit_log,omitempty"`
	SupportLog  *bool  `json:"support_log,omitempty"`
	WebLog      *bool  `json:"web_log,omitempty"`
}

// SyslogServerListResponse represents the response from listing syslog servers
type SyslogServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SyslogServer `json:"objects"`
}
