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

func TestService_ListMjxEndpointGroups(t *testing.T) {
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
				integration1 := "/api/admin/configuration/v1/mjx_integration/1/"
				location1 := "/api/admin/configuration/v1/system_location/1/"
				integration2 := "/api/admin/configuration/v1/mjx_integration/2/"
				location2 := "/api/admin/configuration/v1/system_location/2/"

				expectedResponse := &MjxEndpointGroupListResponse{
					Objects: []MjxEndpointGroup{
						{ID: 1, Name: "floor-1-rooms", Description: "Conference rooms on floor 1", MjxIntegration: &integration1, SystemLocation: &location1, DisableProxy: false, Endpoints: []string{"/api/admin/configuration/v1/mjx_endpoint/1/", "/api/admin/configuration/v1/mjx_endpoint/2/"}},
						{ID: 2, Name: "floor-2-rooms", Description: "Conference rooms on floor 2", MjxIntegration: &integration2, SystemLocation: &location2, DisableProxy: true, Endpoints: []string{"/api/admin/configuration/v1/mjx_endpoint/3/", "/api/admin/configuration/v1/mjx_endpoint/4/"}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_endpoint_group/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxEndpointGroupListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MjxEndpointGroupListResponse)
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
				Search: "floor-1",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				integration := "/api/admin/configuration/v1/mjx_integration/1/"
				location := "/api/admin/configuration/v1/system_location/1/"

				expectedResponse := &MjxEndpointGroupListResponse{
					Objects: []MjxEndpointGroup{
						{ID: 1, Name: "floor-1-rooms", Description: "Conference rooms on floor 1", MjxIntegration: &integration, SystemLocation: &location, DisableProxy: false, Endpoints: []string{"/api/admin/configuration/v1/mjx_endpoint/1/", "/api/admin/configuration/v1/mjx_endpoint/2/"}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_endpoint_group/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxEndpointGroupListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MjxEndpointGroupListResponse)
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
			result, err := service.ListMjxEndpointGroups(t.Context(), tt.opts)

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

func TestService_GetMjxEndpointGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	integration := "/api/admin/configuration/v1/mjx_integration/1/"
	location := "/api/admin/configuration/v1/system_location/1/"

	expectedMjxEndpointGroup := &MjxEndpointGroup{
		ID:             1,
		Name:           "test-group",
		Description:    "Test MJX endpoint group",
		MjxIntegration: &integration,
		SystemLocation: &location,
		DisableProxy:   false,
		Endpoints: []string{
			"/api/admin/configuration/v1/mjx_endpoint/1/",
			"/api/admin/configuration/v1/mjx_endpoint/2/",
			"/api/admin/configuration/v1/mjx_endpoint/3/",
		},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/mjx_endpoint_group/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxEndpointGroup")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxEndpointGroup)
		*result = *expectedMjxEndpointGroup
	})

	service := New(client)
	result, err := service.GetMjxEndpointGroup(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedMjxEndpointGroup, result)
	client.AssertExpectations(t)
}

func TestService_CreateMjxEndpointGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	integration := "/api/admin/configuration/v1/mjx_integration/2/"
	location := "/api/admin/configuration/v1/system_location/3/"

	createRequest := &MjxEndpointGroupCreateRequest{
		Name:           "new-group",
		Description:    "New MJX endpoint group",
		MjxIntegration: &integration,
		SystemLocation: &location,
		DisableProxy:   true,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/mjx_endpoint_group/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/mjx_endpoint_group/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMjxEndpointGroup(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMjxEndpointGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	disableProxy := true
	newLocation := "/api/admin/configuration/v1/system_location/5/"

	updateRequest := &MjxEndpointGroupUpdateRequest{
		Description:    "Updated MJX endpoint group",
		SystemLocation: &newLocation,
		DisableProxy:   &disableProxy,
	}

	integration := "/api/admin/configuration/v1/mjx_integration/1/"

	expectedMjxEndpointGroup := &MjxEndpointGroup{
		ID:             1,
		Name:           "test-group",
		Description:    "Updated MJX endpoint group",
		MjxIntegration: &integration,
		SystemLocation: &newLocation,
		DisableProxy:   true,
		Endpoints: []string{
			"/api/admin/configuration/v1/mjx_endpoint/1/",
			"/api/admin/configuration/v1/mjx_endpoint/2/",
			"/api/admin/configuration/v1/mjx_endpoint/3/",
		},
	}

	client.On("PutJSON", t.Context(), "configuration/v1/mjx_endpoint_group/1/", updateRequest, mock.AnythingOfType("*config.MjxEndpointGroup")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxEndpointGroup)
		*result = *expectedMjxEndpointGroup
	})

	service := New(client)
	result, err := service.UpdateMjxEndpointGroup(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedMjxEndpointGroup, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMjxEndpointGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/mjx_endpoint_group/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMjxEndpointGroup(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
