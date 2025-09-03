/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListTeamsNodeCalls(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	startTime := time.Now().Add(-30 * time.Minute)

	expectedResponse := &TeamsNodeCallListResponse{
		Objects: []TeamsNodeCall{
			{
				ID:          "call-123",
				TeamsNodeID: "1",
				StartTime:   &util.InfinityTime{Time: startTime},
				ResourceURI: "/api/admin/status/v1/teamsnode_call/call-123/",
				// Only fields present in the struct are set
				// Destination, Source, State, HeartbeatTime are omitted for brevity
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/teamsnode_call/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.TeamsNodeCallListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TeamsNodeCallListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListTeamsNodeCalls(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "call-123", result.Objects[0].ID)
	assert.Equal(t, "1", result.Objects[0].TeamsNodeID)
	assert.Equal(t, "/api/admin/status/v1/teamsnode_call/call-123/", result.Objects[0].ResourceURI)
	assert.Equal(t, startTime.Unix(), result.Objects[0].StartTime.Time.Unix())
	client.AssertExpectations(t)
}

func TestService_ListTeamsNodeCalls_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  20,
		Offset: 10,
	}

	startTime := time.Now().Add(-15 * time.Minute)
	expectedResponse := &TeamsNodeCallListResponse{
		Objects: []TeamsNodeCall{
			{
				ID:          "call-options-test",
				TeamsNodeID: "3",
				StartTime:   &util.InfinityTime{Time: startTime},
				// Only fields present in the struct are set
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/teamsnode_call/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.TeamsNodeCallListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TeamsNodeCallListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListTeamsNodeCalls(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "call-options-test", result.Objects[0].ID)
	assert.Equal(t, "3", result.Objects[0].TeamsNodeID)
	assert.Equal(t, startTime.Unix(), result.Objects[0].StartTime.Time.Unix())

	client.AssertExpectations(t)
}

func TestService_GetTeamsNodeCall(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	startTime := time.Now().Add(-1 * time.Hour)
	expectedCall := &TeamsNodeCall{
		ID:          "call-456",
		TeamsNodeID: "2",
		StartTime:   &util.InfinityTime{Time: startTime},
		ResourceURI: "/api/admin/status/v1/teamsnode_call/call-456/",
		// Only fields present in the struct are set
	}

	client.On("GetJSON", t.Context(), "status/v1/teamsnode_call/call-456/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.TeamsNodeCall")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TeamsNodeCall)
		*result = *expectedCall
	})

	service := New(client)
	result, err := service.GetTeamsNodeCall(t.Context(), "call-456")

	assert.NoError(t, err)
	assert.Equal(t, expectedCall, result)
	assert.Equal(t, "call-456", result.ID)
	assert.Equal(t, "2", result.TeamsNodeID)
	assert.Equal(t, "/api/admin/status/v1/teamsnode_call/call-456/", result.ResourceURI)
	assert.Equal(t, startTime.Unix(), result.StartTime.Time.Unix())
	client.AssertExpectations(t)
}
