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

func TestService_ListSystemLocations(t *testing.T) {
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
				expectedResponse := &SystemLocationListResponse{
					Objects: []SystemLocation{
						{ID: 1, Name: "Location 1", Description: "First location"},
						{ID: 2, Name: "Location 2", Description: "Second location"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/system_location/", mock.AnythingOfType("*config.SystemLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SystemLocationListResponse)
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
					Offset: 6,
				},
				Search: "location",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &SystemLocationListResponse{
					Objects: []SystemLocation{
						{ID: 1, Name: "Test Location", Description: "Test system location"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/system_location/?limit=3&name__icontains=location&offset=6", mock.AnythingOfType("*config.SystemLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SystemLocationListResponse)
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
			result, err := service.ListSystemLocations(t.Context(), tt.opts)

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

func TestService_GetSystemLocation(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedLocation := &SystemLocation{
		ID:          1,
		Name:        "Primary Location",
		Description: "Main system location",
		MTU:         1500,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/system_location/1/", mock.AnythingOfType("*config.SystemLocation")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemLocation)
		*result = *expectedLocation
	})

	service := New(client)
	result, err := service.GetSystemLocation(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, result)
	client.AssertExpectations(t)
}

func TestService_CreateSystemLocation(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &SystemLocationCreateRequest{
		Name:        "New Location",
		Description: "New system location",
		MTU:         1400,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/system_location/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/system_location/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateSystemLocation(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateSystemLocation(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &SystemLocationUpdateRequest{
		Description: "Updated location",
		MTU:         1300,
	}

	expectedLocation := &SystemLocation{
		ID:          1,
		Name:        "Test Location",
		Description: "Updated location",
		MTU:         1300,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/system_location/1/", updateRequest, mock.AnythingOfType("*config.SystemLocation")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SystemLocation)
		*result = *expectedLocation
	})

	service := New(client)
	result, err := service.UpdateSystemLocation(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, result)
	client.AssertExpectations(t)
}

func TestService_DeleteSystemLocation(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/system_location/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteSystemLocation(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
