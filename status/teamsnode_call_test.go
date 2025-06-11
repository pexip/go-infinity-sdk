package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListTeamsNodeCalls(t *testing.T) {
	client := &mockClient.Client{}

	startTime := time.Now().Add(-30 * time.Minute)

	expectedResponse := &TeamsNodeCallListResponse{
		Objects: []TeamsNodeCall{
			{
				ID:              "call-123",
				TeamsNodeID:     1,
				ConferenceName:  "Team Meeting",
				ParticipantName: "John Doe",
				CallDirection:   "inbound",
				StartTime:       &startTime,
				Duration:        1800,
				Status:          "active",
				ResourceURI:     "/api/admin/status/v1/teamsnode_call/call-123/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/teamsnode_call/", mock.AnythingOfType("*status.TeamsNodeCallListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*TeamsNodeCallListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListTeamsNodeCalls(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "call-123", result.Objects[0].ID)
	assert.Equal(t, "Team Meeting", result.Objects[0].ConferenceName)
	assert.Equal(t, "John Doe", result.Objects[0].ParticipantName)
	assert.Equal(t, "inbound", result.Objects[0].CallDirection)
	assert.Equal(t, "active", result.Objects[0].Status)
	client.AssertExpectations(t)
}

func TestService_ListTeamsNodeCalls_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  20,
		Offset: 10,
	}

	startTime := time.Now().Add(-15 * time.Minute)
	expectedResponse := &TeamsNodeCallListResponse{
		Objects: []TeamsNodeCall{
			{
				ID:              "call-options-test",
				TeamsNodeID:     3,
				ConferenceName:  "Options Test Meeting",
				ParticipantName: "Test User",
				CallDirection:   "inbound",
				StartTime:       &startTime,
				Duration:        900,
				Status:          "active",
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/teamsnode_call/"
	}), mock.AnythingOfType("*status.TeamsNodeCallListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*TeamsNodeCallListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListTeamsNodeCalls(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Options Test Meeting", result.Objects[0].ConferenceName)

	client.AssertExpectations(t)
}

func TestService_GetTeamsNodeCall(t *testing.T) {
	client := &mockClient.Client{}

	startTime := time.Now().Add(-1 * time.Hour)
	expectedCall := &TeamsNodeCall{
		ID:              "call-456",
		TeamsNodeID:     2,
		ConferenceName:  "Executive Meeting",
		ParticipantName: "Jane Smith",
		CallDirection:   "outbound",
		StartTime:       &startTime,
		Duration:        3600,
		Status:          "completed",
		ResourceURI:     "/api/admin/status/v1/teamsnode_call/call-456/",
	}

	client.On("GetJSON", t.Context(), "status/v1/teamsnode_call/call-456/", mock.AnythingOfType("*status.TeamsNodeCall")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*TeamsNodeCall)
		*result = *expectedCall
	})

	service := New(client)
	result, err := service.GetTeamsNodeCall(t.Context(), "call-456")

	assert.NoError(t, err)
	assert.Equal(t, expectedCall, result)
	assert.Equal(t, "Executive Meeting", result.ConferenceName)
	assert.Equal(t, "completed", result.Status)
	client.AssertExpectations(t)
}