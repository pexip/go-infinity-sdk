/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"strings"
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
							ID:   1,
							Name: "default-theme",
							UUID: "12345678-1234-5678-9abc-123456789012",
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
							UUID:           "87654321-4321-8765-cbaf-987654321098",
							Conference:     []string{},
							CustomLayouts:  `{"layout2": "config2"}`,
							PinningConfigs: `{"pin2": "config2"}`,
							LastUpdated:    lastUpdated,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ivr_theme/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.IVRThemeListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*IVRThemeListResponse)
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
							ID:   1,
							Name: "default-theme",
							UUID: "12345678-1234-5678-9abc-123456789012",
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
				m.On("GetJSON", t.Context(), "configuration/v1/ivr_theme/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.IVRThemeListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*IVRThemeListResponse)
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
		ID:   1,
		Name: "test-theme",
		UUID: "12345678-1234-5678-9abc-123456789012",
		Conference: []string{
			"/api/admin/configuration/v1/conference/1/",
		},
		CustomLayouts:  `{"layout": "test-config"}`,
		PinningConfigs: `{"pin": "test-pin-config"}`,
		LastUpdated:    lastUpdated,
		ResourceURI:    "/api/admin/configuration/v1/ivr_theme/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/ivr_theme/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.IVRTheme")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*IVRTheme)
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

	packageContent := strings.NewReader("mock package content")
	createRequest := &IVRThemeCreateRequest{
		Name: "new-theme",
		UUID: "new12345-1234-5678-9abc-123456789012",
		Conference: []string{
			"/api/admin/configuration/v1/conference/3/",
		},
		CustomLayouts:  `{"newLayout": "newConfig"}`,
		PinningConfigs: `{"newPin": "newPinConfig"}`,
	}

	expectedFields := map[string]string{
		"name":            "new-theme",
		"uuid":            "new12345-1234-5678-9abc-123456789012",
		"conference":      `["/api/admin/configuration/v1/conference/3/"]`,
		"custom_layouts":  `{"newLayout": "newConfig"}`,
		"pinning_configs": `{"newPin": "newPinConfig"}`,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/ivr_theme/123/",
	}

	// Expect single call: POST multipart form with all fields and file
	client.On("PostMultipartFormWithFieldsAndResponse", t.Context(), "configuration/v1/ivr_theme/", expectedFields, "package", "new_package.tar.gz", packageContent, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateIVRTheme(t.Context(), createRequest, "new_package.tar.gz", packageContent)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateIVRTheme(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	packageContent := strings.NewReader("mock updated package content")
	updateRequest := &IVRThemeUpdateRequest{
		Name: "updated-theme",
		Conference: []string{
			"/api/admin/configuration/v1/conference/1/",
			"/api/admin/configuration/v1/conference/4/",
		},
		CustomLayouts: `{"updatedLayout": "updatedConfig"}`,
	}

	expectedFields := map[string]string{
		"name":            "updated-theme",
		"uuid":            "",
		"conference":      `["/api/admin/configuration/v1/conference/1/","/api/admin/configuration/v1/conference/4/"]`,
		"custom_layouts":  `{"updatedLayout": "updatedConfig"}`,
		"pinning_configs": "",
	}

	lastUpdated := util.InfinityTime{Time: time.Date(2024, 1, 2, 12, 0, 0, 0, time.UTC)}
	expectedTheme := &IVRTheme{
		ID:   1,
		Name: "updated-theme",
		UUID: "12345678-1234-5678-9abc-123456789012",
		Conference: []string{
			"/api/admin/configuration/v1/conference/1/",
			"/api/admin/configuration/v1/conference/4/",
		},
		CustomLayouts:  `{"updatedLayout": "updatedConfig"}`,
		PinningConfigs: `{"pin": "test-pin-config"}`,
		LastUpdated:    lastUpdated,
		ResourceURI:    "/api/admin/configuration/v1/ivr_theme/1/",
	}

	// Expect single call: PATCH multipart form with all fields and file
	client.On("PatchMultipartFormWithFieldsAndResponse", t.Context(), "configuration/v1/ivr_theme/1/", expectedFields, "package", "updated_package.tar.gz", packageContent, mock.AnythingOfType("*config.IVRTheme")).Return(nil, nil).Run(func(args mock.Arguments) {
		result := args.Get(6).(*IVRTheme)
		*result = *expectedTheme
	})

	service := New(client)
	result, err := service.UpdateIVRTheme(t.Context(), 1, updateRequest, "updated_package.tar.gz", packageContent)

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
