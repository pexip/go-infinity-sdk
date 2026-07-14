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

// ListManagementVMs retrieves a list of management VMs.
func (s *Service) ListManagementVMs(ctx context.Context, opts *ListOptions) (*ManagementVMListResponse, error) {
	endpoint := "configuration/v1/management_vm/"

	var params *url.Values
	if opts != nil {
		urlValues := opts.ToURLValues()
		params = &urlValues
	}

	var result ManagementVMListResponse
	err := s.client.GetJSON(ctx, endpoint, params, &result)
	return &result, err
}

// GetManagementVM retrieves a management VM by ID. If id is omitted, defaults to 1.
func (s *Service) GetManagementVM(ctx context.Context, id ...int) (*ManagementVM, error) {
	vmID := 1
	if len(id) > 0 {
		vmID = id[0]
	}
	endpoint := fmt.Sprintf("configuration/v1/management_vm/%d/", vmID)

	var result ManagementVM
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// UpdateManagementVM updates an existing management VM by ID. If id is omitted, defaults to 1.
func (s *Service) UpdateManagementVM(ctx context.Context, req *ManagementVMUpdateRequest, id ...int) (*ManagementVM, error) {
	vmID := 1
	if len(id) > 0 {
		vmID = id[0]
	}
	endpoint := fmt.Sprintf("configuration/v1/management_vm/%d/", vmID)

	var result ManagementVM
	err := s.client.PatchJSON(ctx, endpoint, req, &result)
	return &result, err
}
