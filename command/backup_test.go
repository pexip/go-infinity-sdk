/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package command

import (
	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_CreateBackup(t *testing.T) {
	tests := []struct {
		name       string
		passphrase string
		request    bool
		wantErr    bool
	}{
		{
			name:       "create backup immediately",
			passphrase: "secure-passphrase",
			request:    false,
			wantErr:    false,
		},
		{
			name:       "create backup request",
			passphrase: "another-passphrase",
			request:    true,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := interfaces.NewHTTPClientMock()

			expectedRequest := &BackupCreateRequest{
				Passphrase: tt.passphrase,
				Request:    tt.request,
			}

			expectedResponse := &CommandResponse{
				Status:  "success",
				Message: "Backup created successfully",
			}

			client.On("PostJSON", t.Context(), "command/v1/backup/create/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
				result := args.Get(3).(*CommandResponse)
				*result = *expectedResponse
			})

			service := New(client)
			result, err := service.CreateBackup(t.Context(), tt.passphrase, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, result)
			}

			client.AssertExpectations(t)
		})
	}
}

func TestService_RestoreBackup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &BackupRestoreRequest{
		Package:    "backup-file.zip",
		Passphrase: "restore-passphrase",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Backup restored successfully",
	}

	client.On("PostJSON", t.Context(), "command/v1/backup/restore/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.RestoreBackup(t.Context(), "backup-file.zip", "restore-passphrase")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}
