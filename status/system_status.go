/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"context"
)

// GetSystemStatus retrieves the overall system status
func (s *Service) GetSystemStatus(ctx context.Context) (*SystemStatus, error) {
	endpoint := "status/v1/system_status/"

	var result SystemStatus
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
