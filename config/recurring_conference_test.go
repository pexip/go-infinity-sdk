package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListRecurringConferences(t *testing.T) {
	tests := []struct {
		name    string
		opts    *ListOptions
		setup   func(m *interfaces.HTTPClientMock)
		wantErr bool
	}{
		{
			name: "successful list without options",
			opts: nil,
			setup: func(m *interfaces.HTTPClientMock) {
				scheduledAlias1 := "weekly-meeting"
				scheduledAlias2 := "monthly-review"
				expectedResponse := &RecurringConferenceListResponse{
					Objects: []RecurringConference{
						{ID: 1, Conference: "weekly-standup", CurrentIndex: 5, EWSItemID: "ews-id-1", IsDepleted: false, Subject: "Weekly Standup", ScheduledAlias: &scheduledAlias1},
						{ID: 2, Conference: "monthly-review", CurrentIndex: 2, EWSItemID: "ews-id-2", IsDepleted: false, Subject: "Monthly Review", ScheduledAlias: &scheduledAlias2},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/recurring_conference/", mock.AnythingOfType("*config.RecurringConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*RecurringConferenceListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit: 5,
				},
				Search: "weekly",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				scheduledAlias := "weekly-meeting"
				expectedResponse := &RecurringConferenceListResponse{
					Objects: []RecurringConference{
						{ID: 1, Conference: "weekly-standup", CurrentIndex: 5, EWSItemID: "ews-id-1", IsDepleted: false, Subject: "Weekly Standup", ScheduledAlias: &scheduledAlias},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/recurring_conference/?limit=5&name__icontains=weekly", mock.AnythingOfType("*config.RecurringConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*RecurringConferenceListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := interfaces.NewHTTPClientMock()
			tt.setup(client)

			service := New(client)
			result, err := service.ListRecurringConferences(t.Context(), tt.opts)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			client.AssertExpectations(t)
		})
	}
}

func TestService_GetRecurringConference(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	scheduledAlias := "test-alias"
	expectedRecurringConference := &RecurringConference{
		ID:             1,
		Conference:     "test-conference",
		CurrentIndex:   10,
		EWSItemID:      "ews-test-id",
		IsDepleted:     false,
		Subject:        "Test Recurring Conference",
		ScheduledAlias: &scheduledAlias,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/recurring_conference/1/", mock.AnythingOfType("*config.RecurringConference")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*RecurringConference)
		*result = *expectedRecurringConference
	})

	service := New(client)
	result, err := service.GetRecurringConference(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRecurringConference, result)
	client.AssertExpectations(t)
}

func TestService_CreateRecurringConference(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	scheduledAlias := "new-alias"
	createRequest := &RecurringConferenceCreateRequest{
		Conference:     "new-conference",
		CurrentIndex:   0,
		EWSItemID:      "new-ews-id",
		IsDepleted:     false,
		Subject:        "New Recurring Conference",
		ScheduledAlias: &scheduledAlias,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/recurring_conference/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/recurring_conference/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateRecurringConference(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateRecurringConference(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	currentIndex := 15
	isDepleted := true
	scheduledAlias := "updated-alias"

	updateRequest := &RecurringConferenceUpdateRequest{
		Conference:     "updated-conference",
		CurrentIndex:   &currentIndex,
		IsDepleted:     &isDepleted,
		Subject:        "Updated Recurring Conference",
		ScheduledAlias: &scheduledAlias,
	}

	expectedRecurringConference := &RecurringConference{
		ID:             1,
		Conference:     "updated-conference",
		CurrentIndex:   15,
		EWSItemID:      "ews-test-id",
		IsDepleted:     true,
		Subject:        "Updated Recurring Conference",
		ScheduledAlias: &scheduledAlias,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/recurring_conference/1/", updateRequest, mock.AnythingOfType("*config.RecurringConference")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*RecurringConference)
		*result = *expectedRecurringConference
	})

	service := New(client)
	result, err := service.UpdateRecurringConference(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedRecurringConference, result)
	client.AssertExpectations(t)
}

func TestService_DeleteRecurringConference(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/recurring_conference/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteRecurringConference(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
