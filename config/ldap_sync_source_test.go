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

func TestService_ListLdapSyncSources(t *testing.T) {
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
				expectedResponse := &LdapSyncSourceListResponse{
					Objects: []LdapSyncSource{
						{ID: 1, Name: "primary-ldap", Description: "Primary LDAP sync source", LdapServer: "ldap.example.com", LdapBaseDN: "dc=example,dc=com", LdapBindUsername: "cn=admin,dc=example,dc=com", LdapUseGlobalCatalog: false, LdapPermitNoTLS: false},
						{ID: 2, Name: "secondary-ldap", Description: "Secondary LDAP sync source", LdapServer: "ldap2.example.com", LdapBaseDN: "dc=example,dc=com", LdapBindUsername: "cn=admin,dc=example,dc=com", LdapUseGlobalCatalog: true, LdapPermitNoTLS: true},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ldap_sync_source/", mock.AnythingOfType("*config.LdapSyncSourceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*LdapSyncSourceListResponse)
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
				Search: "primary",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &LdapSyncSourceListResponse{
					Objects: []LdapSyncSource{
						{ID: 1, Name: "primary-ldap", Description: "Primary LDAP sync source", LdapServer: "ldap.example.com", LdapBaseDN: "dc=example,dc=com", LdapBindUsername: "cn=admin,dc=example,dc=com", LdapUseGlobalCatalog: false, LdapPermitNoTLS: false},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ldap_sync_source/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.LdapSyncSourceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*LdapSyncSourceListResponse)
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
			result, err := service.ListLdapSyncSources(t.Context(), tt.opts)

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

func TestService_GetLdapSyncSource(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedLdapSyncSource := &LdapSyncSource{
		ID:                   1,
		Name:                 "test-ldap",
		Description:          "Test LDAP sync source",
		LdapServer:           "ldap.example.com",
		LdapBaseDN:           "dc=example,dc=com",
		LdapBindUsername:     "cn=admin,dc=example,dc=com",
		LdapBindPassword:     "secret123",
		LdapUseGlobalCatalog: false,
		LdapPermitNoTLS:      false,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/ldap_sync_source/1/", mock.AnythingOfType("*config.LdapSyncSource")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*LdapSyncSource)
		*result = *expectedLdapSyncSource
	})

	service := New(client)
	result, err := service.GetLdapSyncSource(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedLdapSyncSource, result)
	client.AssertExpectations(t)
}

func TestService_CreateLdapSyncSource(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &LdapSyncSourceCreateRequest{
		Name:                 "new-ldap",
		Description:          "New LDAP sync source",
		LdapServer:           "new-ldap.example.com",
		LdapBaseDN:           "dc=example,dc=com",
		LdapBindUsername:     "cn=admin,dc=example,dc=com",
		LdapBindPassword:     "newsecret123",
		LdapUseGlobalCatalog: true,
		LdapPermitNoTLS:      false,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/ldap_sync_source/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/ldap_sync_source/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateLdapSyncSource(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateLdapSyncSource(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	useGlobalCatalog := true
	updateRequest := &LdapSyncSourceUpdateRequest{
		Description:          "Updated LDAP sync source",
		LdapUseGlobalCatalog: &useGlobalCatalog,
	}

	expectedLdapSyncSource := &LdapSyncSource{
		ID:                   1,
		Name:                 "test-ldap",
		Description:          "Updated LDAP sync source",
		LdapServer:           "ldap.example.com",
		LdapBaseDN:           "dc=example,dc=com",
		LdapBindUsername:     "cn=admin,dc=example,dc=com",
		LdapBindPassword:     "secret123",
		LdapUseGlobalCatalog: true,
		LdapPermitNoTLS:      false,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/ldap_sync_source/1/", updateRequest, mock.AnythingOfType("*config.LdapSyncSource")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*LdapSyncSource)
		*result = *expectedLdapSyncSource
	})

	service := New(client)
	result, err := service.UpdateLdapSyncSource(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedLdapSyncSource, result)
	client.AssertExpectations(t)
}

func TestService_DeleteLdapSyncSource(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/ldap_sync_source/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteLdapSyncSource(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
