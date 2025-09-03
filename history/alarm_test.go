/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package history

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListAlarms(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	expectedResponse := &AlarmListResponse{
		Meta: struct {
			Limit      int    `json:"limit"`
			Next       string `json:"next"`
			Offset     int    `json:"offset"`
			Previous   string `json:"previous"`
			TotalCount int    `json:"total_count"`
		}{
			Limit:      20,
			TotalCount: 1,
		},
		Objects: []Alarm{
			{
				ID:          1,
				Details:     "Test alarm details",
				Identifier:  100,
				Instance:    "test-instance",
				Level:       "warning",
				Name:        "capacity_exhausted",
				Node:        "192.168.1.1",
				TimeRaised:  &util.InfinityTime{},
				ResourceURI: "/api/admin/history/v1/alarm/1/",
			},
		},
	}

	client.On("GetJSON", context.Background(), "history/v1/alarm/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.AlarmListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*AlarmListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListAlarms(context.Background(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "warning", result.Objects[0].Level)
	assert.Equal(t, "capacity_exhausted", result.Objects[0].Name)

	client.AssertExpectations(t)
}

func TestService_GetAlarm(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	expectedAlarm := &Alarm{
		ID:          1,
		Details:     "Test alarm details",
		Identifier:  100,
		Instance:    "test-instance",
		Level:       "critical",
		Name:        "licenses_exhausted",
		Node:        "192.168.1.1",
		TimeRaised:  &util.InfinityTime{},
		ResourceURI: "/api/admin/history/v1/alarm/1/",
	}

	client.On("GetJSON", context.Background(), "history/v1/alarm/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.Alarm")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Alarm)
		*result = *expectedAlarm
	})

	result, err := service.GetAlarm(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedAlarm, result)

	client.AssertExpectations(t)
}

func TestService_ListAlarms_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	startTime := time.Now().Add(-24 * time.Hour)
	endTime := time.Now()
	opts := &ListOptions{
		SearchableListOptions: options.SearchableListOptions{
			BaseListOptions: options.BaseListOptions{
				Limit:  10,
				Offset: 5,
			},
			Search: "critical",
		},
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	expectedResponse := &AlarmListResponse{
		Meta: struct {
			Limit      int    `json:"limit"`
			Next       string `json:"next"`
			Offset     int    `json:"offset"`
			Previous   string `json:"previous"`
			TotalCount int    `json:"total_count"`
		}{
			Limit:      10,
			Offset:     5,
			TotalCount: 1,
		},
		Objects: []Alarm{},
	}

	client.On("GetJSON", context.Background(), "history/v1/alarm/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.AlarmListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*AlarmListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListAlarms(context.Background(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 10, result.Meta.Limit)
	assert.Equal(t, 5, result.Meta.Offset)

	client.AssertExpectations(t)
}

func TestService_ListAlarms_Error(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	client.On("GetJSON", context.Background(), "history/v1/alarm/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.AlarmListResponse")).Return(errors.New("server error"))

	_, err := service.ListAlarms(context.Background(), nil)
	assert.Error(t, err)
	assert.Equal(t, "server error", err.Error())

	client.AssertExpectations(t)
}

func TestService_GetAlarm_NotFound(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	client.On("GetJSON", context.Background(), "history/v1/alarm/999/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.Alarm")).Return(errors.New("alarm not found"))

	_, err := service.GetAlarm(context.Background(), 999)
	assert.Error(t, err)

	client.AssertExpectations(t)
}
