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

func TestService_ListParticipants(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedResponse := &ParticipantListResponse{
		Objects: []Participant{
			{
				ID:               1,
				ConferenceID:     1,
				ConferenceName:   "Test Conference",
				DisplayName:      "John Doe",
				Role:             "chair",
				DurationSeconds:  1800,
				DisconnectReason: "normal",
			},
		},
	}

	client.On("GetJSON", t.Context(), "history/v1/participant/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.ParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ParticipantListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListParticipants(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "John Doe", result.Objects[0].DisplayName)
	client.AssertExpectations(t)
}

func TestService_GetParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedParticipant := &Participant{
		ID:               1,
		ConferenceID:     1,
		ConferenceName:   "Test Conference",
		DisplayName:      "John Doe",
		Role:             "chair",
		DurationSeconds:  1800,
		DisconnectReason: "normal",
	}

	client.On("GetJSON", t.Context(), "history/v1/participant/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.Participant")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Participant)
		*result = *expectedParticipant
	})

	service := New(client)
	result, err := service.GetParticipant(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedParticipant, result)
	client.AssertExpectations(t)
}

func TestService_ListParticipantsByConference(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedResponse := &ParticipantListResponse{
		Objects: []Participant{
			{
				ID:             1,
				ConferenceID:   123,
				ConferenceName: "Test Conference",
				DisplayName:    "Alice",
				Role:           "chair",
			},
			{
				ID:             2,
				ConferenceID:   123,
				ConferenceName: "Test Conference",
				DisplayName:    "Bob",
				Role:           "guest",
			},
		},
	}

	client.On("GetJSON", t.Context(), "history/v1/participant/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.ParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ParticipantListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListParticipantsByConference(t.Context(), 123, nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "Alice", result.Objects[0].DisplayName)
	assert.Equal(t, "Bob", result.Objects[1].DisplayName)
	client.AssertExpectations(t)
}

func TestService_ListParticipants_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	startTime := time.Now().Add(-3 * time.Hour)
	endTime := time.Now()
	opts := &ListOptions{
		SearchableListOptions: options.SearchableListOptions{
			BaseListOptions: options.BaseListOptions{
				Limit:  30,
				Offset: 10,
			},
			Search: "john",
		},
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	expectedResponse := &ParticipantListResponse{
		Objects: []Participant{
			{
				ID:             1,
				ConferenceID:   1,
				ConferenceName: "Test Conference",
				DisplayName:    "John Doe",
				Role:           "chair",
			},
		},
	}

	client.On("GetJSON", t.Context(), "history/v1/participant/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.ParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ParticipantListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListParticipants(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "John Doe", result.Objects[0].DisplayName)

	client.AssertExpectations(t)
}

func TestService_ListParticipantsByConference_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	startTime := time.Now().Add(-1 * time.Hour)
	endTime := time.Now()
	opts := &ListOptions{
		SearchableListOptions: options.SearchableListOptions{
			BaseListOptions: options.BaseListOptions{
				Limit: 5,
			},
			Search: "alice",
		},
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	expectedResponse := &ParticipantListResponse{
		Objects: []Participant{
			{
				ID:             1,
				ConferenceID:   123,
				ConferenceName: "Test Conference",
				DisplayName:    "Alice",
				Role:           "chair",
			},
		},
	}

	client.On("GetJSON", t.Context(), "history/v1/participant/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*history.ParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ParticipantListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListParticipantsByConference(t.Context(), 123, opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Alice", result.Objects[0].DisplayName)

	client.AssertExpectations(t)
}
