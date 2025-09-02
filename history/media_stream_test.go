/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package history

import (
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListMediaStreams(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedResponse := &MediaStreamListResponse{
		Objects: []MediaStream{
			{
				ID:              1,
				ParticipantID:   1,
				StreamType:      "video",
				Direction:       "sendrecv",
				Codec:           "H.264",
				Resolution:      "1920x1080",
				Bitrate:         2000000,
				DurationSeconds: 1800,
			},
		},
	}

	client.On("GetJSON", t.Context(), "history/v1/media_stream/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.MediaStreamListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaStreamListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListMediaStreams(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "video", result.Objects[0].StreamType)
	assert.Equal(t, "H.264", result.Objects[0].Codec)
	client.AssertExpectations(t)
}

func TestService_GetMediaStream(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedStream := &MediaStream{
		ID:              1,
		ParticipantID:   1,
		StreamType:      "audio",
		Direction:       "sendrecv",
		Codec:           "Opus",
		Bitrate:         64000,
		DurationSeconds: 1800,
	}

	client.On("GetJSON", t.Context(), "history/v1/media_stream/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.MediaStream")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaStream)
		*result = *expectedStream
	})

	service := New(client)
	result, err := service.GetMediaStream(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedStream, result)
	client.AssertExpectations(t)
}

func TestService_ListMediaStreamsByParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedResponse := &MediaStreamListResponse{
		Objects: []MediaStream{
			{
				ID:            1,
				ParticipantID: 456,
				StreamType:    "video",
				Direction:     "sendrecv",
				Codec:         "H.264",
			},
			{
				ID:            2,
				ParticipantID: 456,
				StreamType:    "audio",
				Direction:     "sendrecv",
				Codec:         "Opus",
			},
		},
	}

	client.On("GetJSON", t.Context(), "history/v1/media_stream/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.MediaStreamListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaStreamListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListMediaStreamsByParticipant(t.Context(), 456, nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "video", result.Objects[0].StreamType)
	assert.Equal(t, "audio", result.Objects[1].StreamType)
	client.AssertExpectations(t)
}

func TestService_ListMediaStreams_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	startTime := time.Now().Add(-30 * time.Minute)
	endTime := time.Now()
	opts := &ListOptions{
		SearchableListOptions: options.SearchableListOptions{
			BaseListOptions: options.BaseListOptions{
				Limit:  20,
				Offset: 0,
			},
		},
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	expectedResponse := &MediaStreamListResponse{
		Objects: []MediaStream{
			{
				ID:            1,
				ParticipantID: 1,
				StreamType:    "video",
				Direction:     "sendrecv",
				Codec:         "H.264",
			},
		},
	}

	client.On("GetJSON", t.Context(), "history/v1/media_stream/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.MediaStreamListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaStreamListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListMediaStreams(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "video", result.Objects[0].StreamType)

	client.AssertExpectations(t)
}

func TestService_ListMediaStreamsByParticipant_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	startTime := time.Now().Add(-45 * time.Minute)
	endTime := time.Now()
	opts := &ListOptions{
		SearchableListOptions: options.SearchableListOptions{
			BaseListOptions: options.BaseListOptions{
				Limit: 10,
			},
		},
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	expectedResponse := &MediaStreamListResponse{
		Objects: []MediaStream{
			{
				ID:            1,
				ParticipantID: 456,
				StreamType:    "video",
				Direction:     "sendrecv",
				Codec:         "H.264",
			},
			{
				ID:            2,
				ParticipantID: 456,
				StreamType:    "audio",
				Direction:     "sendrecv",
				Codec:         "Opus",
			},
		},
	}

	client.On("GetJSON", t.Context(), "history/v1/media_stream/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.MediaStreamListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaStreamListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListMediaStreamsByParticipant(t.Context(), 456, opts)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "video", result.Objects[0].StreamType)
	assert.Equal(t, "audio", result.Objects[1].StreamType)

	client.AssertExpectations(t)
}
