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

func TestService_ListGoogleAuthServers(t *testing.T) {
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
				clientID1 := "123456789.apps.googleusercontent.com"
				clientID2 := "987654321.apps.googleusercontent.com"
				expectedResponse := &GoogleAuthServerListResponse{
					Objects: []GoogleAuthServer{
						{ID: 1, Name: "primary-google-auth", Description: "Primary Google OAuth server", ApplicationType: "web", ClientID: &clientID1, ClientSecret: "secret1"},
						{ID: 2, Name: "backup-google-auth", Description: "Backup Google OAuth server", ApplicationType: "web", ClientID: &clientID2, ClientSecret: "secret2"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/google_auth_server/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.GoogleAuthServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*GoogleAuthServerListResponse)
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
				clientID := "123456789.apps.googleusercontent.com"
				expectedResponse := &GoogleAuthServerListResponse{
					Objects: []GoogleAuthServer{
						{ID: 1, Name: "primary-google-auth", Description: "Primary Google OAuth server", ApplicationType: "web", ClientID: &clientID, ClientSecret: "secret1"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/google_auth_server/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.GoogleAuthServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*GoogleAuthServerListResponse)
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
			result, err := service.ListGoogleAuthServers(t.Context(), tt.opts)

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

func TestService_GetGoogleAuthServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	clientID := "123456789.apps.googleusercontent.com"
	expectedServer := &GoogleAuthServer{
		ID:              1,
		Name:            "test-google-auth",
		Description:     "Test Google OAuth server",
		ApplicationType: "web",
		ClientID:        &clientID,
		ClientSecret:    "test-secret",
		ResourceURI:     "/api/admin/configuration/v1/google_auth_server/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/google_auth_server/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.GoogleAuthServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*GoogleAuthServer)
		*result = *expectedServer
	})

	service := New(client)
	result, err := service.GetGoogleAuthServer(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedServer, result)
	client.AssertExpectations(t)
}

func TestService_CreateGoogleAuthServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	clientID := "new123456789.apps.googleusercontent.com"
	createRequest := &GoogleAuthServerCreateRequest{
		Name:            "new-google-auth",
		Description:     "New Google OAuth server",
		ApplicationType: "web",
		ClientID:        &clientID,
		ClientSecret:    "new-secret",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/google_auth_server/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/google_auth_server/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateGoogleAuthServer(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateGoogleAuthServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &GoogleAuthServerUpdateRequest{
		Name:            "updated-google-auth",
		Description:     "Updated Google OAuth server",
		ApplicationType: "web",
	}

	clientID := "123456789.apps.googleusercontent.com"
	expectedServer := &GoogleAuthServer{
		ID:              1,
		Name:            "updated-google-auth",
		Description:     "Updated Google OAuth server",
		ApplicationType: "web",
		ClientID:        &clientID,
		ClientSecret:    "test-secret",
		ResourceURI:     "/api/admin/configuration/v1/google_auth_server/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/google_auth_server/1/", updateRequest, mock.AnythingOfType("*config.GoogleAuthServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*GoogleAuthServer)
		*result = *expectedServer
	})

	service := New(client)
	result, err := service.UpdateGoogleAuthServer(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedServer, result)
	client.AssertExpectations(t)
}

func TestService_DeleteGoogleAuthServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/google_auth_server/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteGoogleAuthServer(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
