/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListTLSCertificates retrieves a list of TLS certificates
func (s *Service) ListTLSCertificates(ctx context.Context, opts *ListOptions) (*TLSCertificateListResponse, error) {
	endpoint := "configuration/v1/tls_certificate/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result TLSCertificateListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetTLSCertificate retrieves a specific TLS certificate by ID
func (s *Service) GetTLSCertificate(ctx context.Context, id int) (*TLSCertificate, error) {
	endpoint := fmt.Sprintf("configuration/v1/tls_certificate/%d/", id)

	var result TLSCertificate
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateTLSCertificate creates a new TLS certificate
func (s *Service) CreateTLSCertificate(ctx context.Context, req *TLSCertificateCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/tls_certificate/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateTLSCertificate updates an existing TLS certificate (partial update)
func (s *Service) UpdateTLSCertificate(ctx context.Context, id int, req *TLSCertificateUpdateRequest) (*TLSCertificate, error) {
	endpoint := fmt.Sprintf("configuration/v1/tls_certificate/%d/", id)

	var result TLSCertificate
	err := s.client.PatchJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteTLSCertificate deletes a TLS certificate
func (s *Service) DeleteTLSCertificate(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/tls_certificate/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
