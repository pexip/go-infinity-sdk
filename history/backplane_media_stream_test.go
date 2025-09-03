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

func TestService_ListBackplaneMediaStreams(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	expectedResponse := &BackplaneMediaStreamListResponse{
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
		Objects: []BackplaneMediaStream{
			{
				ID:                1,
				StreamID:          "stream-123",
				StreamType:        "video",
				StartTime:         &util.InfinityTime{},
				EndTime:           &util.InfinityTime{},
				Node:              "node1.example.com",
				RxBitrate:         1024,
				RxCodec:           "H.264",
				RxFPS:             30.0,
				RxPacketLoss:      0.1,
				RxPacketsLost:     5,
				RxPacketsReceived: 5000,
				RxResolution:      "1920x1080",
				TxBitrate:         1024,
				TxCodec:           "H.264",
				TxFPS:             30.0,
				TxPacketLoss:      0.1,
				TxPacketsLost:     3,
				TxPacketsSent:     5000,
				TxResolution:      "1920x1080",
				ResourceURI:       "/api/admin/history/v1/backplane_media_stream/1/",
			},
		},
	}

	client.On("GetJSON", context.Background(), "history/v1/backplane_media_stream/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.BackplaneMediaStreamListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*BackplaneMediaStreamListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListBackplaneMediaStreams(context.Background(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "video", result.Objects[0].StreamType)
	assert.Equal(t, "H.264", result.Objects[0].RxCodec)
	assert.Equal(t, 30.0, result.Objects[0].RxFPS)

	client.AssertExpectations(t)
}

func TestService_GetBackplaneMediaStream(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	expectedStream := &BackplaneMediaStream{
		ID:                1,
		StreamID:          "stream-456",
		StreamType:        "audio",
		StartTime:         &util.InfinityTime{},
		EndTime:           &util.InfinityTime{},
		Node:              "node2.example.com",
		RxBitrate:         64,
		RxCodec:           "Opus",
		RxFPS:             0.0,
		RxPacketLoss:      0.05,
		RxPacketsLost:     2,
		RxPacketsReceived: 4000,
		TxBitrate:         64,
		TxCodec:           "Opus",
		TxFPS:             0.0,
		TxPacketLoss:      0.03,
		TxPacketsLost:     1,
		TxPacketsSent:     4000,
		ResourceURI:       "/api/admin/history/v1/backplane_media_stream/1/",
	}

	client.On("GetJSON", context.Background(), "history/v1/backplane_media_stream/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.BackplaneMediaStream")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*BackplaneMediaStream)
		*result = *expectedStream
	})

	result, err := service.GetBackplaneMediaStream(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedStream, result)

	client.AssertExpectations(t)
}

func TestService_ListBackplaneMediaStreamsByBackplane(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	expectedResponse := &BackplaneMediaStreamListResponse{
		Meta: struct {
			Limit      int    `json:"limit"`
			Next       string `json:"next"`
			Offset     int    `json:"offset"`
			Previous   string `json:"previous"`
			TotalCount int    `json:"total_count"`
		}{
			Limit:      20,
			TotalCount: 2,
		},
		Objects: []BackplaneMediaStream{
			{
				ID:         1,
				StreamID:   "stream-1",
				StreamType: "audio",
			},
			{
				ID:         2,
				StreamID:   "stream-2",
				StreamType: "video",
			},
		},
	}

	client.On("GetJSON", context.Background(), "history/v1/backplane_media_stream/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.BackplaneMediaStreamListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*BackplaneMediaStreamListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListBackplaneMediaStreamsByBackplane(context.Background(), "backplane-123", nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "audio", result.Objects[0].StreamType)
	assert.Equal(t, "video", result.Objects[1].StreamType)

	client.AssertExpectations(t)
}

func TestService_ListBackplaneMediaStreams_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	startTime := time.Now().Add(-2 * time.Hour)
	endTime := time.Now()
	opts := &ListOptions{
		SearchableListOptions: options.SearchableListOptions{
			BaseListOptions: options.BaseListOptions{
				Limit:  25,
				Offset: 0,
			},
		},
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	expectedResponse := &BackplaneMediaStreamListResponse{
		Meta: struct {
			Limit      int    `json:"limit"`
			Next       string `json:"next"`
			Offset     int    `json:"offset"`
			Previous   string `json:"previous"`
			TotalCount int    `json:"total_count"`
		}{
			Limit:      25,
			TotalCount: 0,
		},
		Objects: []BackplaneMediaStream{},
	}

	client.On("GetJSON", context.Background(), "history/v1/backplane_media_stream/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.BackplaneMediaStreamListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*BackplaneMediaStreamListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListBackplaneMediaStreams(context.Background(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 25, result.Meta.Limit)

	client.AssertExpectations(t)
}

func TestService_ListBackplaneMediaStreamsByBackplane_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	startTime := time.Now().Add(-1 * time.Hour)
	endTime := time.Now()
	opts := &ListOptions{
		SearchableListOptions: options.SearchableListOptions{
			BaseListOptions: options.BaseListOptions{
				Limit: 5,
			},
		},
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	expectedResponse := &BackplaneMediaStreamListResponse{
		Meta: struct {
			Limit      int    `json:"limit"`
			Next       string `json:"next"`
			Offset     int    `json:"offset"`
			Previous   string `json:"previous"`
			TotalCount int    `json:"total_count"`
		}{
			Limit:      5,
			TotalCount: 1,
		},
		Objects: []BackplaneMediaStream{
			{
				ID:         1,
				StreamID:   "stream-test",
				StreamType: "audio",
			},
		},
	}

	client.On("GetJSON", context.Background(), "history/v1/backplane_media_stream/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.BackplaneMediaStreamListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*BackplaneMediaStreamListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListBackplaneMediaStreamsByBackplane(context.Background(), "backplane-456", opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "audio", result.Objects[0].StreamType)

	client.AssertExpectations(t)
}

func TestService_ListBackplaneMediaStreams_Error(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	client.On("GetJSON", context.Background(), "history/v1/backplane_media_stream/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.BackplaneMediaStreamListResponse")).Return(errors.New("server error"))

	_, err := service.ListBackplaneMediaStreams(context.Background(), nil)
	assert.Error(t, err)
	assert.Equal(t, "server error", err.Error())

	client.AssertExpectations(t)
}

func TestService_ListBackplaneMediaStreamsByBackplane_Error(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	client.On("GetJSON", context.Background(), "history/v1/backplane_media_stream/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.BackplaneMediaStreamListResponse")).Return(errors.New("server error"))

	_, err := service.ListBackplaneMediaStreamsByBackplane(context.Background(), "backplane-error", nil)
	assert.Error(t, err)
	assert.Equal(t, "server error", err.Error())

	client.AssertExpectations(t)
}

func TestService_GetBackplaneMediaStream_NotFound(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	client.On("GetJSON", context.Background(), "history/v1/backplane_media_stream/999/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.BackplaneMediaStream")).Return(errors.New("backplane media stream not found"))

	_, err := service.GetBackplaneMediaStream(context.Background(), 999)
	assert.Error(t, err)

	client.AssertExpectations(t)
}
