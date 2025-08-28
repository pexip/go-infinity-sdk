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

func TestService_ListAzureTenants(t *testing.T) {
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
				expectedResponse := &AzureTenantListResponse{
					Objects: []AzureTenant{
						{ID: 1, Name: "primary-tenant", Description: "Primary Teams tenant", TenantID: "12345678-1234-1234-1234-123456789abc"},
						{ID: 2, Name: "secondary-tenant", Description: "Secondary Teams tenant", TenantID: "87654321-4321-4321-4321-cba987654321"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/azure_tenant/", mock.AnythingOfType("*config.AzureTenantListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*AzureTenantListResponse)
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
				expectedResponse := &AzureTenantListResponse{
					Objects: []AzureTenant{
						{ID: 1, Name: "primary-tenant", Description: "Primary Teams tenant", TenantID: "12345678-1234-1234-1234-123456789abc"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/azure_tenant/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.AzureTenantListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*AzureTenantListResponse)
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
			result, err := service.ListAzureTenants(t.Context(), tt.opts)

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

func TestService_GetAzureTenant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedAzureTenant := &AzureTenant{
		ID:          1,
		Name:        "test-tenant",
		Description: "Test Teams tenant",
		TenantID:    "12345678-1234-1234-1234-123456789abc",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/azure_tenant/1/", mock.AnythingOfType("*config.AzureTenant")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*AzureTenant)
		*result = *expectedAzureTenant
	})

	service := New(client)
	result, err := service.GetAzureTenant(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedAzureTenant, result)
	client.AssertExpectations(t)
}

func TestService_CreateAzureTenant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &AzureTenantCreateRequest{
		Name:        "new-tenant",
		Description: "New Teams tenant",
		TenantID:    "87654321-4321-4321-4321-cba987654321",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/azure_tenant/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/azure_tenant/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateAzureTenant(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateAzureTenant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &AzureTenantUpdateRequest{
		Description: "Updated Teams tenant",
		TenantID:    "11111111-2222-3333-4444-555555555555",
	}

	expectedAzureTenant := &AzureTenant{
		ID:          1,
		Name:        "test-tenant",
		Description: "Updated Teams tenant",
		TenantID:    "11111111-2222-3333-4444-555555555555",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/azure_tenant/1/", updateRequest, mock.AnythingOfType("*config.AzureTenant")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*AzureTenant)
		*result = *expectedAzureTenant
	})

	service := New(client)
	result, err := service.UpdateAzureTenant(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedAzureTenant, result)
	client.AssertExpectations(t)
}

func TestService_DeleteAzureTenant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/azure_tenant/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteAzureTenant(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
