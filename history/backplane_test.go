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

func TestService_ListBackplanes(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	expectedResponse := &BackplaneListResponse{
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
		Objects: []Backplane{
			{
				ID:                   "test-backplane-1",
				ConferenceName:       "test-conference",
				DisconnectReason:     "Normal call clearing",
				Duration:             &[]int{3600}[0],
				StartTime:            &util.InfinityTime{},
				EndTime:              &util.InfinityTime{},
				MediaNode:            "node1.example.com",
				Protocol:             "INTERNAL",
				RemoteConferenceName: "remote-conference",
				RemoteMediaNode:      "node2.example.com",
				Type:                 "local-backplane",
				ResourceURI:          "/api/admin/history/v1/backplane/test-backplane-1/",
			},
		},
	}

	client.On("GetJSON", context.Background(), "history/v1/backplane/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.BackplaneListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*BackplaneListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListBackplanes(context.Background(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "test-backplane-1", result.Objects[0].ID)
	assert.Equal(t, "test-conference", result.Objects[0].ConferenceName)
	assert.Equal(t, "INTERNAL", result.Objects[0].Protocol)
	assert.Equal(t, "local-backplane", result.Objects[0].Type)

	client.AssertExpectations(t)
}

func TestService_GetBackplane(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	expectedBackplane := &Backplane{
		ID:                   "test-backplane-1",
		ConferenceName:       "test-conference",
		DisconnectReason:     "Normal call clearing",
		Duration:             &[]int{7200}[0],
		StartTime:            &util.InfinityTime{},
		EndTime:              &util.InfinityTime{},
		MediaNode:            "node1.example.com",
		Protocol:             "GMS",
		RemoteConferenceName: "remote-conference",
		RemoteMediaNode:      "node2.example.com",
		Type:                 "geo-backplane",
		ResourceURI:          "/api/admin/history/v1/backplane/test-backplane-1/",
	}

	client.On("GetJSON", context.Background(), "history/v1/backplane/test-backplane-1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.Backplane")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Backplane)
		*result = *expectedBackplane
	})

	result, err := service.GetBackplane(context.Background(), "test-backplane-1")
	assert.NoError(t, err)
	assert.Equal(t, expectedBackplane, result)

	client.AssertExpectations(t)
}

func TestService_ListBackplanes_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	startTime := time.Now().Add(-24 * time.Hour)
	endTime := time.Now()
	opts := &ListOptions{
		SearchableListOptions: options.SearchableListOptions{
			BaseListOptions: options.BaseListOptions{
				Limit:  15,
				Offset: 10,
			},
			Search: "conference",
		},
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	expectedResponse := &BackplaneListResponse{
		Meta: struct {
			Limit      int    `json:"limit"`
			Next       string `json:"next"`
			Offset     int    `json:"offset"`
			Previous   string `json:"previous"`
			TotalCount int    `json:"total_count"`
		}{
			Limit:      15,
			Offset:     10,
			TotalCount: 0,
		},
		Objects: []Backplane{},
	}

	client.On("GetJSON", context.Background(), "history/v1/backplane/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.BackplaneListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*BackplaneListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListBackplanes(context.Background(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 15, result.Meta.Limit)
	assert.Equal(t, 10, result.Meta.Offset)

	client.AssertExpectations(t)
}

func TestService_ListBackplanes_Error(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	client.On("GetJSON", context.Background(), "history/v1/backplane/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.BackplaneListResponse")).Return(errors.New("server error"))

	_, err := service.ListBackplanes(context.Background(), nil)
	assert.Error(t, err)
	assert.Equal(t, "server error", err.Error())

	client.AssertExpectations(t)
}

func TestService_GetBackplane_NotFound(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	client.On("GetJSON", context.Background(), "history/v1/backplane/nonexistent/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.Backplane")).Return(errors.New("backplane not found"))

	_, err := service.GetBackplane(context.Background(), "nonexistent")
	assert.Error(t, err)

	client.AssertExpectations(t)
}
