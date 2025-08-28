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

// ListLicences retrieves a list of licences
func (s *Service) ListLicences(ctx context.Context, opts *ListOptions) (*LicenceListResponse, error) {
	endpoint := "configuration/v1/licence/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result LicenceListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetLicence retrieves a specific licence by fulfillment ID
func (s *Service) GetLicence(ctx context.Context, fulfillmentID string) (*Licence, error) {
	endpoint := fmt.Sprintf("configuration/v1/licence/%s/", fulfillmentID)

	var result Licence
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateLicence creates a new licence (activates a licence)
func (s *Service) CreateLicence(ctx context.Context, req *LicenceCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/licence/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// DeleteLicence deletes a licence
func (s *Service) DeleteLicence(ctx context.Context, fulfillmentID string) error {
	endpoint := fmt.Sprintf("configuration/v1/licence/%s/", fulfillmentID)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
