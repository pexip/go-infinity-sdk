package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListMJXMeetings(t *testing.T) {
	client := &mockClient.Client{}

	startTime := time.Now().Add(-1 * time.Hour)
	endTime := time.Now().Add(1 * time.Hour)

	expectedResponse := &MJXMeetingListResponse{
		Objects: []MJXMeeting{
			{
				ID:               "meeting-123",
				Subject:          "Weekly Team Meeting",
				Organizer:        "john.doe@company.com",
				StartTime:        &startTime,
				EndTime:          &endTime,
				Status:           "active",
				ParticipantCount: 8,
				ConferenceAlias:  "team-weekly",
				ResourceURI:      "/api/admin/status/v1/mjx_meeting/meeting-123/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/mjx_meeting/", mock.AnythingOfType("*status.MJXMeetingListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MJXMeetingListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListMJXMeetings(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Weekly Team Meeting", result.Objects[0].Subject)
	assert.Equal(t, "john.doe@company.com", result.Objects[0].Organizer)
	assert.Equal(t, "active", result.Objects[0].Status)
	assert.Equal(t, 8, result.Objects[0].ParticipantCount)
	client.AssertExpectations(t)
}

func TestService_ListMJXMeetings_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  15,
		Offset: 5,
	}

	startTime := time.Now().Add(-30 * time.Minute)
	endTime := time.Now().Add(30 * time.Minute)
	expectedResponse := &MJXMeetingListResponse{
		Objects: []MJXMeeting{
			{
				ID:               "meeting-options-test",
				Subject:          "Test Meeting With Options",
				Organizer:        "test@company.com",
				StartTime:        &startTime,
				EndTime:          &endTime,
				Status:           "active",
				ParticipantCount: 5,
				ConferenceAlias:  "test-options",
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/mjx_meeting/"
	}), mock.AnythingOfType("*status.MJXMeetingListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MJXMeetingListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListMJXMeetings(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Test Meeting With Options", result.Objects[0].Subject)

	client.AssertExpectations(t)
}

func TestService_GetMJXMeeting(t *testing.T) {
	client := &mockClient.Client{}

	startTime := time.Now().Add(-2 * time.Hour)
	endTime := time.Now().Add(-1 * time.Hour)
	expectedMeeting := &MJXMeeting{
		ID:               "meeting-456",
		Subject:          "Board Meeting",
		Organizer:        "ceo@company.com",
		StartTime:        &startTime,
		EndTime:          &endTime,
		Status:           "completed",
		ParticipantCount: 12,
		ConferenceAlias:  "board-meeting",
		ResourceURI:      "/api/admin/status/v1/mjx_meeting/meeting-456/",
	}

	client.On("GetJSON", t.Context(), "status/v1/mjx_meeting/meeting-456/", mock.AnythingOfType("*status.MJXMeeting")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MJXMeeting)
		*result = *expectedMeeting
	})

	service := New(client)
	result, err := service.GetMJXMeeting(t.Context(), "meeting-456")

	assert.NoError(t, err)
	assert.Equal(t, expectedMeeting, result)
	assert.Equal(t, "Board Meeting", result.Subject)
	assert.Equal(t, "completed", result.Status)
	client.AssertExpectations(t)
}
