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

func TestService_ListExternalWebappHosts(t *testing.T) {
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
				expectedResponse := &ExternalWebappHostListResponse{
					Objects: []ExternalWebappHost{
						{ID: 1, Address: "webapp1.example.com"},
						{ID: 2, Address: "webapp2.example.com"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/external_webapp_host/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ExternalWebappHostListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ExternalWebappHostListResponse)
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
				Search: "webapp1",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &ExternalWebappHostListResponse{
					Objects: []ExternalWebappHost{
						{ID: 1, Address: "webapp1.example.com"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/external_webapp_host/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ExternalWebappHostListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ExternalWebappHostListResponse)
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
			result, err := service.ListExternalWebappHosts(t.Context(), tt.opts)

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

func TestService_GetExternalWebappHost(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedExternalWebappHost := &ExternalWebappHost{
		ID:      1,
		Address: "test-webapp.example.com",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/external_webapp_host/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ExternalWebappHost")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ExternalWebappHost)
		*result = *expectedExternalWebappHost
	})

	service := New(client)
	result, err := service.GetExternalWebappHost(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedExternalWebappHost, result)
	client.AssertExpectations(t)
}

func TestService_CreateExternalWebappHost(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &ExternalWebappHostCreateRequest{
		Address: "new-webapp.example.com",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/external_webapp_host/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/external_webapp_host/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateExternalWebappHost(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateExternalWebappHost(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &ExternalWebappHostUpdateRequest{
		Address: "updated-webapp.example.com",
	}

	expectedExternalWebappHost := &ExternalWebappHost{
		ID:      1,
		Address: "updated-webapp.example.com",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/external_webapp_host/1/", updateRequest, mock.AnythingOfType("*config.ExternalWebappHost")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ExternalWebappHost)
		*result = *expectedExternalWebappHost
	})

	service := New(client)
	result, err := service.UpdateExternalWebappHost(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedExternalWebappHost, result)
	client.AssertExpectations(t)
}

func TestService_DeleteExternalWebappHost(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/external_webapp_host/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteExternalWebappHost(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
