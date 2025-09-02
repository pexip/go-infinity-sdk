/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_GetAutobackup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedAutobackup := &Autobackup{
		ID:                       1,
		AutobackupEnabled:        true,
		AutobackupInterval:       24,
		AutobackupPassphrase:     "test-passphrase",
		AutobackupStartHour:      2,
		AutobackupUploadURL:      "ftp://backup.example.com/",
		AutobackupUploadUsername: "backupuser",
		AutobackupUploadPassword: "backuppass",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/autobackup/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.Autobackup")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Autobackup)
		*result = *expectedAutobackup
	})

	service := New(client)
	result, err := service.GetAutobackup(t.Context())

	assert.NoError(t, err)
	assert.Equal(t, expectedAutobackup, result)
	client.AssertExpectations(t)
}

func TestService_UpdateAutobackup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	enabled := false
	interval := 48
	startHour := 3

	updateRequest := &AutobackupUpdateRequest{
		AutobackupEnabled:        &enabled,
		AutobackupInterval:       &interval,
		AutobackupStartHour:      &startHour,
		AutobackupUploadURL:      "sftp://new-backup.example.com/",
		AutobackupUploadUsername: "newuser",
	}

	expectedAutobackup := &Autobackup{
		ID:                       1,
		AutobackupEnabled:        false,
		AutobackupInterval:       48,
		AutobackupPassphrase:     "test-passphrase",
		AutobackupStartHour:      3,
		AutobackupUploadURL:      "sftp://new-backup.example.com/",
		AutobackupUploadUsername: "newuser",
		AutobackupUploadPassword: "backuppass",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/autobackup/1/", updateRequest, mock.AnythingOfType("*config.Autobackup")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Autobackup)
		*result = *expectedAutobackup
	})

	service := New(client)
	result, err := service.UpdateAutobackup(t.Context(), updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedAutobackup, result)
	client.AssertExpectations(t)
}
