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

// ListLdapSyncFields retrieves a list of LDAP sync fields
func (s *Service) ListLdapSyncFields(ctx context.Context, opts *ListOptions) (*LdapSyncFieldListResponse, error) {
	endpoint := "configuration/v1/ldap_sync_field/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result LdapSyncFieldListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetLdapSyncField retrieves a specific LDAP sync field by ID
func (s *Service) GetLdapSyncField(ctx context.Context, id int) (*LdapSyncField, error) {
	endpoint := fmt.Sprintf("configuration/v1/ldap_sync_field/%d/", id)

	var result LdapSyncField
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateLdapSyncField creates a new LDAP sync field
func (s *Service) CreateLdapSyncField(ctx context.Context, req *LdapSyncFieldCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/ldap_sync_field/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateLdapSyncField updates an existing LDAP sync field
func (s *Service) UpdateLdapSyncField(ctx context.Context, id int, req *LdapSyncFieldUpdateRequest) (*LdapSyncField, error) {
	endpoint := fmt.Sprintf("configuration/v1/ldap_sync_field/%d/", id)

	var result LdapSyncField
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteLdapSyncField deletes an LDAP sync field
func (s *Service) DeleteLdapSyncField(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/ldap_sync_field/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
