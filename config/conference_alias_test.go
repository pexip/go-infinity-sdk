/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListConferenceAliases(t *testing.T) {
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
				expectedResponse := &ConferenceAliasListResponse{
					Objects: []ConferenceAlias{
						{ID: 1, Alias: "test-alias", Conference: "/api/admin/configuration/v1/conference/1/"},
						{ID: 2, Alias: "another-alias", Conference: "/api/admin/configuration/v1/conference/2/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/conference_alias/", mock.AnythingOfType("*config.ConferenceAliasListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceAliasListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit:  5,
					Offset: 10,
				},
				Search: "test",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &ConferenceAliasListResponse{
					Objects: []ConferenceAlias{
						{ID: 1, Alias: "test-alias", Conference: "/api/admin/configuration/v1/conference/1/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/conference_alias/?limit=5&name__icontains=test&offset=10", mock.AnythingOfType("*config.ConferenceAliasListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceAliasListResponse)
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
			result, err := service.ListConferenceAliases(t.Context(), tt.opts)

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

func TestService_GetConferenceAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedAlias := &ConferenceAlias{
		ID:           1,
		Alias:        "test-alias",
		Conference:   "/api/admin/configuration/v1/conference/1/",
		Description:  "Test conference alias",
		CreationTime: util.InfinityTime{Time: time.Now()},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/conference_alias/1/", mock.AnythingOfType("*config.ConferenceAlias")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceAlias)
		*result = *expectedAlias
	})

	service := New(client)
	result, err := service.GetConferenceAlias(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedAlias, result)
	client.AssertExpectations(t)
}

func TestService_CreateConferenceAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &ConferenceAliasCreateRequest{
		Alias:       "new-alias",
		Conference:  "/api/admin/configuration/v1/conference/1/",
		Description: "New test alias",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/conference_alias/123/",
	}
	client.On("PostWithResponse", t.Context(), "configuration/v1/conference_alias/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateConferenceAlias(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateConferenceAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &ConferenceAliasUpdateRequest{
		Description: "Updated description",
	}

	expectedAlias := &ConferenceAlias{
		ID:          1,
		Alias:       "test-alias",
		Description: "Updated description",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/conference_alias/1/", updateRequest, mock.AnythingOfType("*config.ConferenceAlias")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ConferenceAlias)
		*result = *expectedAlias
	})

	service := New(client)
	result, err := service.UpdateConferenceAlias(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedAlias, result)
	client.AssertExpectations(t)
}

func TestService_DeleteConferenceAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/conference_alias/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteConferenceAlias(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
