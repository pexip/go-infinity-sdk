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

func TestService_ListH323Gatekeepers(t *testing.T) {
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
				port := 1719
				expectedResponse := &H323GatekeeperListResponse{
					Objects: []H323Gatekeeper{
						{ID: 1, Name: "primary-gk", Description: "Primary H.323 gatekeeper", Address: "gk.example.com", Port: &port},
						{ID: 2, Name: "backup-gk", Description: "Backup H.323 gatekeeper", Address: "backup-gk.example.com", Port: &port},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/h323_gatekeeper/", mock.AnythingOfType("*config.H323GatekeeperListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*H323GatekeeperListResponse)
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
				port := 1719
				expectedResponse := &H323GatekeeperListResponse{
					Objects: []H323Gatekeeper{
						{ID: 1, Name: "primary-gk", Description: "Primary H.323 gatekeeper", Address: "gk.example.com", Port: &port},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/h323_gatekeeper/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.H323GatekeeperListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*H323GatekeeperListResponse)
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
			result, err := service.ListH323Gatekeepers(t.Context(), tt.opts)

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

func TestService_GetH323Gatekeeper(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	port := 1719
	expectedGatekeeper := &H323Gatekeeper{
		ID:          1,
		Name:        "test-gk",
		Description: "Test H.323 gatekeeper",
		Address:     "gk.example.com",
		Port:        &port,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/h323_gatekeeper/1/", mock.AnythingOfType("*config.H323Gatekeeper")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*H323Gatekeeper)
		*result = *expectedGatekeeper
	})

	service := New(client)
	result, err := service.GetH323Gatekeeper(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedGatekeeper, result)
	client.AssertExpectations(t)
}

func TestService_CreateH323Gatekeeper(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	port := 1719

	createRequest := &H323GatekeeperCreateRequest{
		Name:        "new-gk",
		Description: "New H.323 gatekeeper",
		Address:     "new-gk.example.com",
		Port:        &port,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/h323_gatekeeper/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/h323_gatekeeper/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateH323Gatekeeper(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateH323Gatekeeper(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	port := 1720

	updateRequest := &H323GatekeeperUpdateRequest{
		Description: "Updated H.323 gatekeeper",
		Port:        &port,
	}

	expectedGatekeeper := &H323Gatekeeper{
		ID:          1,
		Name:        "test-gk",
		Description: "Updated H.323 gatekeeper",
		Address:     "gk.example.com",
		Port:        &port,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/h323_gatekeeper/1/", updateRequest, mock.AnythingOfType("*config.H323Gatekeeper")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*H323Gatekeeper)
		*result = *expectedGatekeeper
	})

	service := New(client)
	result, err := service.UpdateH323Gatekeeper(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedGatekeeper, result)
	client.AssertExpectations(t)
}

func TestService_DeleteH323Gatekeeper(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/h323_gatekeeper/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteH323Gatekeeper(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
