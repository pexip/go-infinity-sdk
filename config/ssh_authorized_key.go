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

// ListSSHAuthorizedKeys retrieves a list of SSH authorized keys
func (s *Service) ListSSHAuthorizedKeys(ctx context.Context, opts *ListOptions) (*SSHAuthorizedKeyListResponse, error) {
	endpoint := "configuration/v1/ssh_authorized_key/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result SSHAuthorizedKeyListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetSSHAuthorizedKey retrieves a specific SSH authorized key by ID
func (s *Service) GetSSHAuthorizedKey(ctx context.Context, id int) (*SSHAuthorizedKey, error) {
	endpoint := fmt.Sprintf("configuration/v1/ssh_authorized_key/%d/", id)

	var result SSHAuthorizedKey
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateSSHAuthorizedKey creates a new SSH authorized key
func (s *Service) CreateSSHAuthorizedKey(ctx context.Context, req *SSHAuthorizedKeyCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/ssh_authorized_key/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateSSHAuthorizedKey updates an existing SSH authorized key
func (s *Service) UpdateSSHAuthorizedKey(ctx context.Context, id int, req *SSHAuthorizedKeyUpdateRequest) (*SSHAuthorizedKey, error) {
	endpoint := fmt.Sprintf("configuration/v1/ssh_authorized_key/%d/", id)

	var result SSHAuthorizedKey
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteSSHAuthorizedKey deletes an SSH authorized key
func (s *Service) DeleteSSHAuthorizedKey(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/ssh_authorized_key/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
