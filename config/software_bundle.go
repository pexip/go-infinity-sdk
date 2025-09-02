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
)

// ListSoftwareBundles retrieves a list of software bundles (read-only)
func (s *Service) ListSoftwareBundles(ctx context.Context, opts *ListOptions) (*SoftwareBundleListResponse, error) {
	endpoint := "configuration/v1/software_bundle/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result SoftwareBundleListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetSoftwareBundle retrieves a specific software bundle by ID (read-only)
func (s *Service) GetSoftwareBundle(ctx context.Context, id int) (*SoftwareBundle, error) {
	endpoint := fmt.Sprintf("configuration/v1/software_bundle/%d/", id)

	var result SoftwareBundle
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// UpdateSoftwareBundle updates an existing software bundle (PATCH only)
func (s *Service) UpdateSoftwareBundle(ctx context.Context, id int, req *SoftwareBundleUpdateRequest) (*SoftwareBundle, error) {
	endpoint := fmt.Sprintf("configuration/v1/software_bundle/%d/", id)

	var result SoftwareBundle
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}
