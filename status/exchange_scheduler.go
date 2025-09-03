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

// ListExchangeSchedulers retrieves a list of Exchange scheduler statuses
func (s *Service) ListExchangeSchedulers(ctx context.Context, opts *ListOptions) (*ExchangeSchedulerListResponse, error) {
	endpoint := "status/v1/exchange_scheduler/"

	var result ExchangeSchedulerListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetExchangeScheduler retrieves a specific Exchange scheduler status by ID
func (s *Service) GetExchangeScheduler(ctx context.Context, id int) (*ExchangeScheduler, error) {
	endpoint := fmt.Sprintf("status/v1/exchange_scheduler/%d/", id)

	var result ExchangeScheduler
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
