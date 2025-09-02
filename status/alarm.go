/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"context"
	"fmt"
)

// ListAlarms retrieves a list of system alarms
func (s *Service) ListAlarms(ctx context.Context, opts *ListOptions) (*AlarmListResponse, error) {
	endpoint := "status/v1/alarm/"

	var result AlarmListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetAlarm retrieves a specific alarm by ID
func (s *Service) GetAlarm(ctx context.Context, id int) (*Alarm, error) {
	endpoint := fmt.Sprintf("status/v1/alarm/%d/", id)

	var result Alarm
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
