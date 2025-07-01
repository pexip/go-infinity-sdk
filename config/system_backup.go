package config

import (
	"context"
)

// ListSystemBackups retrieves a list of system backups (read-only)
func (s *Service) ListSystemBackups(ctx context.Context, opts *ListOptions) (*SystemBackupListResponse, error) {
	endpoint := "configuration/v1/system_backup/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result SystemBackupListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetSystemBackup retrieves a specific system backup by filename (read-only)
func (s *Service) GetSystemBackup(ctx context.Context, filename string) (*SystemBackup, error) {
	endpoint := "configuration/v1/system_backup/" + filename + "/"

	var result SystemBackup
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// DeleteSystemBackup deletes a system backup
func (s *Service) DeleteSystemBackup(ctx context.Context, filename string) error {
	endpoint := "configuration/v1/system_backup/" + filename + "/"
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
