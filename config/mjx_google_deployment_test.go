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

func TestService_ListMjxGoogleDeployments(t *testing.T) {
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
				oauthState := "test-state"
				expectedResponse := &MjxGoogleDeploymentListResponse{
					Objects: []MjxGoogleDeployment{
						{
							ID:                         1,
							Name:                       "primary-google",
							Description:                "Primary Google deployment",
							ClientEmail:                "service-account@project.iam.gserviceaccount.com",
							ClientID:                   "123456789012345678901",
							UseUserConsent:             false,
							AuthEndpoint:               "https://accounts.google.com/o/oauth2/auth",
							TokenEndpoint:              "https://oauth2.googleapis.com/token",
							RedirectURI:                "https://pexip.example.com/oauth/google/callback",
							MaximumNumberOfAPIRequests: 1000,
							OAuthState:                 &oauthState,
							MjxIntegrations:            []string{"/api/admin/configuration/v1/mjx_integration/1/"},
						},
						{
							ID:                         2,
							Name:                       "backup-google",
							Description:                "Backup Google deployment",
							ClientEmail:                "backup-service@project.iam.gserviceaccount.com",
							ClientID:                   "987654321098765432109",
							UseUserConsent:             true,
							AuthEndpoint:               "https://accounts.google.com/o/oauth2/auth",
							TokenEndpoint:              "https://oauth2.googleapis.com/token",
							RedirectURI:                "https://backup.example.com/oauth/google/callback",
							MaximumNumberOfAPIRequests: 500,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_google_deployment/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxGoogleDeploymentListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MjxGoogleDeploymentListResponse)
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
				expectedResponse := &MjxGoogleDeploymentListResponse{
					Objects: []MjxGoogleDeployment{
						{
							ID:                         1,
							Name:                       "primary-google",
							Description:                "Primary Google deployment",
							ClientEmail:                "service-account@project.iam.gserviceaccount.com",
							UseUserConsent:             false,
							MaximumNumberOfAPIRequests: 1000,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_google_deployment/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxGoogleDeploymentListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MjxGoogleDeploymentListResponse)
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
			result, err := service.ListMjxGoogleDeployments(t.Context(), tt.opts)

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

func TestService_GetMjxGoogleDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	oauthState := "test-state"

	expectedDeployment := &MjxGoogleDeployment{
		ID:                         1,
		Name:                       "test-google",
		Description:                "Test Google deployment",
		ClientEmail:                "test-service@project.iam.gserviceaccount.com",
		ClientID:                   "123456789012345678901",
		ClientSecret:               "secret123",
		PrivateKey:                 "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC...\n-----END PRIVATE KEY-----",
		UseUserConsent:             false,
		AuthEndpoint:               "https://accounts.google.com/o/oauth2/auth",
		TokenEndpoint:              "https://oauth2.googleapis.com/token",
		RedirectURI:                "https://pexip.example.com/oauth/google/callback",
		RefreshToken:               "refresh-token-123",
		OAuthState:                 &oauthState,
		MaximumNumberOfAPIRequests: 1000,
		MjxIntegrations:            []string{"/api/admin/configuration/v1/mjx_integration/1/", "/api/admin/configuration/v1/mjx_integration/2/"},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/mjx_google_deployment/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxGoogleDeployment")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxGoogleDeployment)
		*result = *expectedDeployment
	})

	service := New(client)
	result, err := service.GetMjxGoogleDeployment(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedDeployment, result)
	client.AssertExpectations(t)
}

func TestService_CreateMjxGoogleDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &MjxGoogleDeploymentCreateRequest{
		Name:                       "new-google",
		Description:                "New Google deployment",
		ClientEmail:                "new-service@project.iam.gserviceaccount.com",
		ClientID:                   "987654321098765432109",
		ClientSecret:               "newsecret123",
		PrivateKey:                 "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC...\n-----END PRIVATE KEY-----",
		UseUserConsent:             true,
		AuthEndpoint:               "https://accounts.google.com/o/oauth2/auth",
		TokenEndpoint:              "https://oauth2.googleapis.com/token",
		RedirectURI:                "https://new.example.com/oauth/google/callback",
		RefreshToken:               "new-refresh-token",
		MaximumNumberOfAPIRequests: 2000,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/mjx_google_deployment/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/mjx_google_deployment/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMjxGoogleDeployment(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMjxGoogleDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	useUserConsent := true
	maxRequests := 1500

	updateRequest := &MjxGoogleDeploymentUpdateRequest{
		Description:                "Updated Google deployment",
		UseUserConsent:             &useUserConsent,
		MaximumNumberOfAPIRequests: &maxRequests,
		RedirectURI:                "https://updated.example.com/oauth/google/callback",
	}

	oauthState := "test-state"
	expectedDeployment := &MjxGoogleDeployment{
		ID:                         1,
		Name:                       "test-google",
		Description:                "Updated Google deployment",
		ClientEmail:                "test-service@project.iam.gserviceaccount.com",
		ClientID:                   "123456789012345678901",
		UseUserConsent:             true,
		AuthEndpoint:               "https://accounts.google.com/o/oauth2/auth",
		TokenEndpoint:              "https://oauth2.googleapis.com/token",
		RedirectURI:                "https://updated.example.com/oauth/google/callback",
		OAuthState:                 &oauthState,
		MaximumNumberOfAPIRequests: 1500,
		MjxIntegrations:            []string{"/api/admin/configuration/v1/mjx_integration/1/"},
	}

	client.On("PutJSON", t.Context(), "configuration/v1/mjx_google_deployment/1/", updateRequest, mock.AnythingOfType("*config.MjxGoogleDeployment")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxGoogleDeployment)
		*result = *expectedDeployment
	})

	service := New(client)
	result, err := service.UpdateMjxGoogleDeployment(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedDeployment, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMjxGoogleDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/mjx_google_deployment/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMjxGoogleDeployment(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
