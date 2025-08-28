/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package command

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_DisconnectParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ParticipantDisconnectRequestLegacy{
		ParticipantUUID: "test-uuid",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Participant disconnected",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/disconnect/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.DisconnectParticipant(t.Context(), "test-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_MuteParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ParticipantMuteRequestLegacy{
		ParticipantUUID: "test-uuid",
		Setting:         "mute",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Participant muted",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/mute/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.MuteParticipant(t.Context(), "test-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UnmuteParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ParticipantMuteRequestLegacy{
		ParticipantUUID: "test-uuid",
		Setting:         "unmute",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Participant unmuted",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/mute/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.UnmuteParticipant(t.Context(), "test-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_ToggleMuteParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ParticipantMuteRequestLegacy{
		ParticipantUUID: "test-uuid",
		Setting:         "toggle",
	}

	expectedResponse := &CommandResponse{
		Status: "success",
		Result: "muted",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/mute/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ToggleMuteParticipant(t.Context(), "test-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_SpotlightParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ParticipantSpotlightRequest{
		ParticipantUUID: "test-uuid",
		Setting:         "on",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Participant spotlighted",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/spotlight/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.SpotlightParticipant(t.Context(), "test-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UnspotlightParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ParticipantSpotlightRequest{
		ParticipantUUID: "test-uuid",
		Setting:         "off",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Participant unspotlighted",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/spotlight/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.UnspotlightParticipant(t.Context(), "test-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_ToggleSpotlightParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ParticipantSpotlightRequest{
		ParticipantUUID: "test-uuid",
		Setting:         "toggle",
	}

	expectedResponse := &CommandResponse{
		Status: "success",
		Result: "spotlighted",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/spotlight/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ToggleSpotlightParticipant(t.Context(), "test-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_SendMessageToParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ParticipantMessageRequest{
		ParticipantUUID: "test-uuid",
		Message:         "Hello, participant!",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Message sent",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/message/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.SendMessageToParticipant(t.Context(), "test-uuid", "Hello, participant!")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_TransferParticipant(t *testing.T) {
	tests := []struct {
		name            string
		participantUUID string
		conferenceAlias string
		opts            *TransferOptions
		expectedRequest *ParticipantTransferRequestLegacy
		wantErr         bool
	}{
		{
			name:            "transfer without options",
			participantUUID: "test-uuid",
			conferenceAlias: "new-conference",
			opts:            nil,
			expectedRequest: &ParticipantTransferRequestLegacy{
				ParticipantUUID: "test-uuid",
				ConferenceAlias: "new-conference",
			},
			wantErr: false,
		},
		{
			name:            "transfer with options",
			participantUUID: "test-uuid",
			conferenceAlias: "new-conference",
			opts: &TransferOptions{
				Role: "chair",
				PIN:  "1234",
			},
			expectedRequest: &ParticipantTransferRequestLegacy{
				ParticipantUUID: "test-uuid",
				ConferenceAlias: "new-conference",
				Role:            "chair",
				PIN:             "1234",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := interfaces.NewHTTPClientMock()

			expectedResponse := &CommandResponse{
				Status:  "success",
				Message: "Participant transferred",
			}

			client.On("PostJSON", t.Context(), "command/v1/participant/transfer/", mock.MatchedBy(func(req *ParticipantTransferRequestLegacy) bool {
				return req.ParticipantUUID == tt.expectedRequest.ParticipantUUID &&
					req.ConferenceAlias == tt.expectedRequest.ConferenceAlias
			}), mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
				result := args.Get(3).(*CommandResponse)
				*result = *expectedResponse
			})

			service := New(client)
			result, err := service.TransferParticipant(t.Context(), tt.participantUUID, tt.conferenceAlias, tt.opts)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, result)
			}

			client.AssertExpectations(t)
		})
	}
}

func TestService_PromoteParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ParticipantRoleRequestLegacy{
		ParticipantUUID: "test-uuid",
		Role:            "chair",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Participant promoted to chair",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/role/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.PromoteParticipant(t.Context(), "test-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_DemoteParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ParticipantRoleRequestLegacy{
		ParticipantUUID: "test-uuid",
		Role:            "guest",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Participant demoted to guest",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/role/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.DemoteParticipant(t.Context(), "test-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}
