/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
)

// GetAutobackup retrieves the autobackup configuration (singleton resource)
func (s *Service) GetAutobackup(ctx context.Context) (*Autobackup, error) {
	endpoint := "configuration/v1/autobackup/1/"

	var result Autobackup
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// UpdateAutobackup updates the autobackup configuration (singleton resource)
func (s *Service) UpdateAutobackup(ctx context.Context, req *AutobackupUpdateRequest) (*Autobackup, error) {
	endpoint := "configuration/v1/autobackup/1/"

	var result Autobackup
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}
