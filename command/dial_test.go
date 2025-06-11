package command

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_DialParticipant(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ParticipantDialRequest{
		ConferenceAlias: "test-conference",
		Destination:     "john@example.com",
		CallType:        "video",
		Protocol:        "sip",
		Role:            "guest",
		Streaming:       false,
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Participant dialed successfully",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/dial/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.DialParticipant(t.Context(), expectedRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_DialParticipantWithOptions(t *testing.T) {
	tests := []struct {
		name            string
		conferenceAlias string
		destination     string
		opts            *DialOptions
		wantErr         bool
	}{
		{
			name:            "dial without options",
			conferenceAlias: "test-conference",
			destination:     "jane@example.com",
			opts:            nil,
			wantErr:         false,
		},
		{
			name:            "dial with basic options",
			conferenceAlias: "test-conference",
			destination:     "bob@example.com",
			opts: &DialOptions{
				CallType: "audio",
				Protocol: "sip",
				Role:     "chair",
			},
			wantErr: false,
		},
		{
			name:            "dial with advanced options",
			conferenceAlias: "streaming-conference",
			destination:     "rtmp://streaming.example.com/live",
			opts: &DialOptions{
				CallType:            "video",
				Protocol:            "rtmp",
				Role:                "guest",
				LocalDisplayName:    "Conference Stream",
				RemoteDisplayName:   "RTMP Stream",
				Streaming:           true,
				KeepConferenceAlive: "keep_conference_alive",
				Routing:             "manual",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockClient.Client{}

			expectedResponse := &CommandResponse{
				Status:  "success",
				Message: "Participant dialed successfully",
			}

			client.On("PostJSON", t.Context(), "command/v1/participant/dial/", mock.AnythingOfType("*command.ParticipantDialRequest"), mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
				result := args.Get(3).(*CommandResponse)
				*result = *expectedResponse
			})

			service := New(client)
			result, err := service.DialParticipantWithOptions(t.Context(), tt.conferenceAlias, tt.destination, tt.opts)

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
