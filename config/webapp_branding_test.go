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
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListWebappBrandings(t *testing.T) {
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
				lastUpdated1 := util.InfinityTime{}
				lastUpdated2 := util.InfinityTime{}

				expectedResponse := &WebappBrandingListResponse{
					Objects: []WebappBranding{
						{Name: "default", Description: "Default branding", UUID: "123e4567-e89b-12d3-a456-426614174000", WebappType: "meeting", IsDefault: true, LastUpdated: lastUpdated1},
						{Name: "custom", Description: "Custom branding", UUID: "123e4567-e89b-12d3-a456-426614174001", WebappType: "meeting", IsDefault: false, LastUpdated: lastUpdated2},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/webapp_branding/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.WebappBrandingListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*WebappBrandingListResponse)
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
				Search: "custom",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				lastUpdated := util.InfinityTime{}

				expectedResponse := &WebappBrandingListResponse{
					Objects: []WebappBranding{
						{Name: "custom", Description: "Custom branding", UUID: "123e4567-e89b-12d3-a456-426614174001", WebappType: "meeting", IsDefault: false, LastUpdated: lastUpdated},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/webapp_branding/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.WebappBrandingListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*WebappBrandingListResponse)
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
			result, err := service.ListWebappBrandings(t.Context(), tt.opts)

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

func TestService_GetWebappBranding(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	lastUpdated := util.InfinityTime{}
	name := "test-branding"

	expectedWebappBranding := &WebappBranding{
		Name:        name,
		Description: "Test webapp branding",
		UUID:        "123e4567-e89b-12d3-a456-426614174000",
		WebappType:  "meeting",
		IsDefault:   false,
		LastUpdated: lastUpdated,
		ResourceURI: "/api/admin/configuration/v1/webapp_branding/" + name + "/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/webapp_branding/"+name+"/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.WebappBranding")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*WebappBranding)
		*result = *expectedWebappBranding
	})

	service := New(client)
	result, err := service.GetWebappBranding(t.Context(), name)

	assert.NoError(t, err)
	assert.Equal(t, expectedWebappBranding, result)
	client.AssertExpectations(t)
}

func TestService_CreateWebappBranding(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &WebappBrandingCreateRequest{
		Name:        "new-branding",
		Description: "New webapp branding",
		WebappType:  "meeting",
	}

	expectedBranding := &WebappBranding{
		UUID:        "123e4567-e89b-12d3-a456-426614174002",
		Name:        "new-branding",
		Description: "New webapp branding",
		WebappType:  "meeting",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/webapp_branding/123e4567-e89b-12d3-a456-426614174002/",
	}

	expectedFields := map[string]string{
		"name":        "new-branding",
		"description": "New webapp branding",
		"webapp_type": "meeting",
	}

	// Step 1: Mock POST without file to create resource
	client.On("PostMultipartFormWithFieldsAndResponse", t.Context(), "configuration/v1/webapp_branding/", expectedFields, "", "", mock.Anything, mock.Anything).Return(expectedResponse, nil)

	// Step 2: Mock PATCH to upload file
	client.On("PatchMultipartFormWithFieldsAndResponse", t.Context(), "configuration/v1/webapp_branding/123e4567-e89b-12d3-a456-426614174002/", mock.Anything, "branding_file", "test.zip", mock.Anything, mock.AnythingOfType("*config.WebappBranding")).Return(nil, nil).Run(func(args mock.Arguments) {
		result := args.Get(6).(*WebappBranding)
		*result = *expectedBranding
	})

	service := New(client)
	result, err := service.CreateWebappBranding(t.Context(), createRequest, "test.zip", nil)

	assert.NoError(t, err)
	assert.Equal(t, expectedBranding, result)
	client.AssertExpectations(t)
}

// TestService_UpdateWebappBranding is commented out because UpdateWebappBranding is not yet implemented
func TestService_UpdateWebappBranding(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	uuid := "123e4567-e89b-12d3-a456-426614174002"

	updateRequest := &WebappBrandingUpdateRequest{
		Name:        "test-branding",
		Description: "Updated webapp branding",
		WebappType:  "meeting",
	}

	lastUpdated := util.InfinityTime{}
	expectedWebappBranding := &WebappBranding{
		Name:        "test-branding",
		Description: "Updated webapp branding",
		UUID:        uuid,
		WebappType:  "meeting",
		IsDefault:   false,
		LastUpdated: lastUpdated,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/webapp_branding/"+uuid+"/", mock.AnythingOfType("*config.WebappBrandingUpdateRequest"), mock.AnythingOfType("*config.WebappBranding")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*WebappBranding)
		*result = *expectedWebappBranding
	})

	service := New(client)
	result, err := service.UpdateWebappBranding(t.Context(), updateRequest, uuid)

	assert.NoError(t, err)
	assert.Equal(t, expectedWebappBranding, result)
	client.AssertExpectations(t)
}

func TestService_DeleteWebappBranding(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	name := "test-branding"

	client.On("DeleteJSON", t.Context(), "configuration/v1/webapp_branding/"+name+"/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteWebappBranding(t.Context(), name)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
