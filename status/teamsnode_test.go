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

func TestService_ListTeamsNodes(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	lastContact := time.Now().Add(-10 * time.Minute)

	expectedResponse := &TeamsNodeListResponse{
		Objects: []TeamsNode{
			{
				ID:            "1",
				Name:          "teams-node-1",
				MediaLoad:     0,
				ResourceURI:   "/api/admin/status/v1/teamsnode/1/",
				CallCount:     5,
				HeartbeatTime: &util.InfinityTime{Time: lastContact},
			},
			{
				ID:            "2",
				Name:          "teams-node-2",
				MediaLoad:     0,
				ResourceURI:   "/api/admin/status/v1/teamsnode/2/",
				CallCount:     0,
				HeartbeatTime: &util.InfinityTime{Time: lastContact},
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/teamsnode/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.TeamsNodeListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TeamsNodeListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListTeamsNodes(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "teams-node-1", result.Objects[0].Name)
	assert.Equal(t, "1", result.Objects[0].ID)
	assert.Equal(t, 5, result.Objects[0].CallCount)
	assert.Equal(t, "/api/admin/status/v1/teamsnode/1/", result.Objects[0].ResourceURI)
	assert.NotNil(t, result.Objects[0].HeartbeatTime)
	assert.Equal(t, lastContact.Unix(), result.Objects[0].HeartbeatTime.Time.Unix())
	assert.Equal(t, "teams-node-2", result.Objects[1].Name)
	assert.Equal(t, "2", result.Objects[1].ID)
	assert.Equal(t, 0, result.Objects[1].CallCount)
	assert.Equal(t, "/api/admin/status/v1/teamsnode/2/", result.Objects[1].ResourceURI)
	assert.NotNil(t, result.Objects[1].HeartbeatTime)
	assert.Equal(t, lastContact.Unix(), result.Objects[1].HeartbeatTime.Time.Unix())
	client.AssertExpectations(t)
}

func TestService_GetTeamsNode(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	lastContact := time.Now().Add(-5 * time.Minute)

	expectedNode := &TeamsNode{
		ID:            "1",
		Name:          "teams-node-primary",
		MediaLoad:     10,
		ResourceURI:   "/api/admin/status/v1/teamsnode/1/",
		CallCount:     12,
		HeartbeatTime: &util.InfinityTime{Time: lastContact},
	}

	client.On("GetJSON", t.Context(), "status/v1/teamsnode/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.TeamsNode")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TeamsNode)
		*result = *expectedNode
	})

	service := New(client)
	result, err := service.GetTeamsNode(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedNode, result)
	assert.Equal(t, "teams-node-primary", result.Name)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, 12, result.CallCount)
	assert.Equal(t, 10, result.MediaLoad)
	assert.Equal(t, "/api/admin/status/v1/teamsnode/1/", result.ResourceURI)
	assert.NotNil(t, result.HeartbeatTime)
	assert.Equal(t, lastContact.Unix(), result.HeartbeatTime.Time.Unix())
	client.AssertExpectations(t)
}

func TestService_ListTeamsNodes_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  3,
		Offset: 0,
	}

	expectedResponse := &TeamsNodeListResponse{
		Objects: []TeamsNode{
			{
				ID:          "3",
				Name:        "teams-node-test",
				MediaLoad:   2,
				CallCount:   8,
				ResourceURI: "/api/admin/status/v1/teamsnode/3/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/teamsnode/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.TeamsNodeListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TeamsNodeListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListTeamsNodes(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "teams-node-test", result.Objects[0].Name)
	assert.Equal(t, "3", result.Objects[0].ID)
	assert.Equal(t, 8, result.Objects[0].CallCount)
	assert.Equal(t, 2, result.Objects[0].MediaLoad)
	assert.Equal(t, "/api/admin/status/v1/teamsnode/3/", result.Objects[0].ResourceURI)

	client.AssertExpectations(t)
}
