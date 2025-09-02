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

// ListLdapSyncSources retrieves a list of LDAP sync sources
func (s *Service) ListLdapSyncSources(ctx context.Context, opts *ListOptions) (*LdapSyncSourceListResponse, error) {
	endpoint := "configuration/v1/ldap_sync_source/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result LdapSyncSourceListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetLdapSyncSource retrieves a specific LDAP sync source by ID
func (s *Service) GetLdapSyncSource(ctx context.Context, id int) (*LdapSyncSource, error) {
	endpoint := fmt.Sprintf("configuration/v1/ldap_sync_source/%d/", id)

	var result LdapSyncSource
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateLdapSyncSource creates a new LDAP sync source
func (s *Service) CreateLdapSyncSource(ctx context.Context, req *LdapSyncSourceCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/ldap_sync_source/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateLdapSyncSource updates an existing LDAP sync source
func (s *Service) UpdateLdapSyncSource(ctx context.Context, id int, req *LdapSyncSourceUpdateRequest) (*LdapSyncSource, error) {
	endpoint := fmt.Sprintf("configuration/v1/ldap_sync_source/%d/", id)

	var result LdapSyncSource
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteLdapSyncSource deletes an LDAP sync source
func (s *Service) DeleteLdapSyncSource(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/ldap_sync_source/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
