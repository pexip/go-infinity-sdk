package status

import (
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListMJXMeetings(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	startTime := time.Now().Add(-1 * time.Hour)
	endTime := time.Now().Add(1 * time.Hour)

	expectedResponse := &MJXMeetingListResponse{
		Objects: []MJXMeeting{
			{
				Alias:                        "team-weekly",
				EndTime:                      &util.InfinityTime{Time: endTime},
				EndpointName:                 "endpoint-1",
				ID:                           123,
				LastModifiedTime:             nil,
				MatchedMeetingProcessingRule: "",
				MeetingID:                    "meeting-123",
				MJXIntegrationID:             0,
				MJXIntegrationName:           "",
				OrganizerEmail:               "john.doe@company.com",
				OrganizerName:                "",
				ResourceURI:                  "/api/admin/status/v1/mjx_meeting/meeting-123/",
				RoomEmail:                    "",
				StartTime:                    &util.InfinityTime{Time: startTime},
				Subject:                      "Weekly Team Meeting",
				WorkerID:                     0,
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
	assert.Equal(t, "john.doe@company.com", result.Objects[0].OrganizerEmail)
	assert.Equal(t, "team-weekly", result.Objects[0].Alias)
	assert.Equal(t, "meeting-123", result.Objects[0].MeetingID)
	assert.Equal(t, "/api/admin/status/v1/mjx_meeting/meeting-123/", result.Objects[0].ResourceURI)
	assert.Equal(t, 123, result.Objects[0].ID)
	assert.Equal(t, "endpoint-1", result.Objects[0].EndpointName)
	assert.Equal(t, startTime.Unix(), result.Objects[0].StartTime.Time.Unix())
	assert.Equal(t, endTime.Unix(), result.Objects[0].EndTime.Time.Unix())
	client.AssertExpectations(t)
}

func TestService_ListMJXMeetings_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
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
				Alias:                        "test-options",
				EndTime:                      &util.InfinityTime{Time: endTime},
				EndpointName:                 "endpoint-test",
				ID:                           456,
				LastModifiedTime:             nil,
				MatchedMeetingProcessingRule: "",
				MeetingID:                    "meeting-options-test",
				MJXIntegrationID:             0,
				MJXIntegrationName:           "",
				OrganizerEmail:               "test@company.com",
				OrganizerName:                "",
				ResourceURI:                  "",
				RoomEmail:                    "",
				StartTime:                    &util.InfinityTime{Time: startTime},
				Subject:                      "Test Meeting With Options",
				WorkerID:                     0,
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
	assert.Equal(t, "test@company.com", result.Objects[0].OrganizerEmail)
	assert.Equal(t, "test-options", result.Objects[0].Alias)
	assert.Equal(t, "meeting-options-test", result.Objects[0].MeetingID)
	assert.Equal(t, "endpoint-test", result.Objects[0].EndpointName)
	assert.Equal(t, startTime.Unix(), result.Objects[0].StartTime.Time.Unix())
	assert.Equal(t, endTime.Unix(), result.Objects[0].EndTime.Time.Unix())

	client.AssertExpectations(t)
}

func TestService_GetMJXMeeting(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	startTime := time.Now().Add(-2 * time.Hour)
	endTime := time.Now().Add(-1 * time.Hour)
	expectedMeeting := &MJXMeeting{
		Alias:                        "board-meeting",
		EndTime:                      &util.InfinityTime{Time: endTime},
		EndpointName:                 "endpoint-board",
		ID:                           789,
		LastModifiedTime:             nil,
		MatchedMeetingProcessingRule: "",
		MeetingID:                    "meeting-456",
		MJXIntegrationID:             0,
		MJXIntegrationName:           "",
		OrganizerEmail:               "ceo@company.com",
		OrganizerName:                "",
		ResourceURI:                  "/api/admin/status/v1/mjx_meeting/meeting-456/",
		RoomEmail:                    "",
		StartTime:                    &util.InfinityTime{Time: startTime},
		Subject:                      "Board Meeting",
		WorkerID:                     0,
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
	assert.Equal(t, "ceo@company.com", result.OrganizerEmail)
	assert.Equal(t, "board-meeting", result.Alias)
	assert.Equal(t, "meeting-456", result.MeetingID)
	assert.Equal(t, "/api/admin/status/v1/mjx_meeting/meeting-456/", result.ResourceURI)
	assert.Equal(t, 789, result.ID)
	assert.Equal(t, "endpoint-board", result.EndpointName)
	assert.Equal(t, startTime.Unix(), result.StartTime.Time.Unix())
	assert.Equal(t, endTime.Unix(), result.EndTime.Time.Unix())
	client.AssertExpectations(t)
}
