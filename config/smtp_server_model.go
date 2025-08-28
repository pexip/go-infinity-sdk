/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// SMTPServer represents an SMTP server configuration
type SMTPServer struct {
	ID                 int    `json:"id,omitempty"`
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	Address            string `json:"address"`
	Port               int    `json:"port"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	FromEmailAddress   string `json:"from_email_address"`
	ConnectionSecurity string `json:"connection_security"`
	ResourceURI        string `json:"resource_uri,omitempty"`
}

// SMTPServerCreateRequest represents a request to create an SMTP server
type SMTPServerCreateRequest struct {
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	Address            string `json:"address"`
	Port               int    `json:"port"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	FromEmailAddress   string `json:"from_email_address"`
	ConnectionSecurity string `json:"connection_security"`
}

// SMTPServerUpdateRequest represents a request to update an SMTP server
type SMTPServerUpdateRequest struct {
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	Address            string `json:"address,omitempty"`
	Port               *int   `json:"port,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	FromEmailAddress   string `json:"from_email_address,omitempty"`
	ConnectionSecurity string `json:"connection_security,omitempty"`
}

// SMTPServerListResponse represents the response from listing SMTP servers
type SMTPServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SMTPServer `json:"objects"`
}
