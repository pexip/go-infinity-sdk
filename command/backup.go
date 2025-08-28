/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package command

import (
	"context"
)

// CreateBackup creates a new system backup
func (s *Service) CreateBackup(ctx context.Context, passphrase string, request bool) (*CommandResponse, error) {
	endpoint := "command/v1/backup/create/"

	req := &BackupCreateRequest{
		Passphrase: passphrase,
		Request:    request,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// RestoreBackup restores from a backup file
func (s *Service) RestoreBackup(ctx context.Context, packageName, passphrase string) (*CommandResponse, error) {
	endpoint := "command/v1/backup/restore/"

	req := &BackupRestoreRequest{
		Package:    packageName,
		Passphrase: passphrase,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}
