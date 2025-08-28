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

func TestService_ListIdentityProviderGroups(t *testing.T) {
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
				expectedResponse := &IdentityProviderGroupListResponse{
					Objects: []IdentityProviderGroup{
						{
							ID:          1,
							Name:        "admin-group",
							Description: "Administrator identity provider group",
							IdentityProvider: []string{
								"/api/admin/configuration/v1/identity_provider/1/",
								"/api/admin/configuration/v1/identity_provider/2/",
							},
						},
						{
							ID:          2,
							Name:        "user-group",
							Description: "Standard user identity provider group",
							IdentityProvider: []string{
								"/api/admin/configuration/v1/identity_provider/3/",
							},
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/identity_provider_group/", mock.AnythingOfType("*config.IdentityProviderGroupListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*IdentityProviderGroupListResponse)
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
				Search: "admin",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &IdentityProviderGroupListResponse{
					Objects: []IdentityProviderGroup{
						{
							ID:          1,
							Name:        "admin-group",
							Description: "Administrator identity provider group",
							IdentityProvider: []string{
								"/api/admin/configuration/v1/identity_provider/1/",
								"/api/admin/configuration/v1/identity_provider/2/",
							},
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/identity_provider_group/?limit=5&name__icontains=admin", mock.AnythingOfType("*config.IdentityProviderGroupListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*IdentityProviderGroupListResponse)
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
			result, err := service.ListIdentityProviderGroups(t.Context(), tt.opts)

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

func TestService_GetIdentityProviderGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedGroup := &IdentityProviderGroup{
		ID:          1,
		Name:        "test-idp-group",
		Description: "Test identity provider group",
		IdentityProvider: []string{
			"/api/admin/configuration/v1/identity_provider/1/",
			"/api/admin/configuration/v1/identity_provider/2/",
		},
		ResourceURI: "/api/admin/configuration/v1/identity_provider_group/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/identity_provider_group/1/", mock.AnythingOfType("*config.IdentityProviderGroup")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*IdentityProviderGroup)
		*result = *expectedGroup
	})

	service := New(client)
	result, err := service.GetIdentityProviderGroup(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedGroup, result)
	client.AssertExpectations(t)
}

func TestService_CreateIdentityProviderGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &IdentityProviderGroupCreateRequest{
		Name:        "new-idp-group",
		Description: "New identity provider group",
		IdentityProvider: []string{
			"/api/admin/configuration/v1/identity_provider/1/",
		},
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/identity_provider_group/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/identity_provider_group/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateIdentityProviderGroup(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateIdentityProviderGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &IdentityProviderGroupUpdateRequest{
		Name:        "updated-idp-group",
		Description: "Updated identity provider group",
		IdentityProvider: []string{
			"/api/admin/configuration/v1/identity_provider/1/",
			"/api/admin/configuration/v1/identity_provider/3/",
		},
	}

	expectedGroup := &IdentityProviderGroup{
		ID:          1,
		Name:        "updated-idp-group",
		Description: "Updated identity provider group",
		IdentityProvider: []string{
			"/api/admin/configuration/v1/identity_provider/1/",
			"/api/admin/configuration/v1/identity_provider/3/",
		},
		ResourceURI: "/api/admin/configuration/v1/identity_provider_group/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/identity_provider_group/1/", updateRequest, mock.AnythingOfType("*config.IdentityProviderGroup")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*IdentityProviderGroup)
		*result = *expectedGroup
	})

	service := New(client)
	result, err := service.UpdateIdentityProviderGroup(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedGroup, result)
	client.AssertExpectations(t)
}

func TestService_DeleteIdentityProviderGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/identity_provider_group/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteIdentityProviderGroup(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
