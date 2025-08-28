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

func TestService_ListRoleMappings(t *testing.T) {
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
				expectedResponse := &RoleMappingListResponse{
					Objects: []RoleMapping{
						{ID: 1, Name: "admin-mapping", Source: "ldap", Value: "CN=Administrators,OU=Groups,DC=example,DC=com", Roles: []string{"admin", "user"}},
						{ID: 2, Name: "user-mapping", Source: "ldap", Value: "CN=Users,OU=Groups,DC=example,DC=com", Roles: []string{"user"}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/role_mapping/", mock.AnythingOfType("*config.RoleMappingListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*RoleMappingListResponse)
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
				expectedResponse := &RoleMappingListResponse{
					Objects: []RoleMapping{
						{ID: 1, Name: "admin-mapping", Source: "ldap", Value: "CN=Administrators,OU=Groups,DC=example,DC=com", Roles: []string{"admin", "user"}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/role_mapping/?limit=5&name__icontains=admin", mock.AnythingOfType("*config.RoleMappingListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*RoleMappingListResponse)
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
			result, err := service.ListRoleMappings(t.Context(), tt.opts)

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

func TestService_GetRoleMapping(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedRoleMapping := &RoleMapping{
		ID:     1,
		Name:   "test-role-mapping",
		Source: "ldap",
		Value:  "CN=TestGroup,OU=Groups,DC=example,DC=com",
		Roles:  []string{"admin", "user", "moderator"},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/role_mapping/1/", mock.AnythingOfType("*config.RoleMapping")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*RoleMapping)
		*result = *expectedRoleMapping
	})

	service := New(client)
	result, err := service.GetRoleMapping(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoleMapping, result)
	client.AssertExpectations(t)
}

func TestService_CreateRoleMapping(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &RoleMappingCreateRequest{
		Name:   "new-role-mapping",
		Source: "oidc",
		Value:  "developer",
		Roles:  []string{"user", "developer"},
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/role_mapping/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/role_mapping/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateRoleMapping(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateRoleMapping(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &RoleMappingUpdateRequest{
		Name:  "updated-role-mapping",
		Value: "CN=UpdatedGroup,OU=Groups,DC=example,DC=com",
		Roles: []string{"admin", "moderator"},
	}

	expectedRoleMapping := &RoleMapping{
		ID:     1,
		Name:   "updated-role-mapping",
		Source: "ldap",
		Value:  "CN=UpdatedGroup,OU=Groups,DC=example,DC=com",
		Roles:  []string{"admin", "moderator"},
	}

	client.On("PutJSON", t.Context(), "configuration/v1/role_mapping/1/", updateRequest, mock.AnythingOfType("*config.RoleMapping")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*RoleMapping)
		*result = *expectedRoleMapping
	})

	service := New(client)
	result, err := service.UpdateRoleMapping(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoleMapping, result)
	client.AssertExpectations(t)
}

func TestService_DeleteRoleMapping(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/role_mapping/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteRoleMapping(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
