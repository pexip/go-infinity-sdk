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

func TestService_ListRoles(t *testing.T) {
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
				expectedResponse := &RoleListResponse{
					Objects: []Role{
						{ID: 1, Name: "admin", Permissions: []Permission{
							{Name: "admin"},
							{Name: "user_management"},
						}},
						{ID: 2, Name: "user", Permissions: []Permission{
							{Name: "basic_access"},
						}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/role/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.RoleListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*RoleListResponse)
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
				expectedResponse := &RoleListResponse{
					Objects: []Role{
						{ID: 1, Name: "admin", Permissions: []Permission{
							{Name: "admin"},
							{Name: "user_management"},
						}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/role/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.RoleListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*RoleListResponse)
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
			result, err := service.ListRoles(t.Context(), tt.opts)

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

func TestService_GetRole(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedRole := &Role{
		ID:          1,
		Name:        "admin",
		Permissions: []Permission{
			{Name: "admin"},
			{Name: "user_management"},
			{Name: "conference_control"},
		},
		ResourceURI: "/api/admin/configuration/v1/role/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/role/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.Role")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Role)
		*result = *expectedRole
	})

	service := New(client)
	result, err := service.GetRole(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRole, result)
	client.AssertExpectations(t)
}

func TestService_CreateRole(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &RoleCreateRequest{
		Name:        "moderator",
		Permissions: []string{"conference_control", "user_management"},
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/role/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/role/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateRole(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateRole(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &RoleUpdateRequest{
		Name:        "updated-admin",
		Permissions: []string{"admin", "user_management", "system_control"},
	}

	expectedRole := &Role{
		ID:          1,
		Name:        "updated-admin",
		Permissions: []Permission{
			{Name: "admin"},
			{Name: "user_management"},
			{Name: "system_control"},
		},
		ResourceURI: "/api/admin/configuration/v1/role/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/role/1/", updateRequest, mock.AnythingOfType("*config.Role")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Role)
		*result = *expectedRole
	})

	service := New(client)
	result, err := service.UpdateRole(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedRole, result)
	client.AssertExpectations(t)
}

func TestService_DeleteRole(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/role/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteRole(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
