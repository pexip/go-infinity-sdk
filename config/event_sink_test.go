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

func TestService_ListEventSinks(t *testing.T) {
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
				desc1 := "Primary event sink"
				desc2 := "Secondary event sink"
				expectedResponse := &EventSinkListResponse{
					Objects: []EventSink{
						{ID: 1, Name: "primary-sink", Description: &desc1, URL: "https://events.example.com/webhook", BulkSupport: true, VerifyTLSCertificate: true, Version: 2},
						{ID: 2, Name: "secondary-sink", Description: &desc2, URL: "https://backup.example.com/events", BulkSupport: false, VerifyTLSCertificate: true, Version: 2},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/event_sink/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.EventSinkListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*EventSinkListResponse)
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
				desc := "Primary event sink"
				expectedResponse := &EventSinkListResponse{
					Objects: []EventSink{
						{ID: 1, Name: "primary-sink", Description: &desc, URL: "https://events.example.com/webhook", BulkSupport: true, VerifyTLSCertificate: true, Version: 2},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/event_sink/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.EventSinkListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*EventSinkListResponse)
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
			result, err := service.ListEventSinks(t.Context(), tt.opts)

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

func TestService_GetEventSink(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	description := "Test event sink"
	username := "testuser"
	password := "testpass"

	expectedEventSink := &EventSink{
		ID:                   1,
		Name:                 "test-sink",
		Description:          &description,
		URL:                  "https://events.example.com/webhook",
		Username:             &username,
		Password:             &password,
		BulkSupport:          true,
		VerifyTLSCertificate: true,
		Version:              2,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/event_sink/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.EventSink")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*EventSink)
		*result = *expectedEventSink
	})

	service := New(client)
	result, err := service.GetEventSink(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedEventSink, result)
	client.AssertExpectations(t)
}

func TestService_CreateEventSink(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	description := "New event sink"
	username := "newuser"
	password := "newpass"

	createRequest := &EventSinkCreateRequest{
		Name:                 "new-sink",
		Description:          &description,
		URL:                  "https://new.example.com/events",
		Username:             &username,
		Password:             &password,
		BulkSupport:          true,
		VerifyTLSCertificate: true,
		Version:              2,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/event_sink/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/event_sink/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateEventSink(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateEventSink(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	newDescription := "Updated event sink"
	bulkSupport := false

	updateRequest := &EventSinkUpdateRequest{
		Description: &newDescription,
		BulkSupport: &bulkSupport,
	}

	expectedEventSink := &EventSink{
		ID:                   1,
		Name:                 "test-sink",
		Description:          &newDescription,
		URL:                  "https://events.example.com/webhook",
		BulkSupport:          false,
		VerifyTLSCertificate: true,
		Version:              2,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/event_sink/1/", updateRequest, mock.AnythingOfType("*config.EventSink")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*EventSink)
		*result = *expectedEventSink
	})

	service := New(client)
	result, err := service.UpdateEventSink(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedEventSink, result)
	client.AssertExpectations(t)
}

func TestService_DeleteEventSink(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/event_sink/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteEventSink(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
