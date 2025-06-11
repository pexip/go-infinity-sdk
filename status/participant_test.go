package status

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListParticipants(t *testing.T) {
	client := &mockClient.Client{}

	expectedResponse := &ParticipantListResponse{
		Objects: []Participant{
			{
				ID:          "participant-1",
				UUID:        "uuid-1",
				DisplayName: "John Doe",
				Role:        "chair",
				IsMuted:     false,
			},
			{
				ID:          "participant-2",
				UUID:        "uuid-2",
				DisplayName: "Jane Smith",
				Role:        "guest",
				IsMuted:     true,
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/participant/", mock.AnythingOfType("*status.ParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ParticipantListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListParticipants(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "John Doe", result.Objects[0].DisplayName)
	assert.Equal(t, "chair", result.Objects[0].Role)
	client.AssertExpectations(t)
}

func TestService_ListParticipants_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  20,
		Offset: 10,
	}

	expectedResponse := &ParticipantListResponse{
		Objects: []Participant{
			{
				ID:          "participant-3",
				UUID:        "uuid-3",
				DisplayName: "Alice Smith",
				Role:        "guest",
				IsMuted:     false,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/participant/"
	}), mock.AnythingOfType("*status.ParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ParticipantListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListParticipants(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Alice Smith", result.Objects[0].DisplayName)

	client.AssertExpectations(t)
}

func TestService_GetParticipant(t *testing.T) {
	client := &mockClient.Client{}
	expectedParticipant := &Participant{
		ID:             "participant-1",
		UUID:           "test-uuid",
		DisplayName:    "John Doe",
		Role:           "chair",
		IsMuted:        false,
		IsPresenting:   true,
		ConferenceID:   1,
		ConferenceName: "Test Conference",
	}

	client.On("GetJSON", t.Context(), "status/v1/participant/test-uuid/", mock.AnythingOfType("*status.Participant")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*Participant)
		*result = *expectedParticipant
	})

	service := New(client)
	result, err := service.GetParticipant(t.Context(), "test-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedParticipant, result)
	client.AssertExpectations(t)
}
