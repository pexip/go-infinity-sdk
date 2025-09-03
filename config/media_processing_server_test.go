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

func TestService_ListMediaProcessingServers(t *testing.T) {
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
				appID1 := "media-app-01"
				appID2 := "media-app-02"
				expectedResponse := &MediaProcessingServerListResponse{
					Objects: []MediaProcessingServer{
						{ID: 1, FQDN: "media1.example.com", AppID: &appID1, PublicJWTKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----"},
						{ID: 2, FQDN: "media2.example.com", AppID: &appID2, PublicJWTKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/media_processing_server/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MediaProcessingServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MediaProcessingServerListResponse)
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
				Search: "media1",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				appID := "media-app-01"
				expectedResponse := &MediaProcessingServerListResponse{
					Objects: []MediaProcessingServer{
						{ID: 1, FQDN: "media1.example.com", AppID: &appID, PublicJWTKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/media_processing_server/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MediaProcessingServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MediaProcessingServerListResponse)
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
			result, err := service.ListMediaProcessingServers(t.Context(), tt.opts)

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

func TestService_GetMediaProcessingServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	appID := "test-media-app"
	expectedMediaProcessingServer := &MediaProcessingServer{
		ID:           1,
		FQDN:         "test-media.example.com",
		AppID:        &appID,
		PublicJWTKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtest...\n-----END PUBLIC KEY-----",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/media_processing_server/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MediaProcessingServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaProcessingServer)
		*result = *expectedMediaProcessingServer
	})

	service := New(client)
	result, err := service.GetMediaProcessingServer(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedMediaProcessingServer, result)
	client.AssertExpectations(t)
}

func TestService_CreateMediaProcessingServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	appID := "new-media-app"
	createRequest := &MediaProcessingServerCreateRequest{
		FQDN:         "new-media.example.com",
		AppID:        &appID,
		PublicJWTKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnew...\n-----END PUBLIC KEY-----",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/media_processing_server/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/media_processing_server/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMediaProcessingServer(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMediaProcessingServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &MediaProcessingServerUpdateRequest{
		PublicJWTKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAupdated...\n-----END PUBLIC KEY-----",
	}

	appID := "test-media-app"
	expectedMediaProcessingServer := &MediaProcessingServer{
		ID:           1,
		FQDN:         "test-media.example.com",
		AppID:        &appID,
		PublicJWTKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAupdated...\n-----END PUBLIC KEY-----",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/media_processing_server/1/", updateRequest, mock.AnythingOfType("*config.MediaProcessingServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaProcessingServer)
		*result = *expectedMediaProcessingServer
	})

	service := New(client)
	result, err := service.UpdateMediaProcessingServer(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedMediaProcessingServer, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMediaProcessingServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/media_processing_server/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMediaProcessingServer(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
