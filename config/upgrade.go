/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// CreateUpgrade initiates a system upgrade (POST only)
func (s *Service) CreateUpgrade(ctx context.Context, req *UpgradeCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/upgrade/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}
