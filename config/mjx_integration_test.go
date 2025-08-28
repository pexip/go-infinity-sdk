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

func TestService_ListMjxIntegrations(t *testing.T) {
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
				exchangeDeployment := "/api/admin/configuration/v1/mjx_exchange_deployment/1/"
				googleDeployment := "/api/admin/configuration/v1/mjx_google_deployment/1/"
				graphDeployment := "/api/admin/configuration/v1/mjx_graph_deployment/1/"
				webexClientID := "webex-client-123"
				webexClientSecret := "webex-secret-456"
				webexOAuthState := "webex-state"
				webexRedirectURI := "https://pexip.example.com/oauth/webex/callback"
				webexRefreshToken := "webex-refresh-token"

				expectedResponse := &MjxIntegrationListResponse{
					Objects: []MjxIntegration{
						{
							ID:                          1,
							Name:                        "primary-mjx",
							Description:                 "Primary MJX integration",
							DisplayUpcomingMeetings:     5,
							EnableNonVideoMeetings:      true,
							EnablePrivateMeetings:       false,
							EndBuffer:                   5,
							StartBuffer:                 10,
							EPUsername:                  "mjx-user",
							EPUseHTTPS:                  true,
							EPVerifyCertificate:         true,
							ExchangeDeployment:          &exchangeDeployment,
							ProcessAliasPrivateMeetings: false,
							ReplaceEmptySubject:         true,
							ReplaceSubjectType:          "template",
							ReplaceSubjectTemplate:      "Meeting: {{subject}}",
							UseWebex:                    false,
							EndpointGroups:              []string{"/api/admin/configuration/v1/mjx_endpoint_group/1/"},
						},
						{
							ID:                          2,
							Name:                        "backup-mjx",
							Description:                 "Backup MJX integration",
							DisplayUpcomingMeetings:     3,
							EnableNonVideoMeetings:      false,
							EnablePrivateMeetings:       true,
							EndBuffer:                   3,
							StartBuffer:                 5,
							EPUsername:                  "backup-mjx-user",
							EPUseHTTPS:                  false,
							EPVerifyCertificate:         false,
							GoogleDeployment:            &googleDeployment,
							GraphDeployment:             &graphDeployment,
							ProcessAliasPrivateMeetings: true,
							ReplaceEmptySubject:         false,
							ReplaceSubjectType:          "none",
							UseWebex:                    true,
							WebexAPIDomain:              "webexapis.com",
							WebexClientID:               &webexClientID,
							WebexClientSecret:           &webexClientSecret,
							WebexOAuthState:             &webexOAuthState,
							WebexRedirectURI:            &webexRedirectURI,
							WebexRefreshToken:           &webexRefreshToken,
							EndpointGroups:              []string{"/api/admin/configuration/v1/mjx_endpoint_group/2/"},
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_integration/", mock.AnythingOfType("*config.MjxIntegrationListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MjxIntegrationListResponse)
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
				exchangeDeployment := "/api/admin/configuration/v1/mjx_exchange_deployment/1/"
				expectedResponse := &MjxIntegrationListResponse{
					Objects: []MjxIntegration{
						{
							ID:                      1,
							Name:                    "primary-mjx",
							Description:             "Primary MJX integration",
							DisplayUpcomingMeetings: 5,
							EnableNonVideoMeetings:  true,
							EnablePrivateMeetings:   false,
							ExchangeDeployment:      &exchangeDeployment,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_integration/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.MjxIntegrationListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MjxIntegrationListResponse)
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
			result, err := service.ListMjxIntegrations(t.Context(), tt.opts)

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

func TestService_GetMjxIntegration(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	exchangeDeployment := "/api/admin/configuration/v1/mjx_exchange_deployment/1/"
	googleDeployment := "/api/admin/configuration/v1/mjx_google_deployment/1/"
	graphDeployment := "/api/admin/configuration/v1/mjx_graph_deployment/1/"
	webexClientID := "webex-client-123"
	webexClientSecret := "webex-secret-456"
	webexOAuthState := "webex-state"
	webexRedirectURI := "https://pexip.example.com/oauth/webex/callback"
	webexRefreshToken := "webex-refresh-token"

	expectedIntegration := &MjxIntegration{
		ID:                          1,
		Name:                        "test-mjx",
		Description:                 "Test MJX integration",
		DisplayUpcomingMeetings:     5,
		EnableNonVideoMeetings:      true,
		EnablePrivateMeetings:       false,
		EndBuffer:                   5,
		StartBuffer:                 10,
		EPUsername:                  "mjx-test-user",
		EPPassword:                  "mjx-password123",
		EPUseHTTPS:                  true,
		EPVerifyCertificate:         true,
		ExchangeDeployment:          &exchangeDeployment,
		GoogleDeployment:            &googleDeployment,
		GraphDeployment:             &graphDeployment,
		ProcessAliasPrivateMeetings: false,
		ReplaceEmptySubject:         true,
		ReplaceSubjectType:          "template",
		ReplaceSubjectTemplate:      "Meeting: {{subject}}",
		UseWebex:                    true,
		WebexAPIDomain:              "webexapis.com",
		WebexClientID:               &webexClientID,
		WebexClientSecret:           &webexClientSecret,
		WebexOAuthState:             &webexOAuthState,
		WebexRedirectURI:            &webexRedirectURI,
		WebexRefreshToken:           &webexRefreshToken,
		EndpointGroups:              []string{"/api/admin/configuration/v1/mjx_endpoint_group/1/", "/api/admin/configuration/v1/mjx_endpoint_group/2/"},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/mjx_integration/1/", mock.AnythingOfType("*config.MjxIntegration")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MjxIntegration)
		*result = *expectedIntegration
	})

	service := New(client)
	result, err := service.GetMjxIntegration(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedIntegration, result)
	client.AssertExpectations(t)
}

func TestService_CreateMjxIntegration(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	exchangeDeployment := "/api/admin/configuration/v1/mjx_exchange_deployment/1/"
	webexClientID := "new-webex-client"
	webexClientSecret := "new-webex-secret"
	webexRedirectURI := "https://new.example.com/oauth/webex/callback"
	webexRefreshToken := "new-webex-refresh-token"

	createRequest := &MjxIntegrationCreateRequest{
		Name:                        "new-mjx",
		Description:                 "New MJX integration",
		DisplayUpcomingMeetings:     10,
		EnableNonVideoMeetings:      true,
		EnablePrivateMeetings:       true,
		EndBuffer:                   10,
		StartBuffer:                 15,
		EPUsername:                  "new-mjx-user",
		EPPassword:                  "new-mjx-password",
		EPUseHTTPS:                  true,
		EPVerifyCertificate:         true,
		ExchangeDeployment:          &exchangeDeployment,
		ProcessAliasPrivateMeetings: true,
		ReplaceEmptySubject:         true,
		ReplaceSubjectType:          "template",
		ReplaceSubjectTemplate:      "Conference: {{subject}}",
		UseWebex:                    true,
		WebexAPIDomain:              "webexapis.com",
		WebexClientID:               &webexClientID,
		WebexClientSecret:           &webexClientSecret,
		WebexRedirectURI:            &webexRedirectURI,
		WebexRefreshToken:           &webexRefreshToken,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/mjx_integration/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/mjx_integration/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMjxIntegration(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMjxIntegration(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	displayUpcomingMeetings := 7
	enableNonVideoMeetings := false
	endBuffer := 8
	epUseHTTPS := false
	useWebex := false

	updateRequest := &MjxIntegrationUpdateRequest{
		Description:             "Updated MJX integration",
		DisplayUpcomingMeetings: &displayUpcomingMeetings,
		EnableNonVideoMeetings:  &enableNonVideoMeetings,
		EndBuffer:               &endBuffer,
		EPUseHTTPS:              &epUseHTTPS,
		ReplaceSubjectTemplate:  "Updated: {{subject}}",
		UseWebex:                &useWebex,
	}

	exchangeDeployment := "/api/admin/configuration/v1/mjx_exchange_deployment/1/"
	expectedIntegration := &MjxIntegration{
		ID:                      1,
		Name:                    "test-mjx",
		Description:             "Updated MJX integration",
		DisplayUpcomingMeetings: 7,
		EnableNonVideoMeetings:  false,
		EnablePrivateMeetings:   false,
		EndBuffer:               8,
		StartBuffer:             10,
		EPUsername:              "mjx-test-user",
		EPUseHTTPS:              false,
		EPVerifyCertificate:     true,
		ExchangeDeployment:      &exchangeDeployment,
		ReplaceEmptySubject:     true,
		ReplaceSubjectType:      "template",
		ReplaceSubjectTemplate:  "Updated: {{subject}}",
		UseWebex:                false,
		EndpointGroups:          []string{"/api/admin/configuration/v1/mjx_endpoint_group/1/"},
	}

	client.On("PutJSON", t.Context(), "configuration/v1/mjx_integration/1/", updateRequest, mock.AnythingOfType("*config.MjxIntegration")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxIntegration)
		*result = *expectedIntegration
	})

	service := New(client)
	result, err := service.UpdateMjxIntegration(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedIntegration, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMjxIntegration(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/mjx_integration/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMjxIntegration(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
