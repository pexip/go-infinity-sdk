/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// TLSCertificate represents a TLS certificate configuration
type TLSCertificate struct {
	ID                   int               `json:"id,omitempty"`
	Certificate          string            `json:"certificate"`
	PrivateKey           string            `json:"private_key"`
	PrivateKeyPassphrase string            `json:"private_key_passphrase,omitempty"`
	Parameters           string            `json:"parameters,omitempty"`
	Nodes                []string          `json:"nodes,omitempty"`
	StartDate            util.InfinityTime `json:"start_date,omitempty"`
	EndDate              util.InfinityTime `json:"end_date,omitempty"`
	SubjectName          string            `json:"subject_name,omitempty"`
	SubjectHash          string            `json:"subject_hash,omitempty"`
	SubjectAltNames      string            `json:"subject_alt_names,omitempty"`
	RawSubject           string            `json:"raw_subject,omitempty"`
	IssuerName           string            `json:"issuer_name,omitempty"`
	IssuerHash           string            `json:"issuer_hash,omitempty"`
	RawIssuer            string            `json:"raw_issuer,omitempty"`
	SerialNo             string            `json:"serial_no,omitempty"`
	KeyID                *string           `json:"key_id,omitempty"`
	IssuerKeyID          *string           `json:"issuer_key_id,omitempty"`
	Text                 string            `json:"text,omitempty"`
	ResourceURI          string            `json:"resource_uri,omitempty"`
}

// TLSCertificateCreateRequest represents a request to create a TLS certificate
type TLSCertificateCreateRequest struct {
	Certificate          string   `json:"certificate"`
	PrivateKey           string   `json:"private_key"`
	PrivateKeyPassphrase string   `json:"private_key_passphrase,omitempty"`
	Parameters           string   `json:"parameters,omitempty"`
	Nodes                []string `json:"nodes,omitempty"`
}

// TLSCertificateUpdateRequest represents a request to update a TLS certificate
type TLSCertificateUpdateRequest struct {
	Certificate          string   `json:"certificate,omitempty"`
	PrivateKey           string   `json:"private_key,omitempty"`
	PrivateKeyPassphrase string   `json:"private_key_passphrase,omitempty"`
	Parameters           string   `json:"parameters,omitempty"`
	Nodes                []string `json:"nodes,omitempty"`
}

// TLSCertificateListResponse represents the response from listing TLS certificates
type TLSCertificateListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []TLSCertificate `json:"objects"`
}
