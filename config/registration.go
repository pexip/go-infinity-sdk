/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
)

// GetRegistration retrieves the registration configuration (singleton resource)
func (s *Service) GetRegistration(ctx context.Context) (*Registration, error) {
	endpoint := "configuration/v1/registration/1/"

	var result Registration
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// UpdateRegistration updates the registration configuration (singleton resource)
func (s *Service) UpdateRegistration(ctx context.Context, req *RegistrationUpdateRequest) (*Registration, error) {
	endpoint := "configuration/v1/registration/1/"

	var result Registration
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}
