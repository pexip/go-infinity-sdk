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

func TestService_ListUserGroups(t *testing.T) {
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
				expectedResponse := &UserGroupListResponse{
					Objects: []UserGroup{
						{ID: 1, Name: "administrators", Description: "System administrators", Users: []string{"/api/admin/configuration/v1/end_user/1/", "/api/admin/configuration/v1/end_user/2/"}, UserGroupEntityMappings: []string{"/api/admin/configuration/v1/user_group_entity_mapping/1/"}},
						{ID: 2, Name: "users", Description: "Regular users", Users: []string{"/api/admin/configuration/v1/end_user/3/"}, UserGroupEntityMappings: []string{}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/user_group/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.UserGroupListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*UserGroupListResponse)
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
				expectedResponse := &UserGroupListResponse{
					Objects: []UserGroup{
						{ID: 1, Name: "administrators", Description: "System administrators", Users: []string{"/api/admin/configuration/v1/end_user/1/", "/api/admin/configuration/v1/end_user/2/"}, UserGroupEntityMappings: []string{"/api/admin/configuration/v1/user_group_entity_mapping/1/"}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/user_group/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.UserGroupListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*UserGroupListResponse)
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
			result, err := service.ListUserGroups(t.Context(), tt.opts)

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

func TestService_GetUserGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedUserGroup := &UserGroup{
		ID:                      1,
		Name:                    "test-group",
		Description:             "Test user group",
		Users:                   []string{"/api/admin/configuration/v1/end_user/1/", "/api/admin/configuration/v1/end_user/2/"},
		UserGroupEntityMappings: []string{"/api/admin/configuration/v1/user_group_entity_mapping/1/"},
		ResourceURI:             "/api/admin/configuration/v1/user_group/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/user_group/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.UserGroup")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*UserGroup)
		*result = *expectedUserGroup
	})

	service := New(client)
	result, err := service.GetUserGroup(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUserGroup, result)
	client.AssertExpectations(t)
}

func TestService_CreateUserGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &UserGroupCreateRequest{
		Name:                    "new-group",
		Description:             "New user group",
		Users:                   []string{"/api/admin/configuration/v1/end_user/3/"},
		UserGroupEntityMappings: []string{"/api/admin/configuration/v1/user_group_entity_mapping/2/"},
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/user_group/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/user_group/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateUserGroup(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateUserGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &UserGroupUpdateRequest{
		Description: "Updated user group",
		Users:       []string{"/api/admin/configuration/v1/end_user/1/"},
	}

	expectedUserGroup := &UserGroup{
		ID:                      1,
		Name:                    "test-group",
		Description:             "Updated user group",
		Users:                   []string{"/api/admin/configuration/v1/end_user/1/"},
		UserGroupEntityMappings: []string{"/api/admin/configuration/v1/user_group_entity_mapping/1/"},
		ResourceURI:             "/api/admin/configuration/v1/user_group/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/user_group/1/", updateRequest, mock.AnythingOfType("*config.UserGroup")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*UserGroup)
		*result = *expectedUserGroup
	})

	service := New(client)
	result, err := service.UpdateUserGroup(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedUserGroup, result)
	client.AssertExpectations(t)
}

func TestService_DeleteUserGroup(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/user_group/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteUserGroup(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
