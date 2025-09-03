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

func TestService_ListMjxExchangeDeployments(t *testing.T) {
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
				oauthClientID := "test-client-id"
				oauthState := "test-state"
				expectedResponse := &MjxExchangeDeploymentListResponse{
					Objects: []MjxExchangeDeployment{
						{
							ID:                       1,
							Name:                     "primary-exchange",
							Description:              "Primary Exchange deployment",
							ServiceAccountUsername:   "svc-exchange@example.com",
							AuthenticationMethod:     "oauth",
							DisableProxy:             false,
							FindItemsRequestQuota:    1000,
							KerberosAuthEveryRequest: false,
							KerberosEnableTLS:        true,
							OAuthClientID:            &oauthClientID,
							OAuthState:               &oauthState,
							OAuthAuthEndpoint:        "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
							OAuthTokenEndpoint:       "https://login.microsoftonline.com/common/oauth2/v2.0/token",
							OAuthRedirectURI:         "https://pexip.example.com/oauth/callback",
						},
						{
							ID:                       2,
							Name:                     "backup-exchange",
							Description:              "Backup Exchange deployment",
							ServiceAccountUsername:   "svc-backup@example.com",
							AuthenticationMethod:     "kerberos",
							DisableProxy:             true,
							FindItemsRequestQuota:    500,
							KerberosRealm:            "EXAMPLE.COM",
							KerberosKDC:              "kdc.example.com",
							KerberosExchangeSPN:      "HTTP/exchange.example.com",
							KerberosAuthEveryRequest: true,
							KerberosEnableTLS:        true,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_exchange_deployment/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxExchangeDeploymentListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MjxExchangeDeploymentListResponse)
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
				oauthClientID := "test-client-id"
				expectedResponse := &MjxExchangeDeploymentListResponse{
					Objects: []MjxExchangeDeployment{
						{
							ID:                     1,
							Name:                   "primary-exchange",
							Description:            "Primary Exchange deployment",
							ServiceAccountUsername: "svc-exchange@example.com",
							AuthenticationMethod:   "oauth",
							OAuthClientID:          &oauthClientID,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_exchange_deployment/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxExchangeDeploymentListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MjxExchangeDeploymentListResponse)
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
			result, err := service.ListMjxExchangeDeployments(t.Context(), tt.opts)

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

func TestService_GetMjxExchangeDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	oauthClientID := "test-client-id"
	oauthState := "test-state"

	expectedDeployment := &MjxExchangeDeployment{
		ID:                             1,
		Name:                           "test-exchange",
		Description:                    "Test Exchange deployment",
		ServiceAccountUsername:         "svc-test@example.com",
		ServiceAccountPassword:         "password123",
		AuthenticationMethod:           "oauth",
		EWSURL:                         "https://exchange.example.com/ews/exchange.asmx",
		DisableProxy:                   false,
		FindItemsRequestQuota:          1000,
		KerberosRealm:                  "EXAMPLE.COM",
		KerberosKDC:                    "kdc.example.com",
		KerberosExchangeSPN:            "HTTP/exchange.example.com",
		KerberosAuthEveryRequest:       false,
		KerberosEnableTLS:              true,
		KerberosKDCHTTPSProxy:          "proxy.example.com:8080",
		KerberosVerifyTLSUsingCustomCA: false,
		OAuthClientID:                  &oauthClientID,
		OAuthAuthEndpoint:              "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
		OAuthTokenEndpoint:             "https://login.microsoftonline.com/common/oauth2/v2.0/token",
		OAuthRedirectURI:               "https://pexip.example.com/oauth/callback",
		OAuthRefreshToken:              "refresh-token-123",
		OAuthState:                     &oauthState,
		AutodiscoverURLs:               []string{"/api/admin/configuration/v1/mjx_exchange_autodiscover_url/1/"},
		MjxIntegrations:                []string{"/api/admin/configuration/v1/mjx_integration/1/"},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/mjx_exchange_deployment/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxExchangeDeployment")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxExchangeDeployment)
		*result = *expectedDeployment
	})

	service := New(client)
	result, err := service.GetMjxExchangeDeployment(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedDeployment, result)
	client.AssertExpectations(t)
}

func TestService_CreateMjxExchangeDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	oauthClientID := "new-client-id"
	createRequest := &MjxExchangeDeploymentCreateRequest{
		Name:                           "new-exchange",
		Description:                    "New Exchange deployment",
		ServiceAccountUsername:         "svc-new@example.com",
		ServiceAccountPassword:         "newpassword123",
		AuthenticationMethod:           "oauth",
		EWSURL:                         "https://new-exchange.example.com/ews/exchange.asmx",
		DisableProxy:                   false,
		FindItemsRequestQuota:          2000,
		KerberosRealm:                  "EXAMPLE.COM",
		KerberosKDC:                    "kdc.example.com",
		KerberosExchangeSPN:            "HTTP/new-exchange.example.com",
		KerberosAuthEveryRequest:       false,
		KerberosEnableTLS:              true,
		KerberosKDCHTTPSProxy:          "proxy.example.com:8080",
		KerberosVerifyTLSUsingCustomCA: false,
		OAuthClientID:                  &oauthClientID,
		OAuthAuthEndpoint:              "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
		OAuthTokenEndpoint:             "https://login.microsoftonline.com/common/oauth2/v2.0/token",
		OAuthRedirectURI:               "https://pexip.example.com/oauth/callback",
		OAuthRefreshToken:              "new-refresh-token",
		AutodiscoverURLs:               []string{"/api/admin/configuration/v1/mjx_exchange_autodiscover_url/1/"},
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/mjx_exchange_deployment/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/mjx_exchange_deployment/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMjxExchangeDeployment(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMjxExchangeDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	disableProxy := true
	findItemsQuota := 1500
	kerberosEnableTLS := false

	updateRequest := &MjxExchangeDeploymentUpdateRequest{
		Description:           "Updated Exchange deployment",
		DisableProxy:          &disableProxy,
		FindItemsRequestQuota: &findItemsQuota,
		KerberosEnableTLS:     &kerberosEnableTLS,
	}

	oauthClientID := "test-client-id"
	expectedDeployment := &MjxExchangeDeployment{
		ID:                     1,
		Name:                   "test-exchange",
		Description:            "Updated Exchange deployment",
		ServiceAccountUsername: "svc-test@example.com",
		AuthenticationMethod:   "oauth",
		DisableProxy:           true,
		FindItemsRequestQuota:  1500,
		KerberosEnableTLS:      false,
		OAuthClientID:          &oauthClientID,
		OAuthAuthEndpoint:      "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
		OAuthTokenEndpoint:     "https://login.microsoftonline.com/common/oauth2/v2.0/token",
		OAuthRedirectURI:       "https://pexip.example.com/oauth/callback",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/mjx_exchange_deployment/1/", updateRequest, mock.AnythingOfType("*config.MjxExchangeDeployment")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxExchangeDeployment)
		*result = *expectedDeployment
	})

	service := New(client)
	result, err := service.UpdateMjxExchangeDeployment(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedDeployment, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMjxExchangeDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/mjx_exchange_deployment/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMjxExchangeDeployment(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
