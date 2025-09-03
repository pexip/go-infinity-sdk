/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package history

import (
	"context"
	"fmt"
)

// ListAlarms retrieves a list of alarm history records
func (s *Service) ListAlarms(ctx context.Context, opts *ListOptions) (*AlarmListResponse, error) {
	endpoint := "history/v1/alarm/"

	var result AlarmListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetAlarm retrieves a specific alarm history record by ID
func (s *Service) GetAlarm(ctx context.Context, id int) (*Alarm, error) {
	endpoint := fmt.Sprintf("history/v1/alarm/%d/", id)

	var result Alarm
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
