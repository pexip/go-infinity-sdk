/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListConferences(t *testing.T) {
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
				expectedResponse := &ConferenceListResponse{
					Objects: []ConferenceStatus{
						{ID: "1", Name: "Test Conference 1", IsStarted: true, ServiceType: "conference"},
						{ID: "2", Name: "Test Conference 2", IsStarted: false, ServiceType: "conference"},
					},
				}
				m.On("GetJSON", t.Context(), "status/v1/conference/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.ConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ConferenceListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				Limit:  5,
				Offset: 10,
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &ConferenceListResponse{
					Objects: []ConferenceStatus{
						{ID: "3", Name: "Test Conference 3", IsStarted: true, ServiceType: "conference"},
					},
				}
				m.On("GetJSON", t.Context(), "status/v1/conference/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.ConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ConferenceListResponse)
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
			result, err := service.ListConferences(t.Context(), tt.opts)

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

func TestService_GetConference(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedConference := &ConferenceStatus{
		ID:          "1",
		Name:        "Test Conference",
		ServiceType: "conference",
		IsStarted:   true,
	}

	client.On("GetJSON", t.Context(), "status/v1/conference/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.ConferenceStatus")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ConferenceStatus)
		*result = *expectedConference
	})

	service := New(client)
	result, err := service.GetConference(t.Context(), "1")

	assert.NoError(t, err)
	assert.Equal(t, expectedConference, result)
	client.AssertExpectations(t)
}
