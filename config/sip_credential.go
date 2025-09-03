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

// ListSIPCredentials retrieves a list of SIP credentials
func (s *Service) ListSIPCredentials(ctx context.Context, opts *ListOptions) (*SIPCredentialListResponse, error) {
	endpoint := "configuration/v1/sip_credential/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result SIPCredentialListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetSIPCredential retrieves a specific SIP credential by ID
func (s *Service) GetSIPCredential(ctx context.Context, id int) (*SIPCredential, error) {
	endpoint := fmt.Sprintf("configuration/v1/sip_credential/%d/", id)

	var result SIPCredential
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateSIPCredential creates a new SIP credential
func (s *Service) CreateSIPCredential(ctx context.Context, req *SIPCredentialCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/sip_credential/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateSIPCredential updates an existing SIP credential
func (s *Service) UpdateSIPCredential(ctx context.Context, id int, req *SIPCredentialUpdateRequest) (*SIPCredential, error) {
	endpoint := fmt.Sprintf("configuration/v1/sip_credential/%d/", id)

	var result SIPCredential
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteSIPCredential deletes a SIP credential
func (s *Service) DeleteSIPCredential(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/sip_credential/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
