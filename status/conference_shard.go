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

// ListConferenceShards retrieves a list of conference shard statuses
func (s *Service) ListConferenceShards(ctx context.Context, opts *ListOptions) (*ConferenceShardListResponse, error) {
	endpoint := "status/v1/conference_shard/"

	var result ConferenceShardListResponse
	err := s.listEndpoint(ctx, endpoint, opts, &result)
	return &result, err
}

// GetConferenceShard retrieves a specific conference shard status by ID
func (s *Service) GetConferenceShard(ctx context.Context, id string) (*ConferenceShard, error) {
	endpoint := fmt.Sprintf("status/v1/conference_shard/%s/", id)

	var result ConferenceShard
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
