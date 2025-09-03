/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListMjxMeetingProcessingRules retrieves a list of MJX meeting processing rules
func (s *Service) ListMjxMeetingProcessingRules(ctx context.Context, opts *ListOptions) (*MjxMeetingProcessingRuleListResponse, error) {
	endpoint := "configuration/v1/mjx_meeting_processing_rule/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result MjxMeetingProcessingRuleListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetMjxMeetingProcessingRule retrieves a specific MJX meeting processing rule by ID
func (s *Service) GetMjxMeetingProcessingRule(ctx context.Context, id int) (*MjxMeetingProcessingRule, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_meeting_processing_rule/%d/", id)

	var result MjxMeetingProcessingRule
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateMjxMeetingProcessingRule creates a new MJX meeting processing rule
func (s *Service) CreateMjxMeetingProcessingRule(ctx context.Context, req *MjxMeetingProcessingRuleCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/mjx_meeting_processing_rule/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateMjxMeetingProcessingRule updates an existing MJX meeting processing rule
func (s *Service) UpdateMjxMeetingProcessingRule(ctx context.Context, id int, req *MjxMeetingProcessingRuleUpdateRequest) (*MjxMeetingProcessingRule, error) {
	endpoint := fmt.Sprintf("configuration/v1/mjx_meeting_processing_rule/%d/", id)

	var result MjxMeetingProcessingRule
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteMjxMeetingProcessingRule deletes a MJX meeting processing rule
func (s *Service) DeleteMjxMeetingProcessingRule(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/mjx_meeting_processing_rule/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
