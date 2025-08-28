/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// CertificateSigningRequest represents a certificate signing request
type CertificateSigningRequest struct {
	ID                        int     `json:"id,omitempty"`
	SubjectName               string  `json:"subject_name"`
	DN                        string  `json:"dn,omitempty"`
	AdditionalSubjectAltNames string  `json:"additional_subject_alt_names,omitempty"`
	PrivateKeyType            string  `json:"private_key_type"`
	PrivateKey                *string `json:"private_key,omitempty"`
	PrivateKeyPassphrase      string  `json:"private_key_passphrase,omitempty"`
	AdCompatible              bool    `json:"ad_compatible"`
	TLSCertificate            *string `json:"tls_certificate,omitempty"`
	CSR                       string  `json:"csr,omitempty"`
	Certificate               string  `json:"certificate,omitempty"`
	ResourceURI               string  `json:"resource_uri,omitempty"`
}

// CertificateSigningRequestCreateRequest represents a request to create a certificate signing request
type CertificateSigningRequestCreateRequest struct {
	SubjectName               string  `json:"subject_name"`
	DN                        string  `json:"dn,omitempty"`
	AdditionalSubjectAltNames string  `json:"additional_subject_alt_names,omitempty"`
	PrivateKeyType            string  `json:"private_key_type"`
	PrivateKey                *string `json:"private_key,omitempty"`
	PrivateKeyPassphrase      string  `json:"private_key_passphrase,omitempty"`
	AdCompatible              bool    `json:"ad_compatible"`
	TLSCertificate            *string `json:"tls_certificate,omitempty"`
}

// CertificateSigningRequestUpdateRequest represents a request to update a certificate signing request
type CertificateSigningRequestUpdateRequest struct {
	SubjectName               string  `json:"subject_name,omitempty"`
	DN                        string  `json:"dn,omitempty"`
	AdditionalSubjectAltNames string  `json:"additional_subject_alt_names,omitempty"`
	PrivateKeyType            string  `json:"private_key_type,omitempty"`
	PrivateKey                *string `json:"private_key,omitempty"`
	PrivateKeyPassphrase      string  `json:"private_key_passphrase,omitempty"`
	AdCompatible              *bool   `json:"ad_compatible,omitempty"`
	TLSCertificate            *string `json:"tls_certificate,omitempty"`
	Certificate               string  `json:"certificate,omitempty"`
}

// CertificateSigningRequestListResponse represents the response from listing certificate signing requests
type CertificateSigningRequestListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []CertificateSigningRequest `json:"objects"`
}
