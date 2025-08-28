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

func TestService_ListWorkerVMStatusEvents(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	state := "running"
	contextValue := "startup"
	details := "Worker VM started successfully"
	configID := 100
	configName := "worker-vm-1"
	locationName := "datacenter-1"

	expectedResponse := &WorkerVMStatusEventListResponse{
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
		Objects: []WorkerVMStatusEvent{
			{
				ID:                                1,
				EventType:                         "workervm_added",
				State:                             &state,
				Context:                           &contextValue,
				Details:                           &details,
				TimeChanged:                       &util.InfinityTime{},
				WorkerVMAddress:                   "192.168.1.10",
				WorkerVMConfigurationID:           &configID,
				WorkerVMConfigurationName:         &configName,
				WorkerVMConfigurationLocationName: &locationName,
				ResourceURI:                       "/api/admin/history/v1/workervm_status_event/1/",
			},
		},
	}

	client.On("GetJSON", context.Background(), "history/v1/workervm_status_event/", mock.AnythingOfType("*history.WorkerVMStatusEventListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*WorkerVMStatusEventListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListWorkerVMStatusEvents(context.Background(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "workervm_added", result.Objects[0].EventType)
	assert.Equal(t, "running", *result.Objects[0].State)
	assert.Equal(t, "192.168.1.10", result.Objects[0].WorkerVMAddress)

	client.AssertExpectations(t)
}

func TestService_GetWorkerVMStatusEvent(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	state := "stopped"
	contextValue := "maintenance"
	details := "Worker VM stopped for maintenance"
	configID := 200
	configName := "worker-vm-2"
	locationName := "datacenter-2"

	expectedEvent := &WorkerVMStatusEvent{
		ID:                                1,
		EventType:                         "workervm_removed",
		State:                             &state,
		Context:                           &contextValue,
		Details:                           &details,
		TimeChanged:                       &util.InfinityTime{},
		WorkerVMAddress:                   "192.168.1.20",
		WorkerVMConfigurationID:           &configID,
		WorkerVMConfigurationName:         &configName,
		WorkerVMConfigurationLocationName: &locationName,
		ResourceURI:                       "/api/admin/history/v1/workervm_status_event/1/",
	}

	client.On("GetJSON", context.Background(), "history/v1/workervm_status_event/1/", mock.AnythingOfType("*history.WorkerVMStatusEvent")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*WorkerVMStatusEvent)
		*result = *expectedEvent
	})

	result, err := service.GetWorkerVMStatusEvent(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvent, result)

	client.AssertExpectations(t)
}

func TestService_ListWorkerVMStatusEvents_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	startTime := time.Now().Add(-12 * time.Hour)
	endTime := time.Now()
	opts := &ListOptions{
		SearchableListOptions: options.SearchableListOptions{
			BaseListOptions: options.BaseListOptions{
				Limit:  100,
				Offset: 50,
			},
			Search: "worker",
		},
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	expectedResponse := &WorkerVMStatusEventListResponse{
		Meta: struct {
			Limit      int    `json:"limit"`
			Next       string `json:"next"`
			Offset     int    `json:"offset"`
			Previous   string `json:"previous"`
			TotalCount int    `json:"total_count"`
		}{
			Limit:      100,
			Offset:     50,
			TotalCount: 150,
		},
		Objects: []WorkerVMStatusEvent{
			{
				ID:              1,
				EventType:       "workervm_added",
				WorkerVMAddress: "192.168.1.10",
			},
		},
	}

	client.On("GetJSON", context.Background(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "history/v1/workervm_status_event/"
	}), mock.AnythingOfType("*history.WorkerVMStatusEventListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*WorkerVMStatusEventListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListWorkerVMStatusEvents(context.Background(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 100, result.Meta.Limit)
	assert.Equal(t, 50, result.Meta.Offset)
	assert.Equal(t, 150, result.Meta.TotalCount)

	client.AssertExpectations(t)
}

func TestService_ListWorkerVMStatusEvents_Error(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	client.On("GetJSON", context.Background(), "history/v1/workervm_status_event/", mock.AnythingOfType("*history.WorkerVMStatusEventListResponse")).Return(errors.New("server error"))

	_, err := service.ListWorkerVMStatusEvents(context.Background(), nil)
	assert.Error(t, err)
	assert.Equal(t, "server error", err.Error())

	client.AssertExpectations(t)
}

func TestService_GetWorkerVMStatusEvent_NotFound(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	client.On("GetJSON", context.Background(), "history/v1/workervm_status_event/999/", mock.AnythingOfType("*history.WorkerVMStatusEvent")).Return(errors.New("worker VM status event not found"))

	_, err := service.GetWorkerVMStatusEvent(context.Background(), 999)
	assert.Error(t, err)

	client.AssertExpectations(t)
}
