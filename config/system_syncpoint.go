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

// CreateSystemSyncpoint creates a new system syncpoint
func (s *Service) CreateSystemSyncpoint(ctx context.Context, req *SystemSyncpointCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/system_syncpoint/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// GetSystemSyncpoint retrieves a specific system syncpoint by ID (read-only)
func (s *Service) GetSystemSyncpoint(ctx context.Context, id int) (*SystemSyncpoint, error) {
	endpoint := fmt.Sprintf("configuration/v1/system_syncpoint/%d/", id)

	var result SystemSyncpoint
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
