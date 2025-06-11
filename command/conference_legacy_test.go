package command

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_LockConference(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ConferenceLockRequestLegacy{
		ConferenceID: 1,
		Setting:      "lock",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Conference locked",
	}

	client.On("PostJSON", t.Context(), "command/v1/conference/lock/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.LockConference(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UnlockConference(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ConferenceLockRequestLegacy{
		ConferenceID: 1,
		Setting:      "unlock",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Conference unlocked",
	}

	client.On("PostJSON", t.Context(), "command/v1/conference/lock/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.UnlockConference(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_SendMessageToConference(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ConferenceMessageRequest{
		ConferenceID: 1,
		Message:      "Hello, everyone!",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Message sent to conference",
	}

	client.On("PostJSON", t.Context(), "command/v1/conference/message/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.SendMessageToConference(t.Context(), 1, "Hello, everyone!")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_StartConference(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ConferenceStartRequest{
		ConferenceAlias: "test-conference",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Conference started",
	}

	client.On("PostJSON", t.Context(), "command/v1/conference/start/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.StartConference(t.Context(), "test-conference")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_ToggleLockConference(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ConferenceLockRequestLegacy{
		ConferenceID: 1,
		Setting:      "toggle",
	}

	expectedResponse := &CommandResponse{
		Status: "success",
		Result: "locked",
	}

	client.On("PostJSON", t.Context(), "command/v1/conference/lock/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ToggleLockConference(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_StopConference(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ConferenceStopRequest{
		ConferenceID: 1,
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Conference stopped",
	}

	client.On("PostJSON", t.Context(), "command/v1/conference/stop/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.StopConference(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}
