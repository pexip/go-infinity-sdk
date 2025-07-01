package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListWebappAliases(t *testing.T) {
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
				bundle1 := "/api/admin/configuration/v1/software_bundle/1/"
				branding1 := "/api/admin/configuration/v1/webapp_branding/custom/"

				expectedResponse := &WebappAliasListResponse{
					Objects: []WebappAlias{
						{ID: 1, Slug: "meeting", Description: "Main meeting interface", WebappType: "meeting", IsEnabled: true, Bundle: &bundle1, Branding: &branding1},
						{ID: 2, Slug: "admin", Description: "Admin interface", WebappType: "admin", IsEnabled: true, Bundle: &bundle1},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/webapp_alias/", mock.AnythingOfType("*config.WebappAliasListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*WebappAliasListResponse)
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
				Search: "meeting",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				bundle := "/api/admin/configuration/v1/software_bundle/1/"
				branding := "/api/admin/configuration/v1/webapp_branding/custom/"

				expectedResponse := &WebappAliasListResponse{
					Objects: []WebappAlias{
						{ID: 1, Slug: "meeting", Description: "Main meeting interface", WebappType: "meeting", IsEnabled: true, Bundle: &bundle, Branding: &branding},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/webapp_alias/?limit=5&name__icontains=meeting", mock.AnythingOfType("*config.WebappAliasListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*WebappAliasListResponse)
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
			result, err := service.ListWebappAliases(t.Context(), tt.opts)

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

func TestService_GetWebappAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	bundle := "/api/admin/configuration/v1/software_bundle/1/"
	branding := "/api/admin/configuration/v1/webapp_branding/custom/"

	expectedWebappAlias := &WebappAlias{
		ID:          1,
		Slug:        "test-webapp",
		Description: "Test web app alias",
		WebappType:  "meeting",
		IsEnabled:   true,
		Bundle:      &bundle,
		Branding:    &branding,
		ResourceURI: "/api/admin/configuration/v1/webapp_alias/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/webapp_alias/1/", mock.AnythingOfType("*config.WebappAlias")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*WebappAlias)
		*result = *expectedWebappAlias
	})

	service := New(client)
	result, err := service.GetWebappAlias(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedWebappAlias, result)
	client.AssertExpectations(t)
}

func TestService_CreateWebappAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	bundle := "/api/admin/configuration/v1/software_bundle/1/"
	branding := "/api/admin/configuration/v1/webapp_branding/new/"

	createRequest := &WebappAliasCreateRequest{
		Slug:        "new-webapp",
		Description: "New web app alias",
		WebappType:  "meeting",
		IsEnabled:   true,
		Bundle:      &bundle,
		Branding:    &branding,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/webapp_alias/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/webapp_alias/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateWebappAlias(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateWebappAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	isEnabled := false
	branding := "/api/admin/configuration/v1/webapp_branding/updated/"

	updateRequest := &WebappAliasUpdateRequest{
		Description: "Updated web app alias",
		IsEnabled:   &isEnabled,
		Branding:    &branding,
	}

	bundle := "/api/admin/configuration/v1/software_bundle/1/"
	expectedWebappAlias := &WebappAlias{
		ID:          1,
		Slug:        "test-webapp",
		Description: "Updated web app alias",
		WebappType:  "meeting",
		IsEnabled:   false,
		Bundle:      &bundle,
		Branding:    &branding,
		ResourceURI: "/api/admin/configuration/v1/webapp_alias/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/webapp_alias/1/", updateRequest, mock.AnythingOfType("*config.WebappAlias")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*WebappAlias)
		*result = *expectedWebappAlias
	})

	service := New(client)
	result, err := service.UpdateWebappAlias(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedWebappAlias, result)
	client.AssertExpectations(t)
}

func TestService_DeleteWebappAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/webapp_alias/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteWebappAlias(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
