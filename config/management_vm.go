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

// ListManagementVMs retrieves a list of management VMs
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

// GetManagementVM retrieves a specific management VM by ID
func (s *Service) GetManagementVM(ctx context.Context, id int) (*ManagementVM, error) {
	endpoint := fmt.Sprintf("configuration/v1/management_vm/%d/", id)

	var result ManagementVM
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// CreateManagementVM creates a new management VM
func (s *Service) CreateManagementVM(ctx context.Context, req *ManagementVMCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/management_vm/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateManagementVM updates an existing management VM
func (s *Service) UpdateManagementVM(ctx context.Context, id int, req *ManagementVMUpdateRequest) (*ManagementVM, error) {
	endpoint := fmt.Sprintf("configuration/v1/management_vm/%d/", id)

	var result ManagementVM
	err := s.client.PatchJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteManagementVM deletes a management VM
func (s *Service) DeleteManagementVM(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/management_vm/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
