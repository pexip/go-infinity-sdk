package status

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListCloudMonitoredLocations(t *testing.T) {
	client := &mockClient.Client{}

	expectedResponse := &CloudMonitoredLocationListResponse{
		Objects: []CloudMonitoredLocation{
			{
				ID:               1,
				Name:             "AWS US West",
				OverflowLocation: "/api/admin/status/v1/cloud_overflow_location/2/",
				MaxHDCalls:       10,
				FreeHDCalls:      8,
				MediaLoad:        20,
				ResourceURI:      "/api/admin/status/v1/cloud_monitored_location/1/",
			},
			{
				ID:               2,
				Name:             "Azure EU West",
				OverflowLocation: "/api/admin/status/v1/cloud_overflow_location/3/",
				MaxHDCalls:       8,
				FreeHDCalls:      6,
				MediaLoad:        40,
				ResourceURI:      "/api/admin/status/v1/cloud_monitored_location/2/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_monitored_location/", mock.AnythingOfType("*status.CloudMonitoredLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudMonitoredLocationListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListCloudMonitoredLocations(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "AWS US West", result.Objects[0].Name)
	assert.Equal(t, "/api/admin/status/v1/cloud_overflow_location/2/", result.Objects[0].OverflowLocation)
	assert.Equal(t, 10, result.Objects[0].MaxHDCalls)
	assert.Equal(t, 8, result.Objects[0].FreeHDCalls)
	assert.Equal(t, 20, result.Objects[0].MediaLoad)
	assert.Equal(t, "Azure EU West", result.Objects[1].Name)
	client.AssertExpectations(t)
}

func TestService_GetCloudMonitoredLocation(t *testing.T) {
	client := &mockClient.Client{}

	expectedLocation := &CloudMonitoredLocation{
		ID:               1,
		Name:             "GCP Asia Pacific",
		OverflowLocation: "/api/admin/status/v1/cloud_overflow_location/4/",
		MaxHDCalls:       12,
		FreeHDCalls:      10,
		MediaLoad:        30,
		ResourceURI:      "/api/admin/status/v1/cloud_monitored_location/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_monitored_location/1/", mock.AnythingOfType("*status.CloudMonitoredLocation")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudMonitoredLocation)
		*result = *expectedLocation
	})

	service := New(client)
	result, err := service.GetCloudMonitoredLocation(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, result)
	assert.Equal(t, "GCP Asia Pacific", result.Name)
	assert.Equal(t, "/api/admin/status/v1/cloud_overflow_location/4/", result.OverflowLocation)
	assert.Equal(t, 12, result.MaxHDCalls)
	assert.Equal(t, 10, result.FreeHDCalls)
	assert.Equal(t, 30, result.MediaLoad)
	client.AssertExpectations(t)
}

func TestService_ListCloudMonitoredLocations_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  10,
		Offset: 0,
	}

	expectedResponse := &CloudMonitoredLocationListResponse{
		Objects: []CloudMonitoredLocation{
			{
				ID:               3,
				Name:             "AWS EU Central",
				OverflowLocation: "/api/admin/status/v1/cloud_overflow_location/5/",
				MaxHDCalls:       9,
				FreeHDCalls:      7,
				MediaLoad:        25,
				ResourceURI:      "/api/admin/status/v1/cloud_monitored_location/3/",
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/cloud_monitored_location/"
	}), mock.AnythingOfType("*status.CloudMonitoredLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudMonitoredLocationListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListCloudMonitoredLocations(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "AWS EU Central", result.Objects[0].Name)
	assert.Equal(t, "/api/admin/status/v1/cloud_overflow_location/5/", result.Objects[0].OverflowLocation)
	assert.Equal(t, 9, result.Objects[0].MaxHDCalls)
	assert.Equal(t, 7, result.Objects[0].FreeHDCalls)
	assert.Equal(t, 25, result.Objects[0].MediaLoad)

	client.AssertExpectations(t)
}
