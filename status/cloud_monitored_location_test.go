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

func TestService_ListCloudMonitoredLocations(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	name1 := "AWS US West"
	name2 := "Azure EU West"
	maxHD1 := 10
	maxHD2 := 8
	freeHD1 := 8
	freeHD2 := 6
	mediaLoad1 := 20
	mediaLoad2 := 40

	expectedResponse := &CloudMonitoredLocationListResponse{
		Meta: Meta{
			Limit:      100,
			Next:       "",
			Offset:     0,
			Previous:   "",
			TotalCount: 2,
		},
		Objects: []CloudMonitoredLocation{
			{
				ID:               1,
				Name:             &name1,
				OverflowLocation: "/api/admin/status/v1/cloud_overflow_location/2/",
				MaxHDCalls:       &maxHD1,
				FreeHDCalls:      &freeHD1,
				MediaLoad:        &mediaLoad1,
				ResourceURI:      "/api/admin/status/v1/cloud_monitored_location/1/",
			},
			{
				ID:               2,
				Name:             &name2,
				OverflowLocation: "/api/admin/status/v1/cloud_overflow_location/3/",
				MaxHDCalls:       &maxHD2,
				FreeHDCalls:      &freeHD2,
				MediaLoad:        &mediaLoad2,
				ResourceURI:      "/api/admin/status/v1/cloud_monitored_location/2/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_monitored_location/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.CloudMonitoredLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CloudMonitoredLocationListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListCloudMonitoredLocations(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, name1, derefString(result.Objects[0].Name))
	assert.Equal(t, "/api/admin/status/v1/cloud_overflow_location/2/", result.Objects[0].OverflowLocation)
	assert.Equal(t, maxHD1, derefInt(result.Objects[0].MaxHDCalls))
	assert.Equal(t, freeHD1, derefInt(result.Objects[0].FreeHDCalls))
	assert.Equal(t, mediaLoad1, derefInt(result.Objects[0].MediaLoad))
	assert.Equal(t, name2, derefString(result.Objects[1].Name))
	client.AssertExpectations(t)
}

func TestService_GetCloudMonitoredLocation(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	name := "GCP Asia Pacific"
	maxHD := 12
	freeHD := 10
	mediaLoad := 30

	expectedLocation := &CloudMonitoredLocation{
		ID:               1,
		Name:             &name,
		OverflowLocation: "/api/admin/status/v1/cloud_overflow_location/4/",
		MaxHDCalls:       &maxHD,
		FreeHDCalls:      &freeHD,
		MediaLoad:        &mediaLoad,
		ResourceURI:      "/api/admin/status/v1/cloud_monitored_location/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_monitored_location/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.CloudMonitoredLocation")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CloudMonitoredLocation)
		*result = *expectedLocation
	})

	service := New(client)
	result, err := service.GetCloudMonitoredLocation(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, result)
	assert.Equal(t, name, derefString(result.Name))
	assert.Equal(t, "/api/admin/status/v1/cloud_overflow_location/4/", result.OverflowLocation)
	assert.Equal(t, maxHD, derefInt(result.MaxHDCalls))
	assert.Equal(t, freeHD, derefInt(result.FreeHDCalls))
	assert.Equal(t, mediaLoad, derefInt(result.MediaLoad))
	client.AssertExpectations(t)
}

func TestService_ListCloudMonitoredLocations_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  10,
		Offset: 0,
	}

	name := "AWS EU Central"
	maxHD := 9
	freeHD := 7
	mediaLoad := 25

	expectedResponse := &CloudMonitoredLocationListResponse{
		Meta: Meta{
			Limit:      10,
			Next:       "",
			Offset:     0,
			Previous:   "",
			TotalCount: 1,
		},
		Objects: []CloudMonitoredLocation{
			{
				ID:               3,
				Name:             &name,
				OverflowLocation: "/api/admin/status/v1/cloud_overflow_location/5/",
				MaxHDCalls:       &maxHD,
				FreeHDCalls:      &freeHD,
				MediaLoad:        &mediaLoad,
				ResourceURI:      "/api/admin/status/v1/cloud_monitored_location/3/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_monitored_location/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.CloudMonitoredLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CloudMonitoredLocationListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListCloudMonitoredLocations(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, name, derefString(result.Objects[0].Name))
	assert.Equal(t, "/api/admin/status/v1/cloud_overflow_location/5/", result.Objects[0].OverflowLocation)
	assert.Equal(t, maxHD, derefInt(result.Objects[0].MaxHDCalls))
	assert.Equal(t, freeHD, derefInt(result.Objects[0].FreeHDCalls))
	assert.Equal(t, mediaLoad, derefInt(result.Objects[0].MediaLoad))
	assert.Equal(t, 1, result.Meta.TotalCount)

	client.AssertExpectations(t)
}
