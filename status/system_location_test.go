/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSystemLocations(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedResponse := &SystemLocationListResponse{
		Objects: []SystemLocation{
			{
				ID:              1,
				Name:            "Main Office",
				ResourceURI:     "/api/admin/status/v1/system_location/1/",
				MaxAudioCalls:   100,
				MaxConnections:  "1000",
				MaxFullHDCalls:  10,
				MaxHDCalls:      20,
				MaxMediaTokens:  200,
				MaxSDCalls:      30,
				MediaLoad:       0.5,
				MediaTokensUsed: 50,
			},
			{
				ID:              2,
				Name:            "Branch Office",
				ResourceURI:     "/api/admin/status/v1/system_location/2/",
				MaxAudioCalls:   50,
				MaxConnections:  "500",
				MaxFullHDCalls:  5,
				MaxHDCalls:      10,
				MaxMediaTokens:  100,
				MaxSDCalls:      15,
				MediaLoad:       0.2,
				MediaTokensUsed: 20,
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/system_location/", mock.AnythingOfType("*status.SystemLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemLocationListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListSystemLocations(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "Main Office", result.Objects[0].Name)
	assert.Equal(t, "/api/admin/status/v1/system_location/1/", result.Objects[0].ResourceURI)
	assert.Equal(t, 100, result.Objects[0].MaxAudioCalls)
	assert.Equal(t, "1000", result.Objects[0].MaxConnections)
	assert.Equal(t, 10, result.Objects[0].MaxFullHDCalls)
	assert.Equal(t, 20, result.Objects[0].MaxHDCalls)
	assert.Equal(t, 200, result.Objects[0].MaxMediaTokens)
	assert.Equal(t, 30, result.Objects[0].MaxSDCalls)
	assert.Equal(t, 0.5, result.Objects[0].MediaLoad)
	assert.Equal(t, 50, result.Objects[0].MediaTokensUsed)
	assert.Equal(t, "Branch Office", result.Objects[1].Name)
	assert.Equal(t, "/api/admin/status/v1/system_location/2/", result.Objects[1].ResourceURI)
	assert.Equal(t, 50, result.Objects[1].MaxAudioCalls)
	assert.Equal(t, "500", result.Objects[1].MaxConnections)
	assert.Equal(t, 5, result.Objects[1].MaxFullHDCalls)
	assert.Equal(t, 10, result.Objects[1].MaxHDCalls)
	assert.Equal(t, 100, result.Objects[1].MaxMediaTokens)
	assert.Equal(t, 15, result.Objects[1].MaxSDCalls)
	assert.Equal(t, 0.2, result.Objects[1].MediaLoad)
	assert.Equal(t, 20, result.Objects[1].MediaTokensUsed)
	client.AssertExpectations(t)
}

func TestService_GetSystemLocation(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedLocation := &SystemLocation{
		ID:              1,
		Name:            "Headquarters",
		ResourceURI:     "/api/admin/status/v1/system_location/1/",
		MaxAudioCalls:   200,
		MaxConnections:  "2000",
		MaxFullHDCalls:  20,
		MaxHDCalls:      40,
		MaxMediaTokens:  400,
		MaxSDCalls:      60,
		MediaLoad:       0.8,
		MediaTokensUsed: 100,
	}

	client.On("GetJSON", t.Context(), "status/v1/system_location/1/", mock.AnythingOfType("*status.SystemLocation")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemLocation)
		*result = *expectedLocation
	})

	service := New(client)
	result, err := service.GetSystemLocation(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, result)
	assert.Equal(t, "Headquarters", result.Name)
	assert.Equal(t, "/api/admin/status/v1/system_location/1/", result.ResourceURI)
	assert.Equal(t, 200, result.MaxAudioCalls)
	assert.Equal(t, "2000", result.MaxConnections)
	assert.Equal(t, 20, result.MaxFullHDCalls)
	assert.Equal(t, 40, result.MaxHDCalls)
	assert.Equal(t, 400, result.MaxMediaTokens)
	assert.Equal(t, 60, result.MaxSDCalls)
	assert.Equal(t, 0.8, result.MediaLoad)
	assert.Equal(t, 100, result.MediaTokensUsed)
	client.AssertExpectations(t)
}

func TestService_ListSystemLocations_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  5,
		Offset: 0,
	}

	expectedResponse := &SystemLocationListResponse{
		Objects: []SystemLocation{
			{
				ID:              3,
				Name:            "Remote Site",
				ResourceURI:     "/api/admin/status/v1/system_location/3/",
				MaxAudioCalls:   10,
				MaxConnections:  "100",
				MaxFullHDCalls:  1,
				MaxHDCalls:      2,
				MaxMediaTokens:  20,
				MaxSDCalls:      3,
				MediaLoad:       0.1,
				MediaTokensUsed: 5,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/system_location/"
	}), mock.AnythingOfType("*status.SystemLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemLocationListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListSystemLocations(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Remote Site", result.Objects[0].Name)
	assert.Equal(t, "/api/admin/status/v1/system_location/3/", result.Objects[0].ResourceURI)
	assert.Equal(t, 10, result.Objects[0].MaxAudioCalls)
	assert.Equal(t, "100", result.Objects[0].MaxConnections)
	assert.Equal(t, 1, result.Objects[0].MaxFullHDCalls)
	assert.Equal(t, 2, result.Objects[0].MaxHDCalls)
	assert.Equal(t, 20, result.Objects[0].MaxMediaTokens)
	assert.Equal(t, 3, result.Objects[0].MaxSDCalls)
	assert.Equal(t, 0.1, result.Objects[0].MediaLoad)
	assert.Equal(t, 5, result.Objects[0].MediaTokensUsed)

	client.AssertExpectations(t)
}
