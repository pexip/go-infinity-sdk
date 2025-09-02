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

func TestService_ListDevices(t *testing.T) {
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
				expectedResponse := &DeviceListResponse{
					Objects: []Device{
						{ID: 1, Alias: "device1", Description: "Test device 1", EnableSIP: true},
						{ID: 2, Alias: "device2", Description: "Test device 2", EnableH323: true},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/device/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.DeviceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*DeviceListResponse)
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
				Search: "device1",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &DeviceListResponse{
					Objects: []Device{
						{ID: 1, Alias: "device1", Description: "Test device 1", EnableSIP: true},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/device/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.DeviceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*DeviceListResponse)
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
			result, err := service.ListDevices(t.Context(), tt.opts)

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

func TestService_GetDevice(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedDevice := &Device{
		ID:                          1,
		Alias:                       "test-device",
		Description:                 "Test Device",
		Username:                    "testuser",
		Password:                    "testpass",
		PrimaryOwnerEmailAddress:    "test@example.com",
		EnableSIP:                   true,
		EnableH323:                  false,
		EnableInfinityConnectNonSSO: true,
		EnableInfinityConnectSSO:    false,
		EnableStandardSSO:           false,
		Tag:                         "test-tag",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/device/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.Device")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Device)
		*result = *expectedDevice
	})

	service := New(client)
	result, err := service.GetDevice(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedDevice, result)
	client.AssertExpectations(t)
}

func TestService_CreateDevice(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &DeviceCreateRequest{
		Alias:                       "new-device",
		Description:                 "New Test Device",
		Username:                    "newuser",
		Password:                    "newpass",
		PrimaryOwnerEmailAddress:    "new@example.com",
		EnableSIP:                   true,
		EnableH323:                  false,
		EnableInfinityConnectNonSSO: true,
		EnableInfinityConnectSSO:    false,
		EnableStandardSSO:           false,
		Tag:                         "new-tag",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/device/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/device/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateDevice(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateDevice(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	enableSIP := true
	updateRequest := &DeviceUpdateRequest{
		Description:              "Updated Device",
		PrimaryOwnerEmailAddress: "updated@example.com",
		EnableSIP:                &enableSIP,
	}

	expectedDevice := &Device{
		ID:                          1,
		Alias:                       "test-device",
		Description:                 "Updated Device",
		Username:                    "testuser",
		Password:                    "testpass",
		PrimaryOwnerEmailAddress:    "updated@example.com",
		EnableSIP:                   true,
		EnableH323:                  false,
		EnableInfinityConnectNonSSO: true,
		EnableInfinityConnectSSO:    false,
		EnableStandardSSO:           false,
		Tag:                         "test-tag",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/device/1/", updateRequest, mock.AnythingOfType("*config.Device")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Device)
		*result = *expectedDevice
	})

	service := New(client)
	result, err := service.UpdateDevice(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedDevice, result)
	client.AssertExpectations(t)
}

func TestService_DeleteDevice(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/device/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteDevice(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
