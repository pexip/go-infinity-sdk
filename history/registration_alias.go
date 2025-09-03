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

// ListRegistrationAliases retrieves a list of registration alias history records
func (s *Service) ListRegistrationAliases(ctx context.Context, opts *ListOptions) (*RegistrationAliasListResponse, error) {
	endpoint := "history/v1/registration_alias/"

	var result RegistrationAliasListResponse

	err := s.listEndpointWithSearchField(ctx, endpoint, opts, "alias__icontains", &result)
	return &result, err
}

// GetRegistrationAlias retrieves a specific registration alias history record by ID
func (s *Service) GetRegistrationAlias(ctx context.Context, id int) (*RegistrationAlias, error) {
	endpoint := fmt.Sprintf("history/v1/registration_alias/%d/", id)

	var result RegistrationAlias
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}
