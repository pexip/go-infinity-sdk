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

func TestService_ListNTPServers(t *testing.T) {
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
				expectedResponse := &NTPServerListResponse{
					Objects: []NTPServer{
						{ID: 1, Address: "pool.ntp.org", Description: "NTP Pool"},
						{ID: 2, Address: "time.google.com", Description: "Google Time"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ntp_server/", mock.AnythingOfType("*config.NTPServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*NTPServerListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit:  3,
					Offset: 1,
				},
				Search: "google",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &NTPServerListResponse{
					Objects: []NTPServer{
						{ID: 2, Address: "time.google.com", Description: "Google Time"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ntp_server/?limit=3&name__icontains=google&offset=1", mock.AnythingOfType("*config.NTPServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*NTPServerListResponse)
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
			result, err := service.ListNTPServers(t.Context(), tt.opts)

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

func TestService_GetNTPServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	keyID := 1
	expectedServer := &NTPServer{
		ID:          1,
		Address:     "pool.ntp.org",
		Description: "NTP Pool Server",
		Key:         "secret-key",
		KeyID:       &keyID,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/ntp_server/1/", mock.AnythingOfType("*config.NTPServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*NTPServer)
		*result = *expectedServer
	})

	service := New(client)
	result, err := service.GetNTPServer(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedServer, result)
	client.AssertExpectations(t)
}

func TestService_CreateNTPServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	keyID := 2
	createRequest := &NTPServerCreateRequest{
		Address:     "time.cloudflare.com",
		Description: "Cloudflare Time",
		Key:         "test-key",
		KeyID:       &keyID,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/ntp_server/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/ntp_server/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateNTPServer(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateNTPServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &NTPServerUpdateRequest{
		Description: "Updated NTP Server",
	}

	expectedServer := &NTPServer{
		ID:          1,
		Address:     "pool.ntp.org",
		Description: "Updated NTP Server",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/ntp_server/1/", updateRequest, mock.AnythingOfType("*config.NTPServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*NTPServer)
		*result = *expectedServer
	})

	service := New(client)
	result, err := service.UpdateNTPServer(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedServer, result)
	client.AssertExpectations(t)
}

func TestService_DeleteNTPServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/ntp_server/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteNTPServer(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
