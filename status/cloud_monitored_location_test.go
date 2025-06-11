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
				ID:            1,
				Name:          "AWS US West",
				Region:        "us-west-2",
				Status:        "healthy",
				InstanceCount: 5,
				ActiveCalls:   15,
				ResourceURI:   "/api/admin/status/v1/cloud_monitored_location/1/",
			},
			{
				ID:            2,
				Name:          "Azure EU West",
				Region:        "eu-west-1",
				Status:        "degraded",
				InstanceCount: 3,
				ActiveCalls:   8,
				ResourceURI:   "/api/admin/status/v1/cloud_monitored_location/2/",
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
	assert.Equal(t, "us-west-2", result.Objects[0].Region)
	assert.Equal(t, "healthy", result.Objects[0].Status)
	assert.Equal(t, 5, result.Objects[0].InstanceCount)
	assert.Equal(t, 15, result.Objects[0].ActiveCalls)
	assert.Equal(t, "Azure EU West", result.Objects[1].Name)
	assert.Equal(t, "degraded", result.Objects[1].Status)
	client.AssertExpectations(t)
}

func TestService_GetCloudMonitoredLocation(t *testing.T) {
	client := &mockClient.Client{}

	expectedLocation := &CloudMonitoredLocation{
		ID:            1,
		Name:          "GCP Asia Pacific",
		Region:        "asia-pacific-1",
		Status:        "healthy",
		InstanceCount: 8,
		ActiveCalls:   32,
		ResourceURI:   "/api/admin/status/v1/cloud_monitored_location/1/",
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
	assert.Equal(t, "asia-pacific-1", result.Region)
	assert.Equal(t, "healthy", result.Status)
	assert.Equal(t, 8, result.InstanceCount)
	assert.Equal(t, 32, result.ActiveCalls)
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
				ID:            3,
				Name:          "AWS EU Central",
				Region:        "eu-central-1",
				Status:        "healthy",
				InstanceCount: 4,
				ActiveCalls:   18,
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
	assert.Equal(t, "eu-central-1", result.Objects[0].Region)
	assert.Equal(t, "healthy", result.Objects[0].Status)

	client.AssertExpectations(t)
}
