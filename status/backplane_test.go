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
	util "github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListBackplanes(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedResponse := &BackplaneListResponse{
		Objects: []Backplane{
			{
				ID:                   "backplane-1",
				Conference:           "Test Conference",
				Type:                 "local-backplane",
				Protocol:             "INTERNAL",
				ServiceTag:           "tag123",
				SystemLocation:       "main-site",
				MediaNode:            "node1",
				ProxyNode:            "proxy1",
				RemoteConferenceName: "Remote Conference",
				RemoteMediaNode:      "remote-node1",
				RemoteNodeName:       "remote-name",
				ResourceURI:          "/api/admin/status/v1/backplane/backplane-1/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/backplane/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.BackplaneListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*BackplaneListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListBackplanes(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "backplane-1", result.Objects[0].ID)
	assert.Equal(t, "Test Conference", result.Objects[0].Conference)
	assert.Equal(t, "local-backplane", result.Objects[0].Type)
	client.AssertExpectations(t)
}

func TestService_GetBackplane(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	connectTime := time.Now()
	expectedBackplane := &Backplane{
		ID:                   "backplane-1",
		Conference:           "Test Conference",
		Type:                 "geo-backplane",
		Protocol:             "MSSIP",
		ConnectTime:          &util.InfinityTime{Time: connectTime},
		ServiceTag:           "tag123",
		SystemLocation:       "main-site",
		MediaNode:            "node1",
		ProxyNode:            "proxy1",
		RemoteConferenceName: "Remote Conference",
		RemoteMediaNode:      "remote-node1",
		RemoteNodeName:       "remote-name",
		ResourceURI:          "/api/admin/status/v1/backplane/backplane-1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/backplane/backplane-1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.Backplane")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Backplane)
		*result = *expectedBackplane
	})

	service := New(client)
	result, err := service.GetBackplane(t.Context(), "backplane-1")

	assert.NoError(t, err)
	assert.Equal(t, expectedBackplane, result)
	client.AssertExpectations(t)
}

func TestService_ListBackplanes_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  10,
		Offset: 5,
	}

	expectedResponse := &BackplaneListResponse{
		Objects: []Backplane{
			{
				ID:         "backplane-1",
				Conference: "Test Conference",
				Type:       "external-backplane",
				Protocol:   "TEAMS",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/backplane/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.BackplaneListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*BackplaneListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListBackplanes(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "external-backplane", result.Objects[0].Type)
	assert.Equal(t, "TEAMS", result.Objects[0].Protocol)

	client.AssertExpectations(t)
}
