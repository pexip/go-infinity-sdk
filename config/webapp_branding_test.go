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
						{Name: "default", Description: "Default branding", UUID: "123e4567-e89b-12d3-a456-426614174000", WebappType: "meeting", IsDefault: true, BrandingFile: "default.zip", LastUpdated: lastUpdated1},
						{Name: "custom", Description: "Custom branding", UUID: "123e4567-e89b-12d3-a456-426614174001", WebappType: "meeting", IsDefault: false, BrandingFile: "custom.zip", LastUpdated: lastUpdated2},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/webapp_branding/", mock.AnythingOfType("*config.WebappBrandingListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*WebappBrandingListResponse)
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
						{Name: "custom", Description: "Custom branding", UUID: "123e4567-e89b-12d3-a456-426614174001", WebappType: "meeting", IsDefault: false, BrandingFile: "custom.zip", LastUpdated: lastUpdated},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/webapp_branding/?limit=5&name__icontains=custom", mock.AnythingOfType("*config.WebappBrandingListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*WebappBrandingListResponse)
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
		Name:         name,
		Description:  "Test webapp branding",
		UUID:         "123e4567-e89b-12d3-a456-426614174000",
		WebappType:   "meeting",
		IsDefault:    false,
		BrandingFile: "test-branding.zip",
		LastUpdated:  lastUpdated,
		ResourceURI:  "/api/admin/configuration/v1/webapp_branding/" + name + "/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/webapp_branding/"+name+"/", mock.AnythingOfType("*config.WebappBranding")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*WebappBranding)
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
		Name:         "new-branding",
		Description:  "New webapp branding",
		UUID:         "123e4567-e89b-12d3-a456-426614174002",
		WebappType:   "meeting",
		IsDefault:    false,
		BrandingFile: "new-branding.zip",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/webapp_branding/new-branding/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/webapp_branding/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateWebappBranding(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateWebappBranding(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	isDefault := true
	name := "test-branding"

	updateRequest := &WebappBrandingUpdateRequest{
		Description:  "Updated webapp branding",
		IsDefault:    &isDefault,
		BrandingFile: "updated-branding.zip",
	}

	lastUpdated := util.InfinityTime{}
	expectedWebappBranding := &WebappBranding{
		Name:         name,
		Description:  "Updated webapp branding",
		UUID:         "123e4567-e89b-12d3-a456-426614174000",
		WebappType:   "meeting",
		IsDefault:    true,
		BrandingFile: "updated-branding.zip",
		LastUpdated:  lastUpdated,
		ResourceURI:  "/api/admin/configuration/v1/webapp_branding/" + name + "/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/webapp_branding/"+name+"/", updateRequest, mock.AnythingOfType("*config.WebappBranding")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*WebappBranding)
		*result = *expectedWebappBranding
	})

	service := New(client)
	result, err := service.UpdateWebappBranding(t.Context(), name, updateRequest)

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
