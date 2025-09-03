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

func TestService_ListGMSAccessTokens(t *testing.T) {
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
				expectedResponse := &GMSAccessTokenListResponse{
					Objects: []GMSAccessToken{
						{ID: 1, Name: "primary-gms-token", Token: "token1"},
						{ID: 2, Name: "secondary-gms-token", Token: "token2"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/gms_access_token/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.GMSAccessTokenListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*GMSAccessTokenListResponse)
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
				expectedResponse := &GMSAccessTokenListResponse{
					Objects: []GMSAccessToken{
						{ID: 1, Name: "primary-gms-token", Token: "token1"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/gms_access_token/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.GMSAccessTokenListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*GMSAccessTokenListResponse)
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
			result, err := service.ListGMSAccessTokens(t.Context(), tt.opts)

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

func TestService_GetGMSAccessToken(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedGMSAccessToken := &GMSAccessToken{
		ID:    1,
		Name:  "test-gms-token",
		Token: "test-token-value",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/gms_access_token/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.GMSAccessToken")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*GMSAccessToken)
		*result = *expectedGMSAccessToken
	})

	service := New(client)
	result, err := service.GetGMSAccessToken(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedGMSAccessToken, result)
	client.AssertExpectations(t)
}

func TestService_CreateGMSAccessToken(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &GMSAccessTokenCreateRequest{
		Name:  "new-gms-token",
		Token: "new-token-value",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/gms_access_token/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/gms_access_token/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateGMSAccessToken(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateGMSAccessToken(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &GMSAccessTokenUpdateRequest{
		Name:  "updated-gms-token",
		Token: "updated-token-value",
	}

	expectedGMSAccessToken := &GMSAccessToken{
		ID:    1,
		Name:  "updated-gms-token",
		Token: "updated-token-value",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/gms_access_token/1/", updateRequest, mock.AnythingOfType("*config.GMSAccessToken")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*GMSAccessToken)
		*result = *expectedGMSAccessToken
	})

	service := New(client)
	result, err := service.UpdateGMSAccessToken(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedGMSAccessToken, result)
	client.AssertExpectations(t)
}

func TestService_DeleteGMSAccessToken(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/gms_access_token/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteGMSAccessToken(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
