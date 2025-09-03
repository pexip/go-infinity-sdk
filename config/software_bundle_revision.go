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

// ListSoftwareBundleRevisions retrieves a list of software bundle revisions (read-only)
func (s *Service) ListSoftwareBundleRevisions(ctx context.Context, opts *ListOptions) (*SoftwareBundleRevisionListResponse, error) {
	endpoint := "configuration/v1/software_bundle_revision/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result SoftwareBundleRevisionListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetSoftwareBundleRevision retrieves a specific software bundle revision by ID (read-only)
func (s *Service) GetSoftwareBundleRevision(ctx context.Context, id int) (*SoftwareBundleRevision, error) {
	endpoint := fmt.Sprintf("configuration/v1/software_bundle_revision/%d/", id)

	var result SoftwareBundleRevision
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
