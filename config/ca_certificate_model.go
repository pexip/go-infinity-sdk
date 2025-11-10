/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// CACertificate represents a CA certificate configuration
type CACertificate struct {
	ID                  int               `json:"id,omitempty"`
	Certificate         string            `json:"certificate"`
	TrustedIntermediate bool              `json:"trusted_intermediate"`
	StartDate           util.InfinityTime `json:"start_date,omitempty"`
	EndDate             util.InfinityTime `json:"end_date,omitempty"`
	SubjectName         string            `json:"subject_name,omitempty"`
	SubjectHash         string            `json:"subject_hash,omitempty"`
	RawSubject          string            `json:"raw_subject,omitempty"`
	IssuerName          string            `json:"issuer_name,omitempty"`
	IssuerHash          string            `json:"issuer_hash,omitempty"`
	RawIssuer           string            `json:"raw_issuer,omitempty"`
	SerialNo            string            `json:"serial_no,omitempty"`
	KeyID               *string           `json:"key_id,omitempty"`
	IssuerKeyID         *string           `json:"issuer_key_id,omitempty"`
	Text                string            `json:"text,omitempty"`
	ResourceURI         string            `json:"resource_uri,omitempty"`
}

// CACertificateCreateRequest represents a request to create a CA certificate
type CACertificateCreateRequest struct {
	Certificate         string `json:"certificate"`
	TrustedIntermediate bool   `json:"trusted_intermediate"`
}

// CACertificateUpdateRequest represents a request to update a CA certificate
type CACertificateUpdateRequest struct {
	Certificate         string `json:"certificate,omitempty"`
	TrustedIntermediate bool   `json:"trusted_intermediate"`
}

// CACertificateListResponse represents the response from listing CA certificates
type CACertificateListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []CACertificate `json:"objects"`
}
