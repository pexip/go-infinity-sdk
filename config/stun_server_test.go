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

func TestService_ListSTUNServers(t *testing.T) {
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
				expectedResponse := &STUNServerListResponse{
					Objects: []STUNServer{
						{ID: 1, Name: "primary-stun", Description: "Primary STUN server", Address: "stun.example.com", Port: 3478},
						{ID: 2, Name: "backup-stun", Description: "Backup STUN server", Address: "stun-backup.example.com", Port: 3478},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/stun_server/", mock.AnythingOfType("*config.STUNServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*STUNServerListResponse)
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
				expectedResponse := &STUNServerListResponse{
					Objects: []STUNServer{
						{ID: 1, Name: "primary-stun", Description: "Primary STUN server", Address: "stun.example.com", Port: 3478},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/stun_server/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.STUNServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*STUNServerListResponse)
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
			result, err := service.ListSTUNServers(t.Context(), tt.opts)

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

func TestService_GetSTUNServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedSTUNServer := &STUNServer{
		ID:          1,
		Name:        "test-stun",
		Description: "Test STUN server",
		Address:     "stun.example.com",
		Port:        3478,
		ResourceURI: "/api/admin/configuration/v1/stun_server/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/stun_server/1/", mock.AnythingOfType("*config.STUNServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*STUNServer)
		*result = *expectedSTUNServer
	})

	service := New(client)
	result, err := service.GetSTUNServer(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedSTUNServer, result)
	client.AssertExpectations(t)
}

func TestService_CreateSTUNServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &STUNServerCreateRequest{
		Name:        "new-stun",
		Description: "New STUN server",
		Address:     "new-stun.example.com",
		Port:        3478,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/stun_server/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/stun_server/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateSTUNServer(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateSTUNServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	port := 3479
	updateRequest := &STUNServerUpdateRequest{
		Description: "Updated STUN server",
		Port:        &port,
	}

	expectedSTUNServer := &STUNServer{
		ID:          1,
		Name:        "test-stun",
		Description: "Updated STUN server",
		Address:     "stun.example.com",
		Port:        3479,
		ResourceURI: "/api/admin/configuration/v1/stun_server/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/stun_server/1/", updateRequest, mock.AnythingOfType("*config.STUNServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*STUNServer)
		*result = *expectedSTUNServer
	})

	service := New(client)
	result, err := service.UpdateSTUNServer(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedSTUNServer, result)
	client.AssertExpectations(t)
}

func TestService_DeleteSTUNServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/stun_server/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteSTUNServer(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
