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

func TestService_ListScheduledScalings(t *testing.T) {
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
				updated1 := util.InfinityTime{Time: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)}
				updated2 := util.InfinityTime{Time: time.Date(2023, 10, 2, 15, 30, 0, 0, time.UTC)}

				expectedResponse := &ScheduledScalingListResponse{
					Objects: []ScheduledScaling{
						{ID: 1, PolicyName: "weekend-scaling", PolicyType: "weekly", ResourceIdentifier: "conferencing-nodes", Enabled: true, InstancesToAdd: 5, MinutesInAdvance: 30, LocalTimezone: "UTC", StartDate: "2023-10-01", TimeFrom: "09:00", TimeTo: "17:00", Mon: true, Tue: true, Wed: true, Thu: true, Fri: true, Sat: false, Sun: false, Updated: &updated1},
						{ID: 2, PolicyName: "peak-hours", PolicyType: "daily", ResourceIdentifier: "media-nodes", Enabled: true, InstancesToAdd: 3, MinutesInAdvance: 15, LocalTimezone: "America/New_York", StartDate: "2023-10-02", TimeFrom: "08:00", TimeTo: "18:00", Mon: true, Tue: true, Wed: true, Thu: true, Fri: true, Sat: true, Sun: true, Updated: &updated2},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/scheduled_scaling/", mock.AnythingOfType("*config.ScheduledScalingListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ScheduledScalingListResponse)
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
				Search: "weekend",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				updated := util.InfinityTime{Time: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)}

				expectedResponse := &ScheduledScalingListResponse{
					Objects: []ScheduledScaling{
						{ID: 1, PolicyName: "weekend-scaling", PolicyType: "weekly", ResourceIdentifier: "conferencing-nodes", Enabled: true, InstancesToAdd: 5, MinutesInAdvance: 30, LocalTimezone: "UTC", StartDate: "2023-10-01", TimeFrom: "09:00", TimeTo: "17:00", Mon: true, Tue: true, Wed: true, Thu: true, Fri: true, Sat: false, Sun: false, Updated: &updated},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/scheduled_scaling/?limit=5&name__icontains=weekend", mock.AnythingOfType("*config.ScheduledScalingListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ScheduledScalingListResponse)
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
			result, err := service.ListScheduledScalings(t.Context(), tt.opts)

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

func TestService_GetScheduledScaling(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	updated := util.InfinityTime{Time: time.Date(2023, 11, 15, 10, 30, 0, 0, time.UTC)}

	expectedScheduledScaling := &ScheduledScaling{
		ID:                 1,
		PolicyName:         "test-scaling-policy",
		PolicyType:         "daily",
		ResourceIdentifier: "test-nodes",
		Enabled:            true,
		InstancesToAdd:     10,
		MinutesInAdvance:   45,
		LocalTimezone:      "Europe/London",
		StartDate:          "2023-11-01",
		TimeFrom:           "08:30",
		TimeTo:             "17:30",
		Mon:                true,
		Tue:                true,
		Wed:                true,
		Thu:                true,
		Fri:                true,
		Sat:                false,
		Sun:                false,
		Updated:            &updated,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/scheduled_scaling/1/", mock.AnythingOfType("*config.ScheduledScaling")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ScheduledScaling)
		*result = *expectedScheduledScaling
	})

	service := New(client)
	result, err := service.GetScheduledScaling(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedScheduledScaling, result)
	client.AssertExpectations(t)
}

func TestService_CreateScheduledScaling(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &ScheduledScalingCreateRequest{
		PolicyName:         "new-scaling-policy",
		PolicyType:         "weekly",
		ResourceIdentifier: "new-nodes",
		Enabled:            true,
		InstancesToAdd:     7,
		MinutesInAdvance:   20,
		LocalTimezone:      "Asia/Tokyo",
		StartDate:          "2024-01-01",
		TimeFrom:           "07:00",
		TimeTo:             "19:00",
		Mon:                true,
		Tue:                true,
		Wed:                true,
		Thu:                true,
		Fri:                true,
		Sat:                true,
		Sun:                false,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/scheduled_scaling/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/scheduled_scaling/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateScheduledScaling(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateScheduledScaling(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	enabled := false
	instancesToAdd := 15
	minutesInAdvance := 60
	sat := true
	sun := true

	updateRequest := &ScheduledScalingUpdateRequest{
		PolicyName:       "updated-scaling-policy",
		Enabled:          &enabled,
		InstancesToAdd:   &instancesToAdd,
		MinutesInAdvance: &minutesInAdvance,
		TimeFrom:         "06:00",
		TimeTo:           "20:00",
		Sat:              &sat,
		Sun:              &sun,
	}

	updated := util.InfinityTime{Time: time.Date(2023, 11, 20, 14, 15, 0, 0, time.UTC)}
	expectedScheduledScaling := &ScheduledScaling{
		ID:                 1,
		PolicyName:         "updated-scaling-policy",
		PolicyType:         "daily",
		ResourceIdentifier: "test-nodes",
		Enabled:            false,
		InstancesToAdd:     15,
		MinutesInAdvance:   60,
		LocalTimezone:      "Europe/London",
		StartDate:          "2023-11-01",
		TimeFrom:           "06:00",
		TimeTo:             "20:00",
		Mon:                true,
		Tue:                true,
		Wed:                true,
		Thu:                true,
		Fri:                true,
		Sat:                true,
		Sun:                true,
		Updated:            &updated,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/scheduled_scaling/1/", updateRequest, mock.AnythingOfType("*config.ScheduledScaling")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ScheduledScaling)
		*result = *expectedScheduledScaling
	})

	service := New(client)
	result, err := service.UpdateScheduledScaling(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedScheduledScaling, result)
	client.AssertExpectations(t)
}

func TestService_DeleteScheduledScaling(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/scheduled_scaling/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteScheduledScaling(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
