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
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListADFSAuthServers(t *testing.T) {
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
				expectedResponse := &ADFSAuthServerListResponse{
					Objects: []ADFSAuthServer{
						{ID: 1, Name: "primary-adfs", Description: "Primary ADFS server", ClientID: "client1", FederationServiceName: "fs.example.com", FederationServiceIdentifier: "http://fs.example.com/adfs/services/trust", RelyingPartyTrustIdentifierURL: "https://example.com/"},
						{ID: 2, Name: "backup-adfs", Description: "Backup ADFS server", ClientID: "client2", FederationServiceName: "fs2.example.com", FederationServiceIdentifier: "http://fs2.example.com/adfs/services/trust", RelyingPartyTrustIdentifierURL: "https://backup.example.com/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/adfs_auth_server/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ADFSAuthServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ADFSAuthServerListResponse)
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
				Search: "primary",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &ADFSAuthServerListResponse{
					Objects: []ADFSAuthServer{
						{ID: 1, Name: "primary-adfs", Description: "Primary ADFS server", ClientID: "client1", FederationServiceName: "fs.example.com", FederationServiceIdentifier: "http://fs.example.com/adfs/services/trust", RelyingPartyTrustIdentifierURL: "https://example.com/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/adfs_auth_server/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ADFSAuthServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ADFSAuthServerListResponse)
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
			result, err := service.ListADFSAuthServers(t.Context(), tt.opts)

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

func TestService_GetADFSAuthServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedADFSAuthServer := &ADFSAuthServer{
		ID:                             1,
		Name:                           "test-adfs",
		Description:                    "Test ADFS server",
		ClientID:                       "test-client",
		FederationServiceName:          "fs.example.com",
		FederationServiceIdentifier:    "http://fs.example.com/adfs/services/trust",
		RelyingPartyTrustIdentifierURL: "https://example.com/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/adfs_auth_server/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ADFSAuthServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ADFSAuthServer)
		*result = *expectedADFSAuthServer
	})

	service := New(client)
	result, err := service.GetADFSAuthServer(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedADFSAuthServer, result)
	client.AssertExpectations(t)
}

func TestService_CreateADFSAuthServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &ADFSAuthServerCreateRequest{
		Name:                           "new-adfs",
		Description:                    "New ADFS server",
		ClientID:                       "new-client",
		FederationServiceName:          "new-fs.example.com",
		FederationServiceIdentifier:    "http://new-fs.example.com/adfs/services/trust",
		RelyingPartyTrustIdentifierURL: "https://new.example.com/",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/adfs_auth_server/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/adfs_auth_server/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateADFSAuthServer(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateADFSAuthServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &ADFSAuthServerUpdateRequest{
		Description:           "Updated ADFS server",
		FederationServiceName: "updated-fs.example.com",
	}

	expectedADFSAuthServer := &ADFSAuthServer{
		ID:                             1,
		Name:                           "test-adfs",
		Description:                    "Updated ADFS server",
		ClientID:                       "test-client",
		FederationServiceName:          "updated-fs.example.com",
		FederationServiceIdentifier:    "http://fs.example.com/adfs/services/trust",
		RelyingPartyTrustIdentifierURL: "https://example.com/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/adfs_auth_server/1/", updateRequest, mock.AnythingOfType("*config.ADFSAuthServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ADFSAuthServer)
		*result = *expectedADFSAuthServer
	})

	service := New(client)
	result, err := service.UpdateADFSAuthServer(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedADFSAuthServer, result)
	client.AssertExpectations(t)
}

func TestService_DeleteADFSAuthServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/adfs_auth_server/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteADFSAuthServer(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
