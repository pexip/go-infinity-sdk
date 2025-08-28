/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// HTTPProxy represents an HTTP proxy configuration
type HTTPProxy struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Port        *int   `json:"port,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Protocol    string `json:"protocol"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// HTTPProxyCreateRequest represents a request to create an HTTP proxy
type HTTPProxyCreateRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Port     *int   `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Protocol string `json:"protocol"`
}

// HTTPProxyUpdateRequest represents a request to update an HTTP proxy
type HTTPProxyUpdateRequest struct {
	Name     string `json:"name,omitempty"`
	Address  string `json:"address,omitempty"`
	Port     *int   `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

// HTTPProxyListResponse represents the response from listing HTTP proxies
type HTTPProxyListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []HTTPProxy `json:"objects"`
}
