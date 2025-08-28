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

func TestService_ListTURNServers(t *testing.T) {
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
				port1 := 3478
				port2 := 5349
				expectedResponse := &TURNServerListResponse{
					Objects: []TURNServer{
						{ID: 1, Name: "primary-turn", Description: "Primary TURN server", Address: "turn.example.com", Port: &port1, Username: "turnuser", ServerType: "external", TransportType: "udp"},
						{ID: 2, Name: "secure-turn", Description: "Secure TURN server", Address: "turns.example.com", Port: &port2, SecretKey: "secret123", ServerType: "external", TransportType: "tls"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/turn_server/", mock.AnythingOfType("*config.TURNServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*TURNServerListResponse)
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
				port := 3478
				expectedResponse := &TURNServerListResponse{
					Objects: []TURNServer{
						{ID: 1, Name: "primary-turn", Description: "Primary TURN server", Address: "turn.example.com", Port: &port, Username: "turnuser", ServerType: "external", TransportType: "udp"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/turn_server/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.TURNServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*TURNServerListResponse)
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
			result, err := service.ListTURNServers(t.Context(), tt.opts)

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

func TestService_GetTURNServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	port := 3478
	expectedTURNServer := &TURNServer{
		ID:            1,
		Name:          "test-turn",
		Description:   "Test TURN server",
		Address:       "turn.example.com",
		Port:          &port,
		Username:      "testuser",
		Password:      "testpass",
		SecretKey:     "testsecret",
		ServerType:    "external",
		TransportType: "udp",
		ResourceURI:   "/api/admin/configuration/v1/turn_server/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/turn_server/1/", mock.AnythingOfType("*config.TURNServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*TURNServer)
		*result = *expectedTURNServer
	})

	service := New(client)
	result, err := service.GetTURNServer(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedTURNServer, result)
	client.AssertExpectations(t)
}

func TestService_CreateTURNServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	port := 3478
	createRequest := &TURNServerCreateRequest{
		Name:          "new-turn",
		Description:   "New TURN server",
		Address:       "new-turn.example.com",
		Port:          &port,
		Username:      "newuser",
		Password:      "newpass",
		ServerType:    "external",
		TransportType: "udp",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/turn_server/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/turn_server/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateTURNServer(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateTURNServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	port := 5349
	updateRequest := &TURNServerUpdateRequest{
		Description:   "Updated TURN server",
		Port:          &port,
		TransportType: "tls",
	}

	expectedTURNServer := &TURNServer{
		ID:            1,
		Name:          "test-turn",
		Description:   "Updated TURN server",
		Address:       "turn.example.com",
		Port:          &port,
		Username:      "testuser",
		Password:      "testpass",
		ServerType:    "external",
		TransportType: "tls",
		ResourceURI:   "/api/admin/configuration/v1/turn_server/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/turn_server/1/", updateRequest, mock.AnythingOfType("*config.TURNServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TURNServer)
		*result = *expectedTURNServer
	})

	service := New(client)
	result, err := service.UpdateTURNServer(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedTURNServer, result)
	client.AssertExpectations(t)
}

func TestService_DeleteTURNServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/turn_server/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteTURNServer(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
