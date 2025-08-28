/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListIVRThemes(t *testing.T) {
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
				lastUpdated := util.InfinityTime{Time: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)}
				expectedResponse := &IVRThemeListResponse{
					Objects: []IVRTheme{
						{
							ID:      1,
							Name:    "default-theme",
							Package: "default_package.tar.gz",
							UUID:    "12345678-1234-5678-9abc-123456789012",
							Conference: []string{
								"/api/admin/configuration/v1/conference/1/",
								"/api/admin/configuration/v1/conference/2/",
							},
							CustomLayouts:  `{"layout1": "config1"}`,
							PinningConfigs: `{"pin1": "config1"}`,
							LastUpdated:    lastUpdated,
						},
						{
							ID:             2,
							Name:           "custom-theme",
							Package:        "custom_package.tar.gz",
							UUID:           "87654321-4321-8765-cbaf-987654321098",
							Conference:     []string{},
							CustomLayouts:  `{"layout2": "config2"}`,
							PinningConfigs: `{"pin2": "config2"}`,
							LastUpdated:    lastUpdated,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ivr_theme/", mock.AnythingOfType("*config.IVRThemeListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*IVRThemeListResponse)
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
				Search: "default",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				lastUpdated := util.InfinityTime{Time: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)}
				expectedResponse := &IVRThemeListResponse{
					Objects: []IVRTheme{
						{
							ID:      1,
							Name:    "default-theme",
							Package: "default_package.tar.gz",
							UUID:    "12345678-1234-5678-9abc-123456789012",
							Conference: []string{
								"/api/admin/configuration/v1/conference/1/",
								"/api/admin/configuration/v1/conference/2/",
							},
							CustomLayouts:  `{"layout1": "config1"}`,
							PinningConfigs: `{"pin1": "config1"}`,
							LastUpdated:    lastUpdated,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ivr_theme/?limit=5&name__icontains=default", mock.AnythingOfType("*config.IVRThemeListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*IVRThemeListResponse)
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
			result, err := service.ListIVRThemes(t.Context(), tt.opts)

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

func TestService_GetIVRTheme(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	lastUpdated := util.InfinityTime{Time: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)}
	expectedTheme := &IVRTheme{
		ID:      1,
		Name:    "test-theme",
		Package: "test_package.tar.gz",
		UUID:    "12345678-1234-5678-9abc-123456789012",
		Conference: []string{
			"/api/admin/configuration/v1/conference/1/",
		},
		CustomLayouts:  `{"layout": "test-config"}`,
		PinningConfigs: `{"pin": "test-pin-config"}`,
		LastUpdated:    lastUpdated,
		ResourceURI:    "/api/admin/configuration/v1/ivr_theme/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/ivr_theme/1/", mock.AnythingOfType("*config.IVRTheme")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*IVRTheme)
		*result = *expectedTheme
	})

	service := New(client)
	result, err := service.GetIVRTheme(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedTheme, result)
	client.AssertExpectations(t)
}

func TestService_CreateIVRTheme(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &IVRThemeCreateRequest{
		Name:    "new-theme",
		Package: "new_package.tar.gz",
		UUID:    "new12345-1234-5678-9abc-123456789012",
		Conference: []string{
			"/api/admin/configuration/v1/conference/3/",
		},
		CustomLayouts:  `{"newLayout": "newConfig"}`,
		PinningConfigs: `{"newPin": "newPinConfig"}`,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/ivr_theme/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/ivr_theme/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateIVRTheme(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateIVRTheme(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &IVRThemeUpdateRequest{
		Name:    "updated-theme",
		Package: "updated_package.tar.gz",
		Conference: []string{
			"/api/admin/configuration/v1/conference/1/",
			"/api/admin/configuration/v1/conference/4/",
		},
		CustomLayouts: `{"updatedLayout": "updatedConfig"}`,
	}

	lastUpdated := util.InfinityTime{Time: time.Date(2024, 1, 2, 12, 0, 0, 0, time.UTC)}
	expectedTheme := &IVRTheme{
		ID:      1,
		Name:    "updated-theme",
		Package: "updated_package.tar.gz",
		UUID:    "12345678-1234-5678-9abc-123456789012",
		Conference: []string{
			"/api/admin/configuration/v1/conference/1/",
			"/api/admin/configuration/v1/conference/4/",
		},
		CustomLayouts:  `{"updatedLayout": "updatedConfig"}`,
		PinningConfigs: `{"pin": "test-pin-config"}`,
		LastUpdated:    lastUpdated,
		ResourceURI:    "/api/admin/configuration/v1/ivr_theme/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/ivr_theme/1/", updateRequest, mock.AnythingOfType("*config.IVRTheme")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*IVRTheme)
		*result = *expectedTheme
	})

	service := New(client)
	result, err := service.UpdateIVRTheme(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedTheme, result)
	client.AssertExpectations(t)
}

func TestService_DeleteIVRTheme(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/ivr_theme/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteIVRTheme(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
