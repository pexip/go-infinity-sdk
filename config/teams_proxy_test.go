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
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListTeamsProxies(t *testing.T) {
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
				eventhubID1 := "eh-primary"
				notificationsQueue1 := "queue-primary"
				updated1 := &util.InfinityTime{}

				expectedResponse := &TeamsProxyListResponse{
					Objects: []TeamsProxy{
						{ID: 1, Name: "primary-teams-proxy", Description: "Primary Teams proxy", Address: "teams-proxy.example.com", Port: 8443, AzureTenant: "tenant1", EventhubID: &eventhubID1, MinNumberOfInstances: 2, NotificationsEnabled: true, NotificationsQueue: &notificationsQueue1, Updated: updated1},
						{ID: 2, Name: "backup-teams-proxy", Description: "Backup Teams proxy", Address: "teams-proxy-backup.example.com", Port: 8443, AzureTenant: "tenant2", MinNumberOfInstances: 1, NotificationsEnabled: false},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/teams_proxy/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.TeamsProxyListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*TeamsProxyListResponse)
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
				eventhubID := "eh-primary"
				notificationsQueue := "queue-primary"
				updated := &util.InfinityTime{}

				expectedResponse := &TeamsProxyListResponse{
					Objects: []TeamsProxy{
						{ID: 1, Name: "primary-teams-proxy", Description: "Primary Teams proxy", Address: "teams-proxy.example.com", Port: 8443, AzureTenant: "tenant1", EventhubID: &eventhubID, MinNumberOfInstances: 2, NotificationsEnabled: true, NotificationsQueue: &notificationsQueue, Updated: updated},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/teams_proxy/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.TeamsProxyListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*TeamsProxyListResponse)
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
			result, err := service.ListTeamsProxies(t.Context(), tt.opts)

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

func TestService_GetTeamsProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	eventhubID := "eh-test"
	notificationsQueue := "queue-test"
	updated := &util.InfinityTime{}

	expectedTeamsProxy := &TeamsProxy{
		ID:                   1,
		Name:                 "test-teams-proxy",
		Description:          "Test Teams proxy",
		Address:              "teams-proxy.example.com",
		Port:                 8443,
		AzureTenant:          "test-tenant",
		EventhubID:           &eventhubID,
		MinNumberOfInstances: 2,
		NotificationsEnabled: true,
		NotificationsQueue:   &notificationsQueue,
		Updated:              updated,
		ResourceURI:          "/api/admin/configuration/v1/teams_proxy/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/teams_proxy/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.TeamsProxy")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TeamsProxy)
		*result = *expectedTeamsProxy
	})

	service := New(client)
	result, err := service.GetTeamsProxy(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedTeamsProxy, result)
	client.AssertExpectations(t)
}

func TestService_CreateTeamsProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	eventhubID := "eh-new"
	notificationsQueue := "queue-new"

	createRequest := &TeamsProxyCreateRequest{
		Name:                 "new-teams-proxy",
		Description:          "New Teams proxy",
		Address:              "new-teams-proxy.example.com",
		Port:                 8443,
		AzureTenant:          "new-tenant",
		EventhubID:           &eventhubID,
		MinNumberOfInstances: 3,
		NotificationsEnabled: true,
		NotificationsQueue:   &notificationsQueue,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/teams_proxy/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/teams_proxy/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateTeamsProxy(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateTeamsProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &TeamsProxyUpdateRequest{
		Description:          "Updated Teams proxy",
		Port:                 8444,
		MinNumberOfInstances: 4,
		NotificationsEnabled: false,
	}

	updated := &util.InfinityTime{}
	expectedTeamsProxy := &TeamsProxy{
		ID:                   1,
		Name:                 "test-teams-proxy",
		Description:          "Updated Teams proxy",
		Address:              "teams-proxy.example.com",
		Port:                 8444,
		AzureTenant:          "test-tenant",
		MinNumberOfInstances: 4,
		NotificationsEnabled: false,
		Updated:              updated,
		ResourceURI:          "/api/admin/configuration/v1/teams_proxy/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/teams_proxy/1/", updateRequest, mock.AnythingOfType("*config.TeamsProxy")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TeamsProxy)
		*result = *expectedTeamsProxy
	})

	service := New(client)
	result, err := service.UpdateTeamsProxy(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedTeamsProxy, result)
	client.AssertExpectations(t)
}

func TestService_DeleteTeamsProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/teams_proxy/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteTeamsProxy(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
