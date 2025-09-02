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

func TestService_ListIdentityProviders(t *testing.T) {
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
				expectedResponse := &IdentityProviderListResponse{
					Objects: []IdentityProvider{
						{
							ID:                          1,
							Name:                        "primary-saml-idp",
							Description:                 "Primary SAML Identity Provider",
							IdpType:                     "saml",
							UUID:                        "12345678-1234-5678-9abc-123456789012",
							SSOUrl:                      "https://idp.example.com/sso",
							IdpEntityID:                 "https://idp.example.com/entity",
							AssertionConsumerServiceURL: "https://example.com/acs",
							SignatureAlgorithm:          "RSA_SHA256",
							DigestAlgorithm:             "SHA256",
							WorkerFQDNACSURLs:           false,
							DisablePopupFlow:            false,
						},
						{
							ID:                          2,
							Name:                        "oidc-provider",
							Description:                 "OIDC Identity Provider",
							IdpType:                     "oidc",
							UUID:                        "87654321-4321-8765-cbaf-987654321098",
							OidcFlow:                    "authorization_code",
							OidcClientID:                "client123",
							OidcTokenURL:                "https://oidc.example.com/token",
							AssertionConsumerServiceURL: "https://example.com/oidc/callback",
							SignatureAlgorithm:          "RSA_SHA256",
							DigestAlgorithm:             "SHA256",
							WorkerFQDNACSURLs:           true,
							DisablePopupFlow:            true,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/identity_provider/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.IdentityProviderListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*IdentityProviderListResponse)
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
				Search: "saml",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &IdentityProviderListResponse{
					Objects: []IdentityProvider{
						{
							ID:                          1,
							Name:                        "primary-saml-idp",
							Description:                 "Primary SAML Identity Provider",
							IdpType:                     "saml",
							UUID:                        "12345678-1234-5678-9abc-123456789012",
							SSOUrl:                      "https://idp.example.com/sso",
							IdpEntityID:                 "https://idp.example.com/entity",
							AssertionConsumerServiceURL: "https://example.com/acs",
							SignatureAlgorithm:          "RSA_SHA256",
							DigestAlgorithm:             "SHA256",
							WorkerFQDNACSURLs:           false,
							DisablePopupFlow:            false,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/identity_provider/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.IdentityProviderListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*IdentityProviderListResponse)
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
			result, err := service.ListIdentityProviders(t.Context(), tt.opts)

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

func TestService_GetIdentityProvider(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedProvider := &IdentityProvider{
		ID:                                  1,
		Name:                                "test-saml-idp",
		Description:                         "Test SAML Identity Provider",
		IdpType:                             "saml",
		UUID:                                "12345678-1234-5678-9abc-123456789012",
		SSOUrl:                              "https://test-idp.example.com/sso",
		IdpEntityID:                         "https://test-idp.example.com/entity",
		IdpPublicKey:                        "-----BEGIN CERTIFICATE-----\ntest cert\n-----END CERTIFICATE-----",
		ServiceEntityID:                     "https://example.com/service",
		AssertionConsumerServiceURL:         "https://example.com/acs",
		AssertionConsumerServiceURL2:        "https://example.com/acs2",
		SignatureAlgorithm:                  "RSA_SHA256",
		DigestAlgorithm:                     "SHA256",
		DisplayNameAttributeName:            "displayName",
		RegistrationAliasAttributeName:      "mail",
		WorkerFQDNACSURLs:                   false,
		DisablePopupFlow:                    false,
		OidcTokenEndpointAuthScheme:         "client_secret_post",
		OidcTokenSignatureScheme:            "RS256",
		OidcFranceConnectRequiredEidasLevel: "eidas1",
		ResourceURI:                         "/api/admin/configuration/v1/identity_provider/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/identity_provider/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.IdentityProvider")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*IdentityProvider)
		*result = *expectedProvider
	})

	service := New(client)
	result, err := service.GetIdentityProvider(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedProvider, result)
	client.AssertExpectations(t)
}

func TestService_CreateIdentityProvider(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &IdentityProviderCreateRequest{
		Name:                                "new-saml-idp",
		Description:                         "New SAML Identity Provider",
		IdpType:                             "saml",
		UUID:                                "new12345-1234-5678-9abc-123456789012",
		SSOUrl:                              "https://new-idp.example.com/sso",
		IdpEntityID:                         "https://new-idp.example.com/entity",
		ServiceEntityID:                     "https://example.com/service",
		AssertionConsumerServiceURL:         "https://example.com/acs",
		SignatureAlgorithm:                  "RSA_SHA256",
		DigestAlgorithm:                     "SHA256",
		DisplayNameAttributeName:            "displayName",
		RegistrationAliasAttributeName:      "mail",
		WorkerFQDNACSURLs:                   false,
		DisablePopupFlow:                    false,
		OidcFlow:                            "authorization_code",
		OidcTokenEndpointAuthScheme:         "client_secret_post",
		OidcTokenSignatureScheme:            "RS256",
		OidcFranceConnectRequiredEidasLevel: "eidas1",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/identity_provider/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/identity_provider/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateIdentityProvider(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateIdentityProvider(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	disablePopupFlow := true
	updateRequest := &IdentityProviderUpdateRequest{
		Name:             "updated-saml-idp",
		Description:      "Updated SAML Identity Provider",
		SSOUrl:           "https://updated-idp.example.com/sso",
		DisablePopupFlow: &disablePopupFlow,
	}

	expectedProvider := &IdentityProvider{
		ID:                                  1,
		Name:                                "updated-saml-idp",
		Description:                         "Updated SAML Identity Provider",
		IdpType:                             "saml",
		UUID:                                "12345678-1234-5678-9abc-123456789012",
		SSOUrl:                              "https://updated-idp.example.com/sso",
		IdpEntityID:                         "https://test-idp.example.com/entity",
		ServiceEntityID:                     "https://example.com/service",
		AssertionConsumerServiceURL:         "https://example.com/acs",
		SignatureAlgorithm:                  "RSA_SHA256",
		DigestAlgorithm:                     "SHA256",
		DisplayNameAttributeName:            "displayName",
		RegistrationAliasAttributeName:      "mail",
		WorkerFQDNACSURLs:                   false,
		DisablePopupFlow:                    true,
		OidcTokenEndpointAuthScheme:         "client_secret_post",
		OidcTokenSignatureScheme:            "RS256",
		OidcFranceConnectRequiredEidasLevel: "eidas1",
		ResourceURI:                         "/api/admin/configuration/v1/identity_provider/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/identity_provider/1/", updateRequest, mock.AnythingOfType("*config.IdentityProvider")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*IdentityProvider)
		*result = *expectedProvider
	})

	service := New(client)
	result, err := service.UpdateIdentityProvider(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedProvider, result)
	client.AssertExpectations(t)
}

func TestService_DeleteIdentityProvider(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/identity_provider/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteIdentityProvider(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
