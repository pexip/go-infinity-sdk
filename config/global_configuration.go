/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
)

// GetGlobalConfiguration retrieves the global configuration (singleton resource)
func (s *Service) GetGlobalConfiguration(ctx context.Context) (*GlobalConfiguration, error) {
	endpoint := "configuration/v1/global/1/"

	var result GlobalConfiguration
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// UpdateGlobalConfiguration updates the global configuration (singleton resource)
func (s *Service) UpdateGlobalConfiguration(ctx context.Context, req *GlobalConfigurationUpdateRequest) (*GlobalConfiguration, error) {
	endpoint := "configuration/v1/global/1/"

	var result GlobalConfiguration
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}
