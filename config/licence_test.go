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

func TestService_ListLicences(t *testing.T) {
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
				expectedResponse := &LicenceListResponse{
					Objects: []Licence{
						{FulfillmentID: "12345", ProductID: "pexip-infinity", Status: "active", Concurrent: 100, Activatable: 25},
						{FulfillmentID: "67890", ProductID: "pexip-vmr", Status: "active", Concurrent: 50, Activatable: 10},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/licence/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.LicenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*LicenceListResponse)
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
				Search: "Infinity",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &LicenceListResponse{
					Objects: []Licence{
						{FulfillmentID: "12345", ProductID: "pexip-infinity", Status: "active", Concurrent: 100, Activatable: 25},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/licence/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.LicenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*LicenceListResponse)
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
			result, err := service.ListLicences(t.Context(), tt.opts)

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

func TestService_GetLicence(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedLicence := &Licence{
		FulfillmentID:  "test-12345",
		ProductID:      "pexip-infinity",
		Status:         "active",
		Concurrent:     100,
		Activatable:    25,
		ExpirationDate: "2024-12-31",
		ResourceURI:    "/api/admin/configuration/v1/licence/test-12345/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/licence/test-12345/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.Licence")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Licence)
		*result = *expectedLicence
	})

	service := New(client)
	result, err := service.GetLicence(t.Context(), "test-12345")

	assert.NoError(t, err)
	assert.Equal(t, expectedLicence, result)
	client.AssertExpectations(t)
}

func TestService_CreateLicence(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &LicenceCreateRequest{
		EntitlementID: "ABC123-DEF456-GHI789",
		OfflineMode:   false,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/licence/new-licence-id/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/licence/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateLicence(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_DeleteLicence(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/licence/test-12345/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteLicence(t.Context(), "test-12345")

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
