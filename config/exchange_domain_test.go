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

func TestService_ListExchangeDomains(t *testing.T) {
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
				expectedResponse := &ExchangeDomainListResponse{
					Objects: []ExchangeDomain{
						{ID: 1, Domain: "example.com", ExchangeConnector: "/api/admin/configuration/v1/exchange_connector/1/"},
						{ID: 2, Domain: "subdomain.example.com", ExchangeConnector: "/api/admin/configuration/v1/exchange_connector/2/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/exchange_domain/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ExchangeDomainListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ExchangeDomainListResponse)
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
				expectedResponse := &ExchangeDomainListResponse{
					Objects: []ExchangeDomain{
						{ID: 1, Domain: "example.com", ExchangeConnector: "/api/admin/configuration/v1/exchange_connector/1/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/exchange_domain/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ExchangeDomainListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ExchangeDomainListResponse)
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
			result, err := service.ListExchangeDomains(t.Context(), tt.opts)

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

func TestService_GetExchangeDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedExchangeDomain := &ExchangeDomain{
		ID:                1,
		Domain:            "test.example.com",
		ExchangeConnector: "/api/admin/configuration/v1/exchange_connector/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/exchange_domain/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ExchangeDomain")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ExchangeDomain)
		*result = *expectedExchangeDomain
	})

	service := New(client)
	result, err := service.GetExchangeDomain(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedExchangeDomain, result)
	client.AssertExpectations(t)
}

func TestService_CreateExchangeDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &ExchangeDomainCreateRequest{
		Domain:            "new.example.com",
		ExchangeConnector: "/api/admin/configuration/v1/exchange_connector/2/",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/exchange_domain/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/exchange_domain/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateExchangeDomain(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateExchangeDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &ExchangeDomainUpdateRequest{
		Domain:            "updated.example.com",
		ExchangeConnector: "/api/admin/configuration/v1/exchange_connector/3/",
	}

	expectedExchangeDomain := &ExchangeDomain{
		ID:                1,
		Domain:            "updated.example.com",
		ExchangeConnector: "/api/admin/configuration/v1/exchange_connector/3/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/exchange_domain/1/", updateRequest, mock.AnythingOfType("*config.ExchangeDomain")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ExchangeDomain)
		*result = *expectedExchangeDomain
	})

	service := New(client)
	result, err := service.UpdateExchangeDomain(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedExchangeDomain, result)
	client.AssertExpectations(t)
}

func TestService_DeleteExchangeDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/exchange_domain/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteExchangeDomain(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
