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

func TestService_ListAutomaticParticipants(t *testing.T) {
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
				expectedResponse := &AutomaticParticipantListResponse{
					Objects: []AutomaticParticipant{
						{ID: 1, Alias: "auto-participant-1", Description: "Primary participant", Conference: "/api/admin/configuration/v1/conference/1/", Protocol: "sip", CallType: "video", Role: "chair", KeepConferenceAlive: "keep_conference_alive_never", Routing: "routing_rule", Streaming: false},
						{ID: 2, Alias: "auto-participant-2", Description: "Secondary participant", Conference: "/api/admin/configuration/v1/conference/2/", Protocol: "h323", CallType: "audio", Role: "guest", KeepConferenceAlive: "keep_conference_alive_if_multiple", Routing: "routing_rule", Streaming: true},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/automatic_participant/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.AutomaticParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*AutomaticParticipantListResponse)
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
				expectedResponse := &AutomaticParticipantListResponse{
					Objects: []AutomaticParticipant{
						{ID: 1, Alias: "auto-participant-1", Description: "Primary participant", Conference: "/api/admin/configuration/v1/conference/1/", Protocol: "sip", CallType: "video", Role: "chair", KeepConferenceAlive: "keep_conference_alive_never", Routing: "routing_rule", Streaming: false},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/automatic_participant/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.AutomaticParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*AutomaticParticipantListResponse)
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
			result, err := service.ListAutomaticParticipants(t.Context(), tt.opts)

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

func TestService_GetAutomaticParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	systemLocation := "/api/admin/configuration/v1/system_location/1/"
	expectedAutomaticParticipant := &AutomaticParticipant{
		ID:                  1,
		Alias:               "test-participant",
		Description:         "Test automatic participant",
		Conference:          "/api/admin/configuration/v1/conference/1/",
		Protocol:            "sip",
		CallType:            "video",
		Role:                "chair",
		DTMFSequence:        "123#",
		KeepConferenceAlive: "keep_conference_alive_never",
		Routing:             "routing_rule",
		SystemLocation:      &systemLocation,
		Streaming:           false,
		RemoteDisplayName:   "Test Participant",
		PresentationURL:     "https://example.com/presentation.pdf",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/automatic_participant/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.AutomaticParticipant")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*AutomaticParticipant)
		*result = *expectedAutomaticParticipant
	})

	service := New(client)
	result, err := service.GetAutomaticParticipant(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedAutomaticParticipant, result)
	client.AssertExpectations(t)
}

func TestService_CreateAutomaticParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	systemLocation := "/api/admin/configuration/v1/system_location/1/"
	createRequest := &AutomaticParticipantCreateRequest{
		Alias:               "new-participant",
		Description:         "New automatic participant",
		Conference:          "/api/admin/configuration/v1/conference/1/",
		Protocol:            "sip",
		CallType:            "video",
		Role:                "guest",
		DTMFSequence:        "456#",
		KeepConferenceAlive: "keep_conference_alive_if_multiple",
		Routing:             "routing_rule",
		SystemLocation:      &systemLocation,
		Streaming:           true,
		RemoteDisplayName:   "New Participant",
		PresentationURL:     "https://new.example.com/presentation.pdf",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/automatic_participant/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/automatic_participant/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateAutomaticParticipant(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateAutomaticParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	streaming := true
	updateRequest := &AutomaticParticipantUpdateRequest{
		Description:       "Updated automatic participant",
		Role:              "chair",
		Streaming:         &streaming,
		RemoteDisplayName: "Updated Participant",
	}

	systemLocation := "/api/admin/configuration/v1/system_location/1/"
	expectedAutomaticParticipant := &AutomaticParticipant{
		ID:                  1,
		Alias:               "test-participant",
		Description:         "Updated automatic participant",
		Conference:          "/api/admin/configuration/v1/conference/1/",
		Protocol:            "sip",
		CallType:            "video",
		Role:                "chair",
		DTMFSequence:        "123#",
		KeepConferenceAlive: "keep_conference_alive_never",
		Routing:             "routing_rule",
		SystemLocation:      &systemLocation,
		Streaming:           true,
		RemoteDisplayName:   "Updated Participant",
		PresentationURL:     "https://example.com/presentation.pdf",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/automatic_participant/1/", updateRequest, mock.AnythingOfType("*config.AutomaticParticipant")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*AutomaticParticipant)
		*result = *expectedAutomaticParticipant
	})

	service := New(client)
	result, err := service.UpdateAutomaticParticipant(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedAutomaticParticipant, result)
	client.AssertExpectations(t)
}

func TestService_DeleteAutomaticParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/automatic_participant/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteAutomaticParticipant(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
