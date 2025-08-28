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

// ListManagementVMs retrieves a list of management VM statuses
func (s *Service) ListManagementVMs(ctx context.Context, opts *ListOptions) (*ManagementVMListResponse, error) {
	endpoint := "status/v1/management_vm/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ManagementVMListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetManagementVM retrieves a specific management VM status by ID
func (s *Service) GetManagementVM(ctx context.Context, id int) (*ManagementVM, error) {
	endpoint := fmt.Sprintf("status/v1/management_vm/%d/", id)

	var result ManagementVM
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
