/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListScheduledConferences(t *testing.T) {
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
				startTime1 := util.InfinityTime{Time: time.Date(2023, 10, 15, 10, 0, 0, 0, time.UTC)}
				endTime1 := util.InfinityTime{Time: time.Date(2023, 10, 15, 11, 0, 0, 0, time.UTC)}
				startTime2 := util.InfinityTime{Time: time.Date(2023, 10, 16, 14, 0, 0, 0, time.UTC)}
				endTime2 := util.InfinityTime{Time: time.Date(2023, 10, 16, 15, 30, 0, 0, time.UTC)}
				recurringConf1 := "/api/admin/configuration/v1/recurring_conference/1/"
				scheduledAlias1 := "/api/admin/configuration/v1/scheduled_alias/10/"
				scheduledAlias2 := "/api/admin/configuration/v1/scheduled_alias/11/"

				expectedResponse := &ScheduledConferenceListResponse{
					Objects: []ScheduledConference{
						{ID: 1, Conference: "/api/admin/configuration/v1/conference/1/", StartTime: startTime1, EndTime: endTime1, Subject: "Weekly Team Meeting", EWSItemID: "ews-id-1", EWSItemUID: "ews-uid-1", RecurringConference: &recurringConf1, ScheduledAlias: &scheduledAlias1},
						{ID: 2, Conference: "/api/admin/configuration/v1/conference/2/", StartTime: startTime2, EndTime: endTime2, Subject: "Project Review", EWSItemID: "ews-id-2", EWSItemUID: "ews-uid-2", ScheduledAlias: &scheduledAlias2},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/scheduled_conference/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ScheduledConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ScheduledConferenceListResponse)
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
				Search: "team",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				startTime := util.InfinityTime{Time: time.Date(2023, 10, 15, 10, 0, 0, 0, time.UTC)}
				endTime := util.InfinityTime{Time: time.Date(2023, 10, 15, 11, 0, 0, 0, time.UTC)}
				recurringConf := "/api/admin/configuration/v1/recurring_conference/1/"
				scheduledAlias := "/api/admin/configuration/v1/scheduled_alias/10/"

				expectedResponse := &ScheduledConferenceListResponse{
					Objects: []ScheduledConference{
						{ID: 1, Conference: "/api/admin/configuration/v1/conference/1/", StartTime: startTime, EndTime: endTime, Subject: "Weekly Team Meeting", EWSItemID: "ews-id-1", EWSItemUID: "ews-uid-1", RecurringConference: &recurringConf, ScheduledAlias: &scheduledAlias},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/scheduled_conference/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ScheduledConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ScheduledConferenceListResponse)
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
			result, err := service.ListScheduledConferences(t.Context(), tt.opts)

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

func TestService_GetScheduledConference(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	startTime := util.InfinityTime{Time: time.Date(2023, 12, 1, 9, 0, 0, 0, time.UTC)}
	endTime := util.InfinityTime{Time: time.Date(2023, 12, 1, 10, 0, 0, 0, time.UTC)}
	recurringConference := "/api/admin/configuration/v1/recurring_conference/5/"
	scheduledAlias := "/api/admin/configuration/v1/scheduled_alias/25/"

	expectedScheduledConference := &ScheduledConference{
		ID:                  1,
		Conference:          "/api/admin/configuration/v1/conference/5/",
		StartTime:           startTime,
		EndTime:             endTime,
		Subject:             "Test Scheduled Conference",
		EWSItemID:           "test-ews-id",
		EWSItemUID:          "test-ews-uid",
		RecurringConference: &recurringConference,
		ScheduledAlias:      &scheduledAlias,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/scheduled_conference/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ScheduledConference")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ScheduledConference)
		*result = *expectedScheduledConference
	})

	service := New(client)
	result, err := service.GetScheduledConference(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedScheduledConference, result)
	client.AssertExpectations(t)
}

func TestService_CreateScheduledConference(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	startTime := util.InfinityTime{Time: time.Date(2024, 1, 15, 14, 0, 0, 0, time.UTC)}
	endTime := util.InfinityTime{Time: time.Date(2024, 1, 15, 15, 30, 0, 0, time.UTC)}
	recurringConference := "/api/admin/configuration/v1/recurring_conference/3/"
	scheduledAlias := "/api/admin/configuration/v1/scheduled_alias/30/"

	createRequest := &ScheduledConferenceCreateRequest{
		Conference:          "/api/admin/configuration/v1/conference/10/",
		StartTime:           startTime,
		EndTime:             endTime,
		Subject:             "New Scheduled Conference",
		EWSItemID:           "new-ews-id",
		EWSItemUID:          "new-ews-uid",
		RecurringConference: &recurringConference,
		ScheduledAlias:      &scheduledAlias,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/scheduled_conference/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/scheduled_conference/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateScheduledConference(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateScheduledConference(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	newStartTime := util.InfinityTime{Time: time.Date(2024, 2, 1, 10, 0, 0, 0, time.UTC)}
	newEndTime := util.InfinityTime{Time: time.Date(2024, 2, 1, 12, 0, 0, 0, time.UTC)}
	scheduledAlias := "/api/admin/configuration/v1/scheduled_alias/35/"

	updateRequest := &ScheduledConferenceUpdateRequest{
		Conference:     "/api/admin/configuration/v1/conference/15/",
		StartTime:      &newStartTime,
		EndTime:        &newEndTime,
		Subject:        "Updated Scheduled Conference",
		ScheduledAlias: &scheduledAlias,
	}

	recurringConference := "/api/admin/configuration/v1/recurring_conference/5/"
	expectedScheduledConference := &ScheduledConference{
		ID:                  1,
		Conference:          "/api/admin/configuration/v1/conference/15/",
		StartTime:           newStartTime,
		EndTime:             newEndTime,
		Subject:             "Updated Scheduled Conference",
		EWSItemID:           "test-ews-id",
		EWSItemUID:          "test-ews-uid",
		RecurringConference: &recurringConference,
		ScheduledAlias:      &scheduledAlias,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/scheduled_conference/1/", updateRequest, mock.AnythingOfType("*config.ScheduledConference")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ScheduledConference)
		*result = *expectedScheduledConference
	})

	service := New(client)
	result, err := service.UpdateScheduledConference(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedScheduledConference, result)
	client.AssertExpectations(t)
}

func TestService_DeleteScheduledConference(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/scheduled_conference/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteScheduledConference(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
