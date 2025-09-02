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

func TestService_ListMJXEndpoints(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	lastContact := time.Now().Add(-15 * time.Minute)

	expectedResponse := &MJXEndpointListResponse{
		Objects: []MJXEndpoint{
			{
				EndpointAddress:    "10.0.0.1",
				EndpointName:       "mjx-endpoint-1",
				EndpointType:       "teams",
				ID:                 1,
				LastContactTime:    &util.InfinityTime{Time: lastContact},
				LastWorker:         "worker-1",
				MJXIntegrationName: "integration-1",
				NumberOfMeetings:   3,
				ResourceURI:        "/api/admin/status/v1/mjx_endpoint/1/",
				RoomEmail:          "room1@example.com",
			},
			{
				EndpointAddress:    "10.0.0.2",
				EndpointName:       "mjx-endpoint-2",
				EndpointType:       "google",
				ID:                 2,
				LastContactTime:    &util.InfinityTime{Time: lastContact},
				LastWorker:         "worker-2",
				MJXIntegrationName: "integration-2",
				NumberOfMeetings:   0,
				ResourceURI:        "/api/admin/status/v1/mjx_endpoint/2/",
				RoomEmail:          "room2@example.com",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/mjx_endpoint/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.MJXEndpointListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MJXEndpointListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListMJXEndpoints(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "mjx-endpoint-1", result.Objects[0].EndpointName)
	assert.Equal(t, "10.0.0.1", result.Objects[0].EndpointAddress)
	assert.Equal(t, "teams", result.Objects[0].EndpointType)
	assert.Equal(t, 3, result.Objects[0].NumberOfMeetings)
	assert.Equal(t, "worker-1", result.Objects[0].LastWorker)
	assert.Equal(t, "integration-1", result.Objects[0].MJXIntegrationName)
	assert.Equal(t, "room1@example.com", result.Objects[0].RoomEmail)
	assert.Equal(t, "/api/admin/status/v1/mjx_endpoint/1/", result.Objects[0].ResourceURI)
	assert.Equal(t, "mjx-endpoint-2", result.Objects[1].EndpointName)
	assert.Equal(t, "10.0.0.2", result.Objects[1].EndpointAddress)
	assert.Equal(t, "google", result.Objects[1].EndpointType)
	assert.Equal(t, 0, result.Objects[1].NumberOfMeetings)
	assert.Equal(t, "worker-2", result.Objects[1].LastWorker)
	assert.Equal(t, "integration-2", result.Objects[1].MJXIntegrationName)
	assert.Equal(t, "room2@example.com", result.Objects[1].RoomEmail)
	assert.Equal(t, "/api/admin/status/v1/mjx_endpoint/2/", result.Objects[1].ResourceURI)
	client.AssertExpectations(t)
}

func TestService_GetMJXEndpoint(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	lastContact := time.Now().Add(-2 * time.Minute)

	expectedEndpoint := &MJXEndpoint{
		EndpointAddress:    "10.0.0.10",
		EndpointName:       "mjx-primary-endpoint",
		EndpointType:       "teams",
		ID:                 1,
		LastContactTime:    &util.InfinityTime{Time: lastContact},
		LastWorker:         "worker-primary",
		MJXIntegrationName: "integration-primary",
		NumberOfMeetings:   8,
		ResourceURI:        "/api/admin/status/v1/mjx_endpoint/1/",
		RoomEmail:          "room-primary@example.com",
	}

	client.On("GetJSON", t.Context(), "status/v1/mjx_endpoint/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.MJXEndpoint")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MJXEndpoint)
		*result = *expectedEndpoint
	})

	service := New(client)
	result, err := service.GetMJXEndpoint(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedEndpoint, result)
	assert.Equal(t, "mjx-primary-endpoint", result.EndpointName)
	assert.Equal(t, "10.0.0.10", result.EndpointAddress)
	assert.Equal(t, "teams", result.EndpointType)
	assert.Equal(t, 8, result.NumberOfMeetings)
	assert.Equal(t, "worker-primary", result.LastWorker)
	assert.Equal(t, "integration-primary", result.MJXIntegrationName)
	assert.Equal(t, "room-primary@example.com", result.RoomEmail)
	assert.Equal(t, "/api/admin/status/v1/mjx_endpoint/1/", result.ResourceURI)
	client.AssertExpectations(t)
}

func TestService_ListMJXEndpoints_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  10,
		Offset: 0,
	}

	expectedResponse := &MJXEndpointListResponse{
		Objects: []MJXEndpoint{
			{
				EndpointAddress:    "10.0.0.3",
				EndpointName:       "mjx-test-endpoint",
				EndpointType:       "webex",
				ID:                 3,
				LastContactTime:    nil,
				LastWorker:         "worker-test",
				MJXIntegrationName: "integration-test",
				NumberOfMeetings:   0,
				ResourceURI:        "/api/admin/status/v1/mjx_endpoint/3/",
				RoomEmail:          "room3@example.com",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/mjx_endpoint/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.MJXEndpointListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MJXEndpointListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListMJXEndpoints(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "mjx-test-endpoint", result.Objects[0].EndpointName)
	assert.Equal(t, "10.0.0.3", result.Objects[0].EndpointAddress)
	assert.Equal(t, "webex", result.Objects[0].EndpointType)
	assert.Equal(t, 0, result.Objects[0].NumberOfMeetings)
	assert.Equal(t, "worker-test", result.Objects[0].LastWorker)
	assert.Equal(t, "integration-test", result.Objects[0].MJXIntegrationName)
	assert.Equal(t, "room3@example.com", result.Objects[0].RoomEmail)
	assert.Equal(t, "/api/admin/status/v1/mjx_endpoint/3/", result.Objects[0].ResourceURI)

	client.AssertExpectations(t)
}
