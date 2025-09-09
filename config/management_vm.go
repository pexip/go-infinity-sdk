/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
)

// GetManagementVM retrieves a specific management VM by ID
func (s *Service) GetManagementVM(ctx context.Context) (*ManagementVM, error) {
	endpoint := "configuration/v1/management_vm/1/"

	var result ManagementVM
	err := s.client.GetJSON(ctx, endpoint, nil, &result)
	return &result, err
}

// UpdateManagementVM updates an existing management VM
func (s *Service) UpdateManagementVM(ctx context.Context, req *ManagementVMUpdateRequest) (*ManagementVM, error) {
	endpoint := "configuration/v1/management_vm/1/"

	var result ManagementVM
	err := s.client.PatchJSON(ctx, endpoint, req, &result)
	return &result, err
}
