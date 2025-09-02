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

func TestService_ListADFSAuthServerDomains(t *testing.T) {
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
				expectedResponse := &ADFSAuthServerDomainListResponse{
					Objects: []ADFSAuthServerDomain{
						{ID: 1, Domain: "example.com", Description: "Primary domain", ADFSAuthServer: "/api/admin/configuration/v1/adfs_auth_server/1/"},
						{ID: 2, Domain: "subdomain.example.com", Description: "Subdomain", ADFSAuthServer: "/api/admin/configuration/v1/adfs_auth_server/1/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/adfs_auth_server_domain/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ADFSAuthServerDomainListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ADFSAuthServerDomainListResponse)
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
				Search: "example",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &ADFSAuthServerDomainListResponse{
					Objects: []ADFSAuthServerDomain{
						{ID: 1, Domain: "example.com", Description: "Primary domain", ADFSAuthServer: "/api/admin/configuration/v1/adfs_auth_server/1/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/adfs_auth_server_domain/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ADFSAuthServerDomainListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ADFSAuthServerDomainListResponse)
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
			result, err := service.ListADFSAuthServerDomains(t.Context(), tt.opts)

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

func TestService_GetADFSAuthServerDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedADFSAuthServerDomain := &ADFSAuthServerDomain{
		ID:             1,
		Domain:         "example.com",
		Description:    "Test domain",
		ADFSAuthServer: "/api/admin/configuration/v1/adfs_auth_server/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/adfs_auth_server_domain/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ADFSAuthServerDomain")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ADFSAuthServerDomain)
		*result = *expectedADFSAuthServerDomain
	})

	service := New(client)
	result, err := service.GetADFSAuthServerDomain(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedADFSAuthServerDomain, result)
	client.AssertExpectations(t)
}

func TestService_CreateADFSAuthServerDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &ADFSAuthServerDomainCreateRequest{
		Domain:         "new.example.com",
		Description:    "New domain",
		ADFSAuthServer: "/api/admin/configuration/v1/adfs_auth_server/1/",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/adfs_auth_server_domain/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/adfs_auth_server_domain/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateADFSAuthServerDomain(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateADFSAuthServerDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &ADFSAuthServerDomainUpdateRequest{
		Description: "Updated domain",
		Domain:      "updated.example.com",
	}

	expectedADFSAuthServerDomain := &ADFSAuthServerDomain{
		ID:             1,
		Domain:         "updated.example.com",
		Description:    "Updated domain",
		ADFSAuthServer: "/api/admin/configuration/v1/adfs_auth_server/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/adfs_auth_server_domain/1/", updateRequest, mock.AnythingOfType("*config.ADFSAuthServerDomain")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ADFSAuthServerDomain)
		*result = *expectedADFSAuthServerDomain
	})

	service := New(client)
	result, err := service.UpdateADFSAuthServerDomain(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedADFSAuthServerDomain, result)
	client.AssertExpectations(t)
}

func TestService_DeleteADFSAuthServerDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/adfs_auth_server_domain/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteADFSAuthServerDomain(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
