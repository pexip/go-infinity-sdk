/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// EventSink represents an event sink configuration
type EventSink struct {
	ID                   int                `json:"id,omitempty"`
	Name                 string             `json:"name"`
	Description          *string            `json:"description,omitempty"`
	URL                  string             `json:"url"`
	Username             *string            `json:"username,omitempty"`
	Password             *string            `json:"password,omitempty"`
	BulkSupport          bool               `json:"bulk_support"`
	VerifyTLSCertificate bool               `json:"verify_tls_certificate"`
	Version              int                `json:"version"`
	RestartDate          *util.InfinityTime `json:"restart_date,omitempty"`
	ResourceURI          string             `json:"resource_uri,omitempty"`
}

// EventSinkCreateRequest represents a request to create an event sink
type EventSinkCreateRequest struct {
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	URL                  string  `json:"url"`
	Username             *string `json:"username,omitempty"`
	Password             *string `json:"password,omitempty"`
	BulkSupport          bool    `json:"bulk_support"`
	VerifyTLSCertificate bool    `json:"verify_tls_certificate"`
	Version              int     `json:"version"`
}

// EventSinkUpdateRequest represents a request to update an event sink
type EventSinkUpdateRequest struct {
	Name                 string  `json:"name,omitempty"`
	Description          *string `json:"description,omitempty"`
	URL                  string  `json:"url,omitempty"`
	Username             *string `json:"username,omitempty"`
	Password             *string `json:"password,omitempty"`
	BulkSupport          *bool   `json:"bulk_support,omitempty"`
	VerifyTLSCertificate *bool   `json:"verify_tls_certificate,omitempty"`
	Version              *int    `json:"version,omitempty"`
}

// EventSinkListResponse represents the response from listing event sinks
type EventSinkListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []EventSink `json:"objects"`
}
