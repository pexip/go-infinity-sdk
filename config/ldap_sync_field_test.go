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

func TestService_ListLdapSyncFields(t *testing.T) {
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
				expectedResponse := &LdapSyncFieldListResponse{
					Objects: []LdapSyncField{
						{
							ID:                   1,
							Name:                 "email-field",
							Description:          "User email address field",
							TemplateVariableName: "user_email",
							IsBinary:             false,
						},
						{
							ID:                   2,
							Name:                 "photo-field",
							Description:          "User photo binary field",
							TemplateVariableName: "user_photo",
							IsBinary:             true,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ldap_sync_field/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.LdapSyncFieldListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*LdapSyncFieldListResponse)
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
				Search: "email",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &LdapSyncFieldListResponse{
					Objects: []LdapSyncField{
						{
							ID:                   1,
							Name:                 "email-field",
							Description:          "User email address field",
							TemplateVariableName: "user_email",
							IsBinary:             false,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ldap_sync_field/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.LdapSyncFieldListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*LdapSyncFieldListResponse)
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
			result, err := service.ListLdapSyncFields(t.Context(), tt.opts)

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

func TestService_GetLdapSyncField(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedField := &LdapSyncField{
		ID:                   1,
		Name:                 "test-sync-field",
		Description:          "Test LDAP sync field for user attributes",
		TemplateVariableName: "test_variable",
		IsBinary:             false,
		ResourceURI:          "/api/admin/configuration/v1/ldap_sync_field/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/ldap_sync_field/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.LdapSyncField")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*LdapSyncField)
		*result = *expectedField
	})

	service := New(client)
	result, err := service.GetLdapSyncField(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedField, result)
	client.AssertExpectations(t)
}

func TestService_CreateLdapSyncField(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &LdapSyncFieldCreateRequest{
		Name:                 "new-sync-field",
		Description:          "New LDAP sync field",
		TemplateVariableName: "new_variable",
		IsBinary:             true,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/ldap_sync_field/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/ldap_sync_field/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateLdapSyncField(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateLdapSyncField(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	isBinary := true
	updateRequest := &LdapSyncFieldUpdateRequest{
		Name:                 "updated-sync-field",
		Description:          "Updated LDAP sync field description",
		TemplateVariableName: "updated_variable",
		IsBinary:             &isBinary,
	}

	expectedField := &LdapSyncField{
		ID:                   1,
		Name:                 "updated-sync-field",
		Description:          "Updated LDAP sync field description",
		TemplateVariableName: "updated_variable",
		IsBinary:             true,
		ResourceURI:          "/api/admin/configuration/v1/ldap_sync_field/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/ldap_sync_field/1/", updateRequest, mock.AnythingOfType("*config.LdapSyncField")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*LdapSyncField)
		*result = *expectedField
	})

	service := New(client)
	result, err := service.UpdateLdapSyncField(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedField, result)
	client.AssertExpectations(t)
}

func TestService_DeleteLdapSyncField(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/ldap_sync_field/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteLdapSyncField(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
