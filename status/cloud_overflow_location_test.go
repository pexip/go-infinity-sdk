package status

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListCloudOverflowLocations(t *testing.T) {
	client := &mockClient.Client{}

	expectedResponse := &CloudOverflowLocationListResponse{
		Objects: []CloudOverflowLocation{
			{
				ID:            1,
				Name:          "Overflow US East",
				Region:        "us-east-1",
				Status:        "active",
				InstanceCount: 2,
				ResourceURI:   "/api/admin/status/v1/cloud_overflow_location/1/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_overflow_location/", mock.AnythingOfType("*status.CloudOverflowLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudOverflowLocationListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListCloudOverflowLocations(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Overflow US East", result.Objects[0].Name)
	assert.Equal(t, "active", result.Objects[0].Status)
	client.AssertExpectations(t)
}

func TestService_ListCloudOverflowLocations_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  3,
		Offset: 1,
	}

	expectedResponse := &CloudOverflowLocationListResponse{
		Objects: []CloudOverflowLocation{
			{
				ID:            2,
				Name:          "Test Overflow",
				Region:        "eu-west-1",
				Status:        "active",
				InstanceCount: 2,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/cloud_overflow_location/"
	}), mock.AnythingOfType("*status.CloudOverflowLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudOverflowLocationListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListCloudOverflowLocations(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Test Overflow", result.Objects[0].Name)

	client.AssertExpectations(t)
}

func TestService_GetCloudOverflowLocation(t *testing.T) {
	client := &mockClient.Client{}

	expectedLocation := &CloudOverflowLocation{
		ID:            1,
		Name:          "Primary Overflow",
		Region:        "us-west-2",
		Status:        "active",
		InstanceCount: 5,
		ResourceURI:   "/api/admin/status/v1/cloud_overflow_location/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_overflow_location/1/", mock.AnythingOfType("*status.CloudOverflowLocation")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudOverflowLocation)
		*result = *expectedLocation
	})

	service := New(client)
	result, err := service.GetCloudOverflowLocation(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, result)
	client.AssertExpectations(t)
}
