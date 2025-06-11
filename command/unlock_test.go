package command

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_UnlockParticipant(t *testing.T) {
	client := &mockClient.Client{}

	expectedRequest := &ParticipantUnlockRequest{
		ParticipantID: "test-participant-id",
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Participant unlocked",
	}

	client.On("PostJSON", t.Context(), "command/v1/participant/unlock/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.UnlockParticipant(t.Context(), "test-participant-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}
