package command

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_DisconnectParticipantByID(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ParticipantDisconnectRequest{
		ParticipantID: "test-participant-id",
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
	result, err := service.DisconnectParticipantByID(t.Context(), "test-participant-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_MuteParticipantByID(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ParticipantMuteRequest{
		ParticipantID: "test-participant-id",
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
	result, err := service.MuteParticipantByID(t.Context(), "test-participant-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UnmuteParticipantByID(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ParticipantUnmuteRequest{
		ParticipantID: "test-participant-id",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Participant unmuted",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/unmute/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.UnmuteParticipantByID(t.Context(), "test-participant-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_ChangeParticipantRoleByID(t *testing.T) {
	tests := []struct {
		name          string
		participantID string
		role          string
		wantErr       bool
	}{
		{
			name:          "promote to chair",
			participantID: "test-participant-id",
			role:          "chair",
			wantErr:       false,
		},
		{
			name:          "demote to guest",
			participantID: "test-participant-id",
			role:          "guest",
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockClient.Client{}

			expectedRequest := &ParticipantRoleRequest{
				ParticipantID: tt.participantID,
				Role:          tt.role,
			}

			expectedResponse := &CommandResponse{
				Status:  "success",
				Message: "Participant role changed",
			}

			client.On("PostJSON", t.Context(), "command/v1/participant/role/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
				result := args.Get(3).(*CommandResponse)
				*result = *expectedResponse
			})

			service := New(client)
			result, err := service.ChangeParticipantRoleByID(t.Context(), tt.participantID, tt.role)

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

func TestService_TransferParticipantByID(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ParticipantTransferRequest{
		ParticipantID:   "test-participant-id",
		ConferenceAlias: "target-conference",
		Role:            "guest",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Participant transferred",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/transfer/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.TransferParticipantByID(t.Context(), "test-participant-id", "target-conference", "guest")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_PromoteParticipantByID(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ParticipantRoleRequest{
		ParticipantID: "test-participant-id",
		Role:          "chair",
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
	result, err := service.PromoteParticipantByID(t.Context(), "test-participant-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_DemoteParticipantByID(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ParticipantRoleRequest{
		ParticipantID: "test-participant-id",
		Role:          "guest",
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
	result, err := service.DemoteParticipantByID(t.Context(), "test-participant-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}
