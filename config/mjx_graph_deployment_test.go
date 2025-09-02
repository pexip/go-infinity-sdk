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

func TestService_ListMjxGraphDeployments(t *testing.T) {
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
				expectedResponse := &MjxGraphDeploymentListResponse{
					Objects: []MjxGraphDeployment{
						{
							ID:              1,
							Name:            "primary-graph",
							Description:     "Primary Microsoft Graph deployment",
							ClientID:        "12345678-1234-1234-1234-123456789012",
							OAuthTokenURL:   "https://login.microsoftonline.com/tenant/oauth2/v2.0/token",
							GraphAPIDomain:  "graph.microsoft.com",
							RequestQuota:    1000,
							DisableProxy:    false,
							MjxIntegrations: []string{"/api/admin/configuration/v1/mjx_integration/1/"},
						},
						{
							ID:              2,
							Name:            "backup-graph",
							Description:     "Backup Microsoft Graph deployment",
							ClientID:        "87654321-4321-4321-4321-210987654321",
							OAuthTokenURL:   "https://login.microsoftonline.com/tenant2/oauth2/v2.0/token",
							GraphAPIDomain:  "graph.microsoft.com",
							RequestQuota:    500,
							DisableProxy:    true,
							MjxIntegrations: []string{"/api/admin/configuration/v1/mjx_integration/2/"},
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_graph_deployment/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxGraphDeploymentListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MjxGraphDeploymentListResponse)
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
				expectedResponse := &MjxGraphDeploymentListResponse{
					Objects: []MjxGraphDeployment{
						{
							ID:             1,
							Name:           "primary-graph",
							Description:    "Primary Microsoft Graph deployment",
							ClientID:       "12345678-1234-1234-1234-123456789012",
							OAuthTokenURL:  "https://login.microsoftonline.com/tenant/oauth2/v2.0/token",
							GraphAPIDomain: "graph.microsoft.com",
							RequestQuota:   1000,
							DisableProxy:   false,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_graph_deployment/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxGraphDeploymentListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MjxGraphDeploymentListResponse)
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
			result, err := service.ListMjxGraphDeployments(t.Context(), tt.opts)

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

func TestService_GetMjxGraphDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedDeployment := &MjxGraphDeployment{
		ID:             1,
		Name:           "test-graph",
		Description:    "Test Microsoft Graph deployment",
		ClientID:       "12345678-1234-1234-1234-123456789012",
		ClientSecret:   "client-secret-123",
		OAuthTokenURL:  "https://login.microsoftonline.com/tenant/oauth2/v2.0/token",
		GraphAPIDomain: "graph.microsoft.com",
		RequestQuota:   1000,
		DisableProxy:   false,
		MjxIntegrations: []string{
			"/api/admin/configuration/v1/mjx_integration/1/",
			"/api/admin/configuration/v1/mjx_integration/2/",
		},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/mjx_graph_deployment/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MjxGraphDeployment")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxGraphDeployment)
		*result = *expectedDeployment
	})

	service := New(client)
	result, err := service.GetMjxGraphDeployment(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedDeployment, result)
	client.AssertExpectations(t)
}

func TestService_CreateMjxGraphDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &MjxGraphDeploymentCreateRequest{
		Name:           "new-graph",
		Description:    "New Microsoft Graph deployment",
		ClientID:       "87654321-4321-4321-4321-210987654321",
		ClientSecret:   "new-client-secret",
		OAuthTokenURL:  "https://login.microsoftonline.com/newtenant/oauth2/v2.0/token",
		GraphAPIDomain: "graph.microsoft.com",
		RequestQuota:   2000,
		DisableProxy:   false,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/mjx_graph_deployment/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/mjx_graph_deployment/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMjxGraphDeployment(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMjxGraphDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	requestQuota := 1500
	disableProxy := true

	updateRequest := &MjxGraphDeploymentUpdateRequest{
		Description:   "Updated Microsoft Graph deployment",
		RequestQuota:  &requestQuota,
		DisableProxy:  &disableProxy,
		OAuthTokenURL: "https://login.microsoftonline.com/updatedtenant/oauth2/v2.0/token",
	}

	expectedDeployment := &MjxGraphDeployment{
		ID:              1,
		Name:            "test-graph",
		Description:     "Updated Microsoft Graph deployment",
		ClientID:        "12345678-1234-1234-1234-123456789012",
		ClientSecret:    "client-secret-123",
		OAuthTokenURL:   "https://login.microsoftonline.com/updatedtenant/oauth2/v2.0/token",
		GraphAPIDomain:  "graph.microsoft.com",
		RequestQuota:    1500,
		DisableProxy:    true,
		MjxIntegrations: []string{"/api/admin/configuration/v1/mjx_integration/1/"},
	}

	client.On("PutJSON", t.Context(), "configuration/v1/mjx_graph_deployment/1/", updateRequest, mock.AnythingOfType("*config.MjxGraphDeployment")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxGraphDeployment)
		*result = *expectedDeployment
	})

	service := New(client)
	result, err := service.UpdateMjxGraphDeployment(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedDeployment, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMjxGraphDeployment(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/mjx_graph_deployment/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMjxGraphDeployment(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
