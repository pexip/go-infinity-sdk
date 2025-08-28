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

func TestService_ListMSSIPProxies(t *testing.T) {
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
				port1 := 5060
				port2 := 5061
				expectedResponse := &MSSIPProxyListResponse{
					Objects: []MSSIPProxy{
						{
							ID:          1,
							Name:        "primary-mssip-proxy",
							Description: "Primary MS-SIP proxy",
							Address:     "mssip.example.com",
							Port:        &port1,
							Transport:   "tcp",
						},
						{
							ID:          2,
							Name:        "backup-mssip-proxy",
							Description: "Backup MS-SIP proxy",
							Address:     "backup-mssip.example.com",
							Port:        &port2,
							Transport:   "tls",
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mssip_proxy/", mock.AnythingOfType("*config.MSSIPProxyListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MSSIPProxyListResponse)
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
				port := 5060
				expectedResponse := &MSSIPProxyListResponse{
					Objects: []MSSIPProxy{
						{
							ID:          1,
							Name:        "primary-mssip-proxy",
							Description: "Primary MS-SIP proxy",
							Address:     "mssip.example.com",
							Port:        &port,
							Transport:   "tcp",
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mssip_proxy/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.MSSIPProxyListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MSSIPProxyListResponse)
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
			result, err := service.ListMSSIPProxies(t.Context(), tt.opts)

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

func TestService_GetMSSIPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	port := 5060

	expectedProxy := &MSSIPProxy{
		ID:          1,
		Name:        "test-mssip-proxy",
		Description: "Test MS-SIP proxy",
		Address:     "test-mssip.example.com",
		Port:        &port,
		Transport:   "tcp",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/mssip_proxy/1/", mock.AnythingOfType("*config.MSSIPProxy")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MSSIPProxy)
		*result = *expectedProxy
	})

	service := New(client)
	result, err := service.GetMSSIPProxy(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedProxy, result)
	client.AssertExpectations(t)
}

func TestService_CreateMSSIPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	port := 5060
	createRequest := &MSSIPProxyCreateRequest{
		Name:        "new-mssip-proxy",
		Description: "New MS-SIP proxy",
		Address:     "new-mssip.example.com",
		Port:        &port,
		Transport:   "tls",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/mssip_proxy/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/mssip_proxy/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMSSIPProxy(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMSSIPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	port := 5061
	updateRequest := &MSSIPProxyUpdateRequest{
		Description: "Updated MS-SIP proxy",
		Port:        &port,
		Transport:   "tls",
	}

	expectedProxy := &MSSIPProxy{
		ID:          1,
		Name:        "test-mssip-proxy",
		Description: "Updated MS-SIP proxy",
		Address:     "test-mssip.example.com",
		Port:        &port,
		Transport:   "tls",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/mssip_proxy/1/", updateRequest, mock.AnythingOfType("*config.MSSIPProxy")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MSSIPProxy)
		*result = *expectedProxy
	})

	service := New(client)
	result, err := service.UpdateMSSIPProxy(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedProxy, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMSSIPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/mssip_proxy/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMSSIPProxy(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
