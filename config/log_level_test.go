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

func TestService_ListLogLevels(t *testing.T) {
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
				expectedResponse := &LogLevelListResponse{
					Objects: []LogLevel{
						{ID: 1, Name: "media_processing", Level: "INFO"},
						{ID: 2, Name: "conference_manager", Level: "DEBUG"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/log_level/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.LogLevelListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*LogLevelListResponse)
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
				Search: "media",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &LogLevelListResponse{
					Objects: []LogLevel{
						{ID: 1, Name: "media_processing", Level: "INFO"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/log_level/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.LogLevelListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*LogLevelListResponse)
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
			result, err := service.ListLogLevels(t.Context(), tt.opts)

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

func TestService_GetLogLevel(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedLogLevel := &LogLevel{
		ID:    1,
		Name:  "test_service",
		Level: "WARN",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/log_level/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.LogLevel")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*LogLevel)
		*result = *expectedLogLevel
	})

	service := New(client)
	result, err := service.GetLogLevel(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedLogLevel, result)
	client.AssertExpectations(t)
}

func TestService_CreateLogLevel(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &LogLevelCreateRequest{
		Name:  "new_service",
		Level: "ERROR",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/log_level/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/log_level/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateLogLevel(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateLogLevel(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &LogLevelUpdateRequest{
		Level: "DEBUG",
	}

	expectedLogLevel := &LogLevel{
		ID:    1,
		Name:  "test_service",
		Level: "DEBUG",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/log_level/1/", updateRequest, mock.AnythingOfType("*config.LogLevel")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*LogLevel)
		*result = *expectedLogLevel
	})

	service := New(client)
	result, err := service.UpdateLogLevel(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedLogLevel, result)
	client.AssertExpectations(t)
}

func TestService_DeleteLogLevel(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/log_level/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteLogLevel(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
