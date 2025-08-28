/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
)

// GetAuthentication retrieves the authentication configuration (singleton resource)
func (s *Service) GetAuthentication(ctx context.Context) (*Authentication, error) {
	endpoint := "configuration/v1/authentication/1/"

	var result Authentication
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// UpdateAuthentication updates the authentication configuration
func (s *Service) UpdateAuthentication(ctx context.Context, req *AuthenticationUpdateRequest) (*Authentication, error) {
	endpoint := "configuration/v1/authentication/1/"

	var result Authentication
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}
