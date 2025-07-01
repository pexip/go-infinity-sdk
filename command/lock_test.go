package command

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_LockConferenceByID(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ConferenceLockRequest{
		ConferenceID: "test-conference-id",
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
	result, err := service.LockConferenceByID(t.Context(), "test-conference-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UnlockConferenceByID(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRequest := &ConferenceUnlockRequest{
		ConferenceID: "test-conference-id",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Conference unlocked",
	}

	client.On("PostJSON", t.Context(), "command/v1/conference/unlock/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.UnlockConferenceByID(t.Context(), "test-conference-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}
