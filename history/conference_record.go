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

// ListConferenceRecords retrieves a list of conference history records
func (s *Service) ListConferenceRecords(ctx context.Context, opts *ListOptions) (*ConferenceRecordListResponse, error) {
	endpoint := "history/v1/conference/"

	var result ConferenceRecordListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetConferenceRecord retrieves a specific conference history record by ID
func (s *Service) GetConferenceRecord(ctx context.Context, id int) (*ConferenceRecord, error) {
	endpoint := fmt.Sprintf("history/v1/conference/%d/", id)

	var result ConferenceRecord
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
