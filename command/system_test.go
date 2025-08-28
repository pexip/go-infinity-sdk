/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package command

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_CreateSnapshot(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	includeDiagnostics := true
	expectedRequest := &SnapshotRequest{
		Limit:                    intPtr(100),
		IncludeDiagnosticMetrics: &includeDiagnostics,
		Request:                  boolPtr(false),
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Snapshot created successfully",
	}

	client.On("PostJSON", t.Context(), "command/v1/snapshot/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.CreateSnapshot(t.Context(), expectedRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_CreateSnapshotSimple(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Snapshot created successfully",
	}

	client.On("PostJSON", t.Context(), "command/v1/snapshot/", mock.AnythingOfType("*command.SnapshotRequest"), mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.CreateSnapshotSimple(t.Context())

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_ImportCertificates(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &CertificatesImportRequest{
		Bundle:               "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
		PrivateKeyPassphrase: "cert-passphrase",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Certificates imported successfully",
	}

	client.On("PostJSON", t.Context(), "command/v1/certificates/import/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ImportCertificates(t.Context(), "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----", "cert-passphrase")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_ManageSoftwareBundle(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &SoftwareBundleRequest{
		Package: "software-bundle.zip",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Software bundle managed successfully",
	}

	client.On("PostJSON", t.Context(), "command/v1/software/bundle/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ManageSoftwareBundle(t.Context(), "software-bundle.zip")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpgradeSystem(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &UpgradeRequest{
		Package: "upgrade-package.zip",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "System upgrade initiated",
	}

	client.On("PostJSON", t.Context(), "command/v1/upgrade/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.UpgradeSystem(t.Context(), "upgrade-package.zip")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_StartCloudNode(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &StartCloudNodeRequest{
		InstanceID: "cloud-instance-123",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Cloud node started successfully",
	}

	client.On("PostJSON", t.Context(), "command/v1/cloudnode/start/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.StartCloudNode(t.Context(), "cloud-instance-123")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_Sync(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &SyncRequest{
		ConferenceSyncTemplateID: "sync-template-456",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Sync completed successfully",
	}

	client.On("PostJSON", t.Context(), "command/v1/sync/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.Sync(t.Context(), "sync-template-456")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}
