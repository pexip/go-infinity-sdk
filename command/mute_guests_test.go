package command

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_MuteGuests(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ConferenceMuteGuestsRequest{
		ConferenceID: "test-conference-id",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "All guests muted",
	}

	client.On("PostJSON", t.Context(), "command/v1/conference/mute_guests/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.MuteGuests(t.Context(), "test-conference-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UnmuteGuests(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ConferenceUnmuteGuestsRequest{
		ConferenceID: "test-conference-id",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "All guests unmuted",
	}

	client.On("PostJSON", t.Context(), "command/v1/conference/unmute_guests/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.UnmuteGuests(t.Context(), "test-conference-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}
