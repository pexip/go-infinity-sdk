/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSystemBackups(t *testing.T) {
	tests := []struct {
		name    string
		opts    *ListOptions
		setup   func(m *interfaces.HTTPClientMock)
		wantErr bool
	}{
		{
			name: "successful list without options",
			opts: nil,
			setup: func(m *interfaces.HTTPClientMock) {
				date1 := &util.InfinityTime{}
				date2 := &util.InfinityTime{}
				expectedResponse := &SystemBackupListResponse{
					Objects: []SystemBackup{
						{Filename: "backup-20231201.tar.gz", Date: date1, Build: "27.4.1.1234", Version: "27.4.1", Size: 1048576},
						{Filename: "backup-20231202.tar.gz", Date: date2, Build: "28.0.0.1235", Version: "28.0.0", Size: 2097152},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/system_backup/", mock.AnythingOfType("*config.SystemBackupListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SystemBackupListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit: 5,
				},
				Search: "20231201",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				date := &util.InfinityTime{}
				expectedResponse := &SystemBackupListResponse{
					Objects: []SystemBackup{
						{Filename: "backup-20231201.tar.gz", Date: date, Build: "27.4.1.1234", Version: "27.4.1", Size: 1048576},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/system_backup/?limit=5&name__icontains=20231201", mock.AnythingOfType("*config.SystemBackupListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SystemBackupListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := interfaces.NewHTTPClientMock()
			tt.setup(client)

			service := New(client)
			result, err := service.ListSystemBackups(t.Context(), tt.opts)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			client.AssertExpectations(t)
		})
	}
}

func TestService_GetSystemBackup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	date := &util.InfinityTime{}
	filename := "backup-20231201.tar.gz"

	expectedSystemBackup := &SystemBackup{
		Filename:    filename,
		Date:        date,
		Build:       "27.4.1.1234",
		Version:     "27.4.1",
		Size:        1048576,
		ResourceURI: "/api/admin/configuration/v1/system_backup/" + filename + "/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/system_backup/"+filename+"/", mock.AnythingOfType("*config.SystemBackup")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemBackup)
		*result = *expectedSystemBackup
	})

	service := New(client)
	result, err := service.GetSystemBackup(t.Context(), filename)

	assert.NoError(t, err)
	assert.Equal(t, expectedSystemBackup, result)
	client.AssertExpectations(t)
}

func TestService_DeleteSystemBackup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	filename := "backup-20231201.tar.gz"

	client.On("DeleteJSON", t.Context(), "configuration/v1/system_backup/"+filename+"/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteSystemBackup(t.Context(), filename)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
