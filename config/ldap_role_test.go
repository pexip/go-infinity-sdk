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

func TestService_ListLdapRoles(t *testing.T) {
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
				expectedResponse := &LdapRoleListResponse{
					Objects: []LdapRole{
						{
							ID:          1,
							Name:        "admin-ldap-role",
							LdapGroupDN: "CN=Administrators,CN=Users,DC=example,DC=com",
							Roles: []string{
								"/api/admin/configuration/v1/role/1/",
								"/api/admin/configuration/v1/role/2/",
							},
						},
						{
							ID:          2,
							Name:        "user-ldap-role",
							LdapGroupDN: "CN=Users,CN=Users,DC=example,DC=com",
							Roles: []string{
								"/api/admin/configuration/v1/role/3/",
							},
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ldap_role/", mock.AnythingOfType("*config.LdapRoleListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*LdapRoleListResponse)
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
				expectedResponse := &LdapRoleListResponse{
					Objects: []LdapRole{
						{
							ID:          1,
							Name:        "admin-ldap-role",
							LdapGroupDN: "CN=Administrators,CN=Users,DC=example,DC=com",
							Roles: []string{
								"/api/admin/configuration/v1/role/1/",
								"/api/admin/configuration/v1/role/2/",
							},
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ldap_role/?limit=5&name__icontains=admin", mock.AnythingOfType("*config.LdapRoleListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*LdapRoleListResponse)
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
			result, err := service.ListLdapRoles(t.Context(), tt.opts)

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

func TestService_GetLdapRole(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedRole := &LdapRole{
		ID:          1,
		Name:        "test-ldap-role",
		LdapGroupDN: "CN=TestGroup,CN=Users,DC=example,DC=com",
		Roles: []string{
			"/api/admin/configuration/v1/role/1/",
			"/api/admin/configuration/v1/role/2/",
		},
		ResourceURI: "/api/admin/configuration/v1/ldap_role/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/ldap_role/1/", mock.AnythingOfType("*config.LdapRole")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*LdapRole)
		*result = *expectedRole
	})

	service := New(client)
	result, err := service.GetLdapRole(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRole, result)
	client.AssertExpectations(t)
}

func TestService_CreateLdapRole(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &LdapRoleCreateRequest{
		Name:        "new-ldap-role",
		LdapGroupDN: "CN=NewGroup,CN=Users,DC=example,DC=com",
		Roles: []string{
			"/api/admin/configuration/v1/role/3/",
		},
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/ldap_role/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/ldap_role/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateLdapRole(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateLdapRole(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &LdapRoleUpdateRequest{
		Name:        "updated-ldap-role",
		LdapGroupDN: "CN=UpdatedGroup,CN=Users,DC=example,DC=com",
		Roles: []string{
			"/api/admin/configuration/v1/role/1/",
			"/api/admin/configuration/v1/role/4/",
		},
	}

	expectedRole := &LdapRole{
		ID:          1,
		Name:        "updated-ldap-role",
		LdapGroupDN: "CN=UpdatedGroup,CN=Users,DC=example,DC=com",
		Roles: []string{
			"/api/admin/configuration/v1/role/1/",
			"/api/admin/configuration/v1/role/4/",
		},
		ResourceURI: "/api/admin/configuration/v1/ldap_role/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/ldap_role/1/", updateRequest, mock.AnythingOfType("*config.LdapRole")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*LdapRole)
		*result = *expectedRole
	})

	service := New(client)
	result, err := service.UpdateLdapRole(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedRole, result)
	client.AssertExpectations(t)
}

func TestService_DeleteLdapRole(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/ldap_role/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteLdapRole(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
