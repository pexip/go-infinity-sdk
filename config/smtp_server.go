/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListSMTPServers retrieves a list of SMTP servers
func (s *Service) ListSMTPServers(ctx context.Context, opts *ListOptions) (*SMTPServerListResponse, error) {
	endpoint := "configuration/v1/smtp_server/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result SMTPServerListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetSMTPServer retrieves a specific SMTP server by ID
func (s *Service) GetSMTPServer(ctx context.Context, id int) (*SMTPServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/smtp_server/%d/", id)

	var result SMTPServer
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateSMTPServer creates a new SMTP server
func (s *Service) CreateSMTPServer(ctx context.Context, req *SMTPServerCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/smtp_server/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateSMTPServer updates an existing SMTP server
func (s *Service) UpdateSMTPServer(ctx context.Context, id int, req *SMTPServerUpdateRequest) (*SMTPServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/smtp_server/%d/", id)

	var result SMTPServer
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteSMTPServer deletes an SMTP server
func (s *Service) DeleteSMTPServer(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/smtp_server/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
